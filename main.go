package main

import (
	"Gin_study/table"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var user table.UserInfo
var users []table.UserInfo

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/blog?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	// 创建表，自动迁移(把结构体和表进行对应)
	db.AutoMigrate(&table.UserInfo{})
	//// 获取第一条记录（主键升序）
	//db.First(&user)
	//fmt.Println(user)
	//// 随机获取一条记录，没有指定排序字段
	//db.Take(&user)
	//fmt.Println(user)
	//// 获取最后一条记录（主键降序）
	//db.Last(&user)
	//fmt.Println(user)
	//
	//result := db.First(&user)
	//fmt.Println(result.RowsAffected) // 返回找到的记录数
	//fmt.Println(result.Error)        // returns error or nil
	//// 检查 ErrRecordNotFound 错误
	//errors.Is(result.Error, gorm.ErrRecordNotFound)

	//result := map[string]interface{}{}
	//db.Model(&user).First(&result)
	//fmt.Println(result)
	db.Save(&user).Where("name = hello,id")
	db.Model(&user).Where("name = ?", "max").Update("name", "hello")
}
