package models

import (
	"encoding/json"
	"errors"
	"time"
)

// ProjectStatus 定义项目状态的枚举类型
type ProjectStatus string

const (
	StatusPlanned    ProjectStatus = "Planned"
	StatusInProgress ProjectStatus = "InProgress"
	StatusCompleted  ProjectStatus = "Completed"
	StatusOnHold     ProjectStatus = "OnHold"
	StatusCancelled  ProjectStatus = "Cancelled"
)

// MarshalJSON 实现 json.Marshaler 接口
func (ps ProjectStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(ps))
}

// UnmarshalJSON 实现 json.Unmarshaler 接口
func (ps *ProjectStatus) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case string(StatusPlanned), string(StatusInProgress), string(StatusCompleted), string(StatusOnHold), string(StatusCancelled):
		*ps = ProjectStatus(s)
		return nil
	default:
		return errors.New("invalid ProjectStatus")
	}
}

// Project 模型
type Project struct {
	ID          string        `gorm:"column:projectID;not null" json:"ID"`
	Name        string        `gorm:"column:project_name;not null" json:"name"`
	Description string        `gorm:"column:description" json:"description"`
	Status      ProjectStatus `gorm:"column:Status;type:enum('Planned','InProgress','Completed','OnHold','Cancelled');not null" json:"status"` // 使用枚举类型
	CreatorID   string        `gorm:"column:creator_id;not null" json:"creator_id"`
	CreatedAt   time.Time     `gorm:"column:created_at;autoCreateTime" json:"created_at"` // 自动设置创建时间
	UpdatedAt   time.Time     `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"` // 自动设置更新时间
}

// TableName 自定义表名
func (Project) TableName() string {
	return "projects"
}
