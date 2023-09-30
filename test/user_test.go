package test

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
	"testing"
)

type UserInfo struct { // 数据库的表名为 user_infos
	ID    int
	Name  string
	Hobby string
}

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
	u1 := UserInfo{2, "max", "run"}
	DB.Create(u1)
}

// TestUpdateUser 更新用户
func TestUpdateUser(t *testing.T) {
	u1 := UserInfo{2, "max2", "run"}
	DB.Updates(u1)
}

// TestDeleteUser 删除用户
func TestDeleteUser(t *testing.T) {
	DB.Delete(UserInfo{2, "max2", "run"})
}

// TestGetUser 获取用户
func TestGetUser(t *testing.T) {
	DB.Find(UserInfo{2, "max", "run"})
}
