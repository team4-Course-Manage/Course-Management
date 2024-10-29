package services

import (
	"Course-Management/app/models"
	"errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

var db *gorm.DB

func InitDB() error {
	var err error
	dsn := "user:password@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	// 从环境变量中获取 DSN
	dsn = os.Getenv("DATABASE_DSN")
	if dsn == "" {
		return errors.New("数据库 DSN 未设置")
	}
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	err = db.AutoMigrate(&models.User{})
	return err
}

func init() {
	if err := InitDB(); err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
}

func LoginHandler(username, password string) (*models.User, error) {
	var user models.User
	result := db.Find(&user, "username = ? AND password = ?", username, password)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("无效的用户名或密码")
	}
	return &user, result.Error
}
