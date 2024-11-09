package models

import (
	"gorm.io/gorm"
)

type Institute struct {
	InstituteID string         `gorm:"column:institute_ID"`
	Name        string         `gorm:"column:institute_name"`
	CreatedAt   gorm.DeletedAt `gorm:"autoCreateTime"`
}
