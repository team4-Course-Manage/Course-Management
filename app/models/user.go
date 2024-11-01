package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"uniqueIndex;size:255"`
	Password string `gorm:"size:255"`
	Email    string `gorm:"uniqueIndex;size:255"`
	Phone    string `gorm:"uniqueIndex;size:255"`
	TOTPSecret string `gorm:"size:255"` //存储TOTP密钥
}
