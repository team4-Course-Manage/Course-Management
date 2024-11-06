package models

import (
	"gorm.io/gorm"
)

type Courses struct {
	CourseID   string         `gorm:"column:course_ID"`
	CourseName string         `gorm:"column:course_name"`
	Credit     int            `gorm:"column:credit"`
	Institute  string         `gorm:"column:institute"`
	CreatedAt  gorm.DeletedAt `gorm:"autoCreateTime"`
}
