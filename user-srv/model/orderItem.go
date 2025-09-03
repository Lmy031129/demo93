package model

import (
	__ "user-srv/basic/proto"

	"gorm.io/gorm"
)

type OrderItem struct {
	gorm.Model
	UserId     int64   `gorm:"type:int(11)" json:"user_id"`
	OrderSn    string  `gorm:"type:varchar(30)" json:"order_sn"`
	OrderNum   int64   `gorm:"type:int(11)" json:"order_num"`
	OrderPrice float64 `gorm:"type:decimal(10,2)" json:"order_price"`
	Status     int     `gorm:"type:int(11)" json:"status"`
}

func (o *OrderItem) GetCover(db *gorm.DB, Page, Size int64) (list []*__.OrderItemList, err error) {
	page := Page
	if page <= 0 {
		page = 1
	}
	pageSize := Size
	switch {
	case pageSize > 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}

	offset := (page - 1) * pageSize
	err = db.Model(&o).Offset(int(offset)).Limit(int(pageSize)).Find(&list).Error
	return list, err
}
