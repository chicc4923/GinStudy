package model

import (
	"Gin_study/table"
)

// CreateTodo 创建事项的数据库命令
func CreateTodo(data table.Todo) error {
	if err := GetDB().Create(&data).Error; err != nil {
		return err
	}
	return nil
}

// GetTodo 获取所有事项
func GetTodo() error {
	lists := make([]table.Todo, 0)
	if err := GetDB().Model(&table.Todo{}).Find(&lists).Error; err != nil {
		return err
	}
	return nil
}

// GetTodoByStatus 获取某个事项
func GetTodoByStatus(status bool) error {
	lists := make([]table.Todo, 0)
	if err := GetDB().Model(&table.Todo{}).Find(&lists).Where("status = ?", status).Error; err != nil {
		return err
	}
	return nil
}

func UpdateTodo(data table.Todo) error {
	if err := GetDB().Debug().Model(&table.Todo{}).Where("id = ?", data.ID).Updates(&data).Error; err != nil {
		return err
	}
	return nil
}
