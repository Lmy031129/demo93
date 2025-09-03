package model

import (
	__ "user-srv/basic/proto"

	"gorm.io/gorm"
)

type Shop struct {
	gorm.Model
	Title  string  `gorm:"type:varchar(30)" json:"title"`
	Num    int     `gorm:"type:int(11)" json:"num"`
	Price  float64 `gorm:"type:decimal(10,2)" json:"price"`
	Count  int     `gorm:"type:int(11)" json:"count"`
	Status int     `gorm:"type:int(11)" json:"status"`
}

func (s *Shop) GetCover(db *gorm.DB, Page, Size int64) (list []*__.ShopList, err error) {
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
	err = db.Model(&s).Offset(int(offset)).Limit(int(pageSize)).Find(&list).Error
	return list, err
}
