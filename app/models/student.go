package models

import (
	"gorm.io/gorm"
)

type Student struct {
	StudentID string         `gorm:"column:student_id"`
	Name      string         `gorm:"column:name"`
	Institute string         `gorm:"column:institute"`
	Password  string         `gorm:"column:password"`
	CreatedAt gorm.DeletedAt `gorm:"autoCreateTime"`
}
