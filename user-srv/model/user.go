package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Mobile   string `gorm:"type:varchar(20)" json:"mobile"`
	Password string `gorm:"type:varchar(32)" json:"password"`
}
