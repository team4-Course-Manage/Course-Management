package models

import (
	"time"
)

type Task struct {
	TaskId          int       `gorm:"column:task_id;primaryKey;autoIncrement"`
	TaskName        string    `gorm:"column:task_name"`
	ProjectId       int       `gorm:"column:projeceID"`
	Subject         string    `gorm:"column:theme"`
	TaskDescription string    `gorm:"column:description"`
	ReceiverId      int       `gorm:"column:acceptor_ID"`
	PrincipalId     int       `gorm:"column:responible_member"`
	Priority        string    `gorm:"column:priority"`
	TaskDate        time.Time `gorm:"column:task_date"`
	TaskStatus      string    `gorm:"column:task_status"`
}
