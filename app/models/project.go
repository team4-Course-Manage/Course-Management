package models

import (
	"encoding/json"
	"errors"
	"time"
)

// ProjectStatus 定义项目状态的枚举类型
type ProjectStatus string

const (
	StatusAsPlanned ProjectStatus = "As_Planned"
	StatusAtRisk    ProjectStatus = "At_risk"
	StatusDeviated  ProjectStatus = "Deviated"
	StatusStopped   ProjectStatus = "Stopped"
	StatusNotBegun  ProjectStatus = "Not_Begun"
	StatusFinished  ProjectStatus = "Finished"
)

func (ps ProjectStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(ps))
}

func (ps *ProjectStatus) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case string(StatusAsPlanned), string(StatusAtRisk), string(StatusDeviated), string(StatusStopped), string(StatusNotBegun), string(StatusFinished):
		*ps = ProjectStatus(s)
		return nil
	default:
		return errors.New("invalid ProjectStatus")
	}
}

type Project struct {
	ID          string        `gorm:"column:projectID;not null" json:"ID"`
	Name        string        `gorm:"column:project_name;not null" json:"name"`
	Description string        `gorm:"column:description" json:"description"`
	Status      ProjectStatus `gorm:"column:Status;type:enum('As_Planned','At_risk','Deviated','Stopped','Not_Begun','Finished');not null" json:"status"` // 使用枚举类型
	CreatorID   string        `gorm:"column:creator_id;not null" json:"creator_id"`
	CreatedAt   time.Time     `gorm:"column:created_at;autoCreateTime" json:"created_at"` // 自动设置创建时间
	UpdatedAt   time.Time     `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"` // 自动设置更新时间
}

// TableName 自定义表名
func (Project) TableName() string {
	return "project"
}
