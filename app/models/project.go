package models

import "time"

// Project 模型
type Project struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Name        string    `gorm:"column:name;not null" json:"name"`
	Description string    `gorm:"column:description" json:"description"`
	Status      int       `gorm:"column:status;not null" json:"status"` // 使用整数表示状态
	CreatorID   string    `gorm:"column:creator_id;not null" json:"creator_id"`
	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"` // 自动设置创建时间
	UpdatedAt   time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"` // 自动设置更新时间
}

// TableName 自定义表名
func (Project) TableName() string {
	return "projects"
}