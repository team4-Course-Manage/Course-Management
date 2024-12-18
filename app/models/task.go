package models

import (
	"time"
)
type Task struct {
	TaskID      string `gorm:"column:task_id;not null"`      // 任务 ID
	ProjectID   string `gorm:"column:project_id;not null"`   // 项目 ID
	AssigneeID  string `gorm:"column:assignee_id;not null"`  // 分配任务的成员 ID
	Title       string `gorm:"column:title;not null"`        // 任务标题
	Description string `gorm:"column:description"`           // 任务描述
	Status      string `gorm:"column:status;default:'todo'"` // 状态，默认值为 "todo"
	DueDate     string `gorm:"column:due_date"`              // 截止日期
}

// TableName 指定表名
func (Task) TableName() string {
	return "task"
}
