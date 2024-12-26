package models

import (
	"time"
)

type Task struct {
	TaskId          int       `gorm:"column:task_id;primaryKey;autoIncrement"`
	TaskName        string    `gorm:"column:task_name"`
	ProjectId       int       `gorm:"column:project_name"`
	Subject         string    `gorm:"column:subject"`
	TaskDescription string    `gorm:"column:task_description"`
	ReceiverId      int       `gorm:"column:receiver_id"`
	PrincipalId     int       `gorm:"column:principal_id"`
	Priority        string    `gorm:"column:priority"`
	TaskDate        time.Time `gorm:"column:task_date"`
	TaskStatus      string    `gorm:"column:task_status"`
}
