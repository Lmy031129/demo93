package handler

import (
	"context"
	"errors"
	"fmt"
	"user-srv/basic/globar"
	__ "user-srv/basic/proto"
	"user-srv/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// server is used to implement helloworld.GreeterServer.
type Server struct {
	__.UnimplementedUserServer
}

// SayHello implements helloworld.GreeterServer
func (s *Server) Register(_ context.Context, in *__.RegisterReq) (*__.RegisterResp, error) {

	return &__.RegisterResp{}, nil
}

func (s *Server) Shopadd(_ context.Context, in *__.ShopaddReq) (*__.ShopaddResp, error) {
	var shop model.Shop
	shop = model.Shop{
		Title:  in.Title,
		Num:    int(in.Num),
		Price:  float64(in.Price),
		Count:  int(in.Count),
		Status: int(in.Status),
	}
	err := globar.DB.Create(&shop).Error
	if err != nil {
		return nil, errors.New("添加失败")
	}
	return &__.ShopaddResp{
		Id: int64(shop.ID),
	}, nil
}
func (s *Server) OrderItemadd(_ context.Context, in *__.OrderItemaddReq) (*__.OrderItemaddResp, error) {
	var item model.OrderItem
	tx := globar.DB.Begin()
	result, _ := globar.Rdb.Get(context.Background(), "shop").Result()
	if result == "" {
		return nil, errors.New("下单失败")
		tx.Rollback()
	}

	item = model.OrderItem{
		UserId:     in.UserId,
		OrderSn:    in.OrderSn,
		OrderNum:   in.OrderNum,
		OrderPrice: float64(in.OrderPrice),
	}
	err := globar.DB.Create(&item).Error
	fmt.Println(err)
	if err != nil {
		return nil, errors.New("下单失败")
		tx.Rollback()
	}
	tx.Commit()
	return &__.OrderItemaddResp{
		Id: int64(item.ID),
	}, nil
}
func (s *Server) OrderItemUpdate(_ context.Context, in *__.OrderItemUpdateReq) (*__.OrderItemUpdateResp, error) {
	var item model.OrderItem
	tx := globar.DB.Begin()
	item = model.OrderItem{
		Model:  gorm.Model{ID: uint(in.Id)},
		Status: int(in.Status),
	}
	err := globar.DB.Where("id = ?", in.Id).Updates(&item).Error
	if err != nil {
		return nil, errors.New("修改失败")
		tx.Rollback()
	}
	tx.Commit()
	return &__.OrderItemUpdateResp{
		Id: int64(item.ID),
	}, nil
}
func (s *Server) OrderItemLike(_ context.Context, in *__.OrderItemLikeReq) (*__.OrderItemLikeResp, error) {
	var item model.OrderItem
	tx := globar.DB.Begin()

	err := globar.DB.Where("order_sn like ?", "%"+in.OrderSn+"%").Find(&item).Limit(1).Error
	if err != nil {
		return nil, errors.New("模糊搜索失败")
		tx.Rollback()
	}
	tx.Commit()
	return &__.OrderItemLikeResp{
		UserId:     item.UserId,
		OrderSn:    item.OrderSn,
		OrderNum:   item.OrderNum,
		OrderPrice: float32(item.OrderPrice),
		Status:     int64(item.Status),
	}, nil
}
func (s *Server) ShopList(_ context.Context, in *__.ShopListReq) (*__.ShopListResp, error) {
	var p model.Shop
	cover, err := p.GetCover(globar.DB, in.Page, in.Size)
	if err != nil {
		return nil, errors.New("查询失败")
	}
	return &__.ShopListResp{
		List: cover,
	}, nil
}

func (s *Server) OrderItemShow(_ context.Context, in *__.OrderItemShowReq) (*__.OrderItemShowResp, error) {
	var u model.OrderItem
	err := globar.DB.Where("id = ?", in.Id).Find(&u).Limit(1).Error
	if err != nil {
		return nil, errors.New("详情查询失败")
	}
	return &__.OrderItemShowResp{
		UserId:     u.UserId,
		OrderSn:    u.OrderSn,
		OrderNum:   u.OrderNum,
		OrderPrice: float32(u.OrderPrice),
	}, nil
}
func (s *Server) OrderItemList(_ context.Context, in *__.OrderItemListReq) (*__.OrderItemListResp, error) {
	var o model.OrderItem
	cover, err := o.GetCover(globar.DB, in.Page, in.Size)
	fmt.Println(err)
	if err != nil {
		return nil, errors.New("查询失败")
	}
	return &__.OrderItemListResp{
		List: cover,
	}, nil
}
func (s *Server) OrderItemdel(_ context.Context, in *__.OrderItemdelReq) (*__.OrderItemdelResp, error) {
	var u model.User
	globar.DB.Find(&u)
	var p model.OrderItem
	globar.DB.Find(&p)
	if in.UserId == in.ShopId {
		return nil, errors.New("重复Id禁止下单")
	}
	order := uuid.New().String()
	
	return &__.OrderItemdelResp{
		OrderSn: order,
	}, nil
}
