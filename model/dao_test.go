package model

import (
	"Gin_study/table"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
	"testing"
)

// TestMain:Golang 约定 TestMain 函数是所有单元测试的入口
func TestMain(m *testing.M) {
	db, err := gorm.Open(mysql.Open("root:123456@tcp(127.0.0.1:3306)/buble?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		log.Fatal("can not connect db:", err)
	}
	DB = db
	os.Exit(m.Run())
}

// TestCreateTodo 新增事项测试
func TestCreateTodo(t *testing.T) {
	data := table.Todo{
		ID:    0,
		Title: "testdefault1",
		//Status: ,
	}
	//data := table.Todo{ID: 0, Title: "1", Status: sql.NullBool{Bool: false, Valid: false}}
	err := CreateTodo(data)
	if err != nil {
		log.Fatal(err)
	}
}

// TestGetTodo 查询所有事项测试
func TestGetTodo(t *testing.T) {
	list := make([]*table.Todo, 0)
	var total int64
	var err error
	if err, total = GetTodo(&list); err != nil {
		log.Fatal(err)
	}
	if len(list) == 0 || total == 0 {
		log.Fatal("null data")
	}
	fmt.Println(list, total)
}

// TestGetTodoByStatus 根据可用与否获取事项
func TestGetTodosByStatus(t *testing.T) {
	if err := GetTodosByStatus(true); err != nil {
		log.Fatal(err)
	}
}

func TestUpdateTodo(t *testing.T) {
	data := table.Todo{
		ID:     31,
		Title:  "code",
		Status: false,
		//Status: sql.NullBool{ // 如果不这样定义，因为 bool 类型的零值为 false，所以使用 struct 构建时条件语句时不会影响 status字段
		//	Bool:  false,
		//	Valid: true,
		//},
	}

	if err := UpdateTodosByID(data, 31); err != nil {
		log.Fatal(err)
	}
}

// TestDeleteTodo 测试根据 ID 删除事项
func TestDeleteTodo(t *testing.T) {
	id := 2
	err := DeleteTodo(id)
	if err != nil {
		log.Fatal(err)
	}
}

// TestGetTodoByID 测试根据 ID 获取某个事项
func TestGetTodoByID(t *testing.T) {
	id := 2
	todo, err := GetTodoByID(id)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(todo)
}
