package models

import (
	"time"
)

type WeeklyReport struct {
	ReportID   string `gorm:"column:report_id;not null"`  // 周报 ID
	ProjectID  string `gorm:"column:project_id;not null"` // 项目 ID
	WeekNumber int    `gorm:"column:week_number;not null"` // 周数
	Content    string `gorm:"column:content;not null"`    // 周报内容
	CreatedAt  time.Time `gorm:"column:PublishedAt;autoCreateTime"` // 默认当前时间
}

// TableName 指定表名
func (WeeklyReport) TableName() string {
	return "weekly_report"
}
