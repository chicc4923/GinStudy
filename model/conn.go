package model

import (
	"Gin_study/table"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// SetupDB 链接并初始化数据库
func SetupDB() error {
	dsn := "root:123456@tcp(127.0.0.1:3306)/buble?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	// 创建表，自动迁移(把结构体和表进行对应)
	//db.AutoMigrate(&table.UserInfo{})
	err = db.AutoMigrate(&table.Todo{})

	if err != nil {
		return err
	}
	DB = db
	return nil
}

// GetDB 返回初始化后的数据库实例
func GetDB() *gorm.DB {
	return DB
}
