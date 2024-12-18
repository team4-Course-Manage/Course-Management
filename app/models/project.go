package models
import (
	"time"
)
type Project struct {
	ProjectID   string `gorm:"column:project_id;not null"`   // 项目 ID
	Name        string `gorm:"column:name;not null"`        // 项目名称
	Description string `gorm:"column:description"`          // 项目描述
	StartDate   time.Time `gorm:"column:PublishedAt;autoCreateTime"` // 默认当前时间
	EndDate     string `gorm:"column:end_date"`             // 结束日期
}

// TableName 指定表名
func (Project) TableName() string {
	return "project"
}
