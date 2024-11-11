package models

import (
	"time"
)

// Resource 表示资源发布的模型
type Resource struct {
	ResID       uint      `gorm:"column:RES_ID;primaryKey;autoIncrement"`
	CourseID    string    `gorm:"column:Course_ID"`
	Title       string    `gorm:"column:Title;not null"`
	Description string    `gorm:"column:Description"`
	URL         string    `gorm:"column:URL;not null"`
	PublisherID string    `gorm:"column:PublisherID;not null"`
	PublishedAt time.Time `gorm:"column:PublishedAt;autoCreateTime"` // 默认当前时间
}
