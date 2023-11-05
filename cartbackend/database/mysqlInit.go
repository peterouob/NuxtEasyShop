package database

import (
	"easyshop/model"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB
var err error

func MysqlInit() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/easyshop?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn))
	if err != nil {
		log.Println("mysql init error")
	}
	DB.AutoMigrate(&model.User{}, &model.Good{})
	fmt.Println("mysql init...")
}
