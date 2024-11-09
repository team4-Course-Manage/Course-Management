package models

import (
	"gorm.io/gorm"
)

type Grade struct {
	StudentID string         `gorm:"column:student_id"`
	CourseID  string         `gorm:"column:course_id"`
	Grade     string         `gorm:"column:grade"`
	CreatedAt gorm.DeletedAt `gorm:"autoCreateTime"`
}
