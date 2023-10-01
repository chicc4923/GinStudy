package demo

import (
	"Gin_study/table"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// var user table.UserInfo
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
	//user.ID = 3
	//user.Name = "xiaoming"
	//user.Hobby = "game"
	//db.Save(&user)
	//db.Model(&table.UserInfo{}).Where("name = ?", "max").Update("name", "hello")
	//user.ID = 1
	//db.Model(&user).Update("name", "max")
	// 使用 map or struct 更新多个字段，只能更新非零值字段
	//db.Model(&user).Where("id = ?", 1).Updates(map[string]interface{}{"name": "max", "hobby": "ride"})
	//db.Model(&user).Where("id = ?", 1).Updates(table.UserInfo{
	//	ID:    1,
	//	Name:  "max2",
	//	Hobby: "code",
	//})
	//// 指定字段更新：
	//user.ID = 1
	////db.Model(&user).Select("name").Updates(map[string]interface{}{"name": "max1", "hobby": "cook"}) //这里只选择了 name 字段，所以哪怕其他字段不同也只更新 name
	//db.Model(&user).Omit("name").Updates(map[string]interface{}{"name": "test", "hobby": "cook"}) // 这里忽略了 name，所以会更新其他字段
	//db.Debug().Model(&user).Select("name").Updates(table.UserInfo{
	//	ID:    1,
	//	Name:  "max",
	//	Hobby: "code",
	//})
	//db.Model(&table.UserInfo{}).Where("name = ?", "xiaoming").Updates(table.UserInfo{
	//	Name:  "xiaohua",
	//	Hobby: "IT",
	//})
	// 删除一条记录
	//user.ID = 3
	//db.Delete(&user)
	//db.Where("name = ?", "xiaohua").Delete(&user)
}
