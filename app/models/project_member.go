package models

import (
	"time"
)

type ProjectMember struct {
	ProjectID string `gorm:"column:project_id;not null"` // 项目 ID
	MemberID  string `gorm:"column:member_id;not null"` // 成员 ID
}

// TableName 指定表名
func (ProjectMember) TableName() string {
	return "project_member"
}
