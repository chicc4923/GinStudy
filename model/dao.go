package model

import (
	"Gin_study/table"
	"errors"
)

// CreateTodo 创建事项的数据库命令
func CreateTodo(data table.Todo) error {
	if err := GetDB().Create(&data).Error; err != nil {
		return err
	}
	return nil
}

// GetTodo 获取所有事项
func GetTodo(lists *[]*table.Todo) (error, int64) {
	var total int64
	if err := GetDB().Model(&table.Todo{}).Debug().Limit(10).Offset(0).Count(&total).Find(&lists).Error; err != nil {
		return err, 0
	}
	return nil, total
}

// GetTodosByStatus 获取符合条件的事项
func GetTodosByStatus(status bool) error {
	lists := make([]table.Todo, 0)
	if err := GetDB().Model(&table.Todo{}).Find(&lists).Where("status = ?", status).Error; err != nil {
		return err
	}
	return nil
}

// GetTodoByID 根据 ID 获取某个事项
func GetTodoByID(id int) (table.Todo, error) {
	var todo table.Todo
	if err := GetDB().Model(&table.Todo{}).First(&todo).Where("id = ?", id).Error; err != nil {
		return table.Todo{}, err
	}
	return todo, nil
}

// UpdateTodosByID 根据 ID 和前端传入的 json 更新数据
func UpdateTodosByID(data table.Todo, id int) error {
	if id == 0 {
		return errors.New("ID 不存在！")
	}
	if err := GetDB().Debug().Model(&table.Todo{}).Where("id = ?", id).Save(&data).Error; err != nil {
		return err
	}
	return nil
}
func DeleteTodo(id int) error {
	if id == 0 {
		return errors.New("ID 不存在！")
	}
	if err := GetDB().Debug().Model(&table.Todo{}).Where("id = ?", id).Delete(&table.Todo{}).Error; err != nil {
		return err
	}
	return nil
}
