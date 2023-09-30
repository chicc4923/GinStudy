package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type UserInfo struct { // 数据库的表名为 user_infos
	ID    int
	Name  string
	Hobby string
}

func connDB() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/blog?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	// 创建表，自动迁移(把结构体和表进行对应)
	db.AutoMigrate(&UserInfo{})
}
