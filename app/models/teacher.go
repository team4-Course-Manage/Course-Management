package models

import (
	"gorm.io/gorm"
)

type Teacher struct {
	TeacherID string         `gorm:"column:teacher_ID"`
	Name      string         `gorm:"column:name"`
	Institute string         `gorm:"column:institute"`
	Password  string         `gorm:"column:password"`
	CreatedAt gorm.DeletedAt `gorm:"autoCreateTime"`
}
