package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	UserId        int64   `gorm:"type:int(11)" json:"user_id"`
	ProductName   string  `gorm:"type:varchar(30)" json:"product_name"`
	ProductNum    int     `gorm:"type:int(11)" json:"product_num"`
	ProductPrice  float64 `gorm:"type:decimal(10,2)" json:"product_price"`
	ProductStatus int     `gorm:"type:int(11)" json:"product_status"`
}
