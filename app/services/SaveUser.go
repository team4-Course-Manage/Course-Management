package services

import (
	"Course-Management/app/models"
)

// 保存用户信息到数据库
func SaveUser(user *models.User) error {
	// 使用 GORM 保存用户
	return db.Create(user).Error
}
