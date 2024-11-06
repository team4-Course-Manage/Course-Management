package models

import (
	"gorm.io/gorm"
)

type Enroll struct {
	EnrollID    string         `gorm:"column:enroll_ID"`
	StudentName string         `gorm:"column:student_name"`
	CourseID    string         `gorm:"column:course_id"`
	CreatedAt   gorm.DeletedAt `gorm:"autoCreateTime"`
}
