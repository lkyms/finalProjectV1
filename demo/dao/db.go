package dao

import (
	"demo/util"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True",
		util.GetConfig("mysql.user"),
		util.GetConfig("mysql.pwd"),
		util.GetConfig("mysql.address"),
		util.GetConfig("mysql.dbname"))
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("connect DB failed, err:%v\n", err)
		return
	}

	return
}
