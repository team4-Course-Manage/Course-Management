package config

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB1 *gorm.DB
	DB2 *gorm.DB
)

func ConnectDatabase() {
	// 业务数据库
	dsn1 := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB1_USER"),
		os.Getenv("DB1_PASSWORD"),
		os.Getenv("DB1_HOST"),
		os.Getenv("DB1_PORT"),
		os.Getenv("DB1_NAME"),
	)
	var err error
	DB1, err = gorm.Open(mysql.Open(dsn1), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database 1: %v", err)
	} else {
		log.Println("Database 1 connection established")
	}

	// OAuth 认证数据库
	dsn2 := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB2_USER"),
		os.Getenv("DB2_PASSWORD"),
		os.Getenv("DB2_HOST"),
		os.Getenv("DB2_PORT"),
		os.Getenv("DB2_NAME"),
	)
	DB2, err = gorm.Open(mysql.Open(dsn2), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database 2: %v", err)
	} else {
		log.Println("Database 2 connection established")
	}
}

func AutoMigrate(models ...interface{}) {
	err := DB1.AutoMigrate(models...)
	if err != nil {
		log.Fatalf("Failed to auto-migrate database 1: %v", err)
	} else {
		log.Println("Database 1 auto-migration completed")
	}

	err = DB2.AutoMigrate(models...)
	if err != nil {
		log.Fatalf("Failed to auto-migrate database 2: %v", err)
	} else {
		log.Println("Database 2 auto-migration completed")
	}
}
