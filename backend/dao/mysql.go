package dao

import (
	"fmt"
	"log"
	"project/backend/setting"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type user struct {
}

var DB *gorm.DB

func connect() {
	//建立连接
	var set setting.Mysql
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", set.Username, set.Password, set.Host, set.Dbname)
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	//创建表格
	if err := DB.Migrator().HasTable(&user{}); err == false {
		DB.AutoMigrate(&user{})
	}
}
