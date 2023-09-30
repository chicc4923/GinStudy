package test

import (
	"Gin_study/table"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
	"testing"
)

const (
	dbDriver = "mysql"
	dbSource = "root:123456@tcp(127.0.0.1:3306)/blog?charset=utf8mb4&parseTime=True&loc=Local"
)

var DB *gorm.DB

// TestMain:Golang 约定 TestMain 函数是所有单元测试的入口
func TestMain(m *testing.M) {
	db, err := gorm.Open(mysql.Open(dbSource), &gorm.Config{})
	if err != nil {
		log.Fatal("can not connect db:", err)
	}
	DB = db
	os.Exit(m.Run())
}

// Go中的每个单元测试函数都必须以 Test 开头，并且以 testing.T 作为输入参数。
// TestAddUser 新增用户
func TestAddUser(t *testing.T) {
	u1 := table.UserInfo{2, "max", "run"}
	DB.Create(u1)
}

// TestUpdateUser 更新用户
func TestUpdateUser(t *testing.T) {
	u1 := table.UserInfo{2, "max2", "run"}
	DB.Updates(u1)
}

// TestDeleteUser 删除用户
func TestDeleteUser(t *testing.T) {
	DB.Delete(table.UserInfo{2, "max2", "run"})
}

// TestGetUser 获取用户
func TestGetUser(t *testing.T) {

	var u table.UserInfo
	// 根据主键查询第一个记录
	DB.First(&u) // First 和 Last 会根据主键排序，分别查询第一条和最后一条数据。只有在目标 struct 是指针或者通过 db.Model() 指定 model 时，该方法才有效。
	//  这里只能传指针，我的理解是查询到的结果需要传入原结构体而不是副本。
	fmt.Println(u)
	// DB.Find() 也是需要传入结构体切片的指针，返回的是所有记录
	//DB.Find(table.UserInfo{2, "max", "run"})
}
