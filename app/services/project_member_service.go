package services

import (
	"Course-Management/app/models"
	"gorm.io/gorm"
)

type ProjectMemberService struct {
	DB *gorm.DB // 数据库实例
}

// NewProjectMemberService 构造函数
func NewProjectMemberService(db *gorm.DB) *ProjectMemberService {
	if db == nil {
		panic("Database connection must not be nil")
	}
	return &ProjectMemberService{DB: db}
}

// AddMember 向项目中添加成员
func (s *ProjectMemberService) AddMember(projectID, memberID string) error {
	// 创建项目成员关联
	projectMember := models.ProjectMember{
		ProjectID: projectID,
		MemberID:  memberID,
	}

	// 插入关联数据
	if err := s.DB.Create(&projectMember).Error; err != nil {
		return err
	}

	return nil
}


// GetMembers 获取项目的所有成员（仅返回成员 ID）
func (s *ProjectMemberService) GetMembers(projectID string) ([]string, error) {
	var memberIDs []string

	// 查询项目成员的 MemberID
	err := s.DB.Model(&models.ProjectMember{}).
		Select("member_id").
		Where("project_id = ?", projectID).
		Pluck("member_id", &memberIDs).Error

	if err != nil {
		return nil, err
	}

	return memberIDs, nil
}
