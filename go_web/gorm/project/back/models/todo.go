package models

import (
	"time"

	"go_web/gorm/project/back/dao"

	"gorm.io/gorm"
)

type Todo struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	Title     string         `json:"title"`
	Status    bool           `json:"status"`
}

// 新增笔记
func CreateTodo(todo *Todo) error {
	return dao.DB.Create(todo).Error
}

// 获取笔记
func GetTodoList() (todos []Todo, err error) {
	err = dao.DB.Find(&todos).Error
	return
}

// 更新笔记状态
func UpdateTodo(id string, todo *Todo) error {
	return dao.DB.Model(&Todo{}).Where("id = ?", id).Update("status", todo.Status).Error
}

// 删除笔记
func DeleteTodo(id string) error {
	return dao.DB.Delete(&Todo{}, id).Error
}
