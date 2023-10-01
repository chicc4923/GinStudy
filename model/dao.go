package model

import (
	"Gin_study/table"
)

// CreateTodo 创建事项的数据库命令
func CreateTodo(data table.TodoList) error {
	if err := GetDB().Model(&table.TodoList{}).Create(data).Error; err != nil {
		return err
	}
	return nil
}

// GetTodoList 获取所有事项
func GetTodoList() error {
	lists := make([]table.TodoList, 0)
	if err := GetDB().Model(&table.TodoList{}).Find(&lists).Error; err != nil {
		return err
	}
	return nil
}

// GetTodoByStatus 获取某个事项
func GetTodoByStatus(status bool) error {
	lists := make([]table.TodoList, 0)
	if err := GetDB().Model(&table.TodoList{}).Find(&lists).Where("status = ?", status).Error; err != nil {
		return err
	}
	return nil
}

func UpdateTodo(data table.TodoList) error {
	if err := GetDB().Debug().Model(&table.TodoList{}).Where("id = ?", data.ID).Updates(&data).Error; err != nil {
		return err
	}
	return nil
}
