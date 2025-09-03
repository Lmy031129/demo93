package model

import "gorm.io/gorm"

type UserAddress struct {
	gorm.Model
	Name    string `gorm:"type:varchar(30)" json:"name"`
	Phone   int64  `gorm:"type:int(11)" json:"phone"`
	Address string `gorm:"type:varchar(40)" json:"address"`
}
