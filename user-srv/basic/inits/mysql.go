package inits

import (
	"fmt"
	"log"
	"sync"
	"time"
	"user-srv/basic/config"
	"user-srv/basic/globar"
	"user-srv/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitMysql() {
	var err error
	data := config.ConfigAppData.Mysql
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", data.User, data.Password, data.Host, data.Port, data.DataBase)
	once := sync.Once{}
	once.Do(func() {
		globar.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			panic(err)
		}
	})
	globar.DB.AutoMigrate(&model.User{}, &model.Shop{}, &model.OrderItem{}, &model.Product{}, &model.UserAddress{}, &model.PaymentRecord{}, &model.UserAddress{})
	sqlDB, err := globar.DB.DB()

	// SetMaxIdleConns 设置空闲连接池中连接的最大数量。
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(5)

	// SetConnMaxLifetime 设置了可以重新使用连接的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)
	log.Println("Mysql is success")
}
