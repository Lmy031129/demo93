package model

import "gorm.io/gorm"

type PaymentRecord struct {
	gorm.Model
	UserId      int64  `gorm:"type:int(11)" json:"user_id"`
	OrderSn     string `gorm:"type:varchar(100)" json:"order_sn"`
	OrderNum    int64  `gorm:"type:int(11)" json:"order_num"`
	ProductName string `gorm:"type:varchar(30)" json:"product_name"`
}
