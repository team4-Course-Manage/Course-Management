package services

import (
	"Course-Management/app/models"
	"errors"

	"gorm.io/gorm"
)

// 定义有效的项目状态
var validStatuses = []models.ProjectStatus{
	models.StatusPlanned,
	models.StatusInProgress,
	models.StatusCompleted,
	models.StatusOnHold,
	models.StatusCancelled,
}

// IsValidStatus 验证状态是否合法
func IsValidStatus(status models.ProjectStatus) bool {
	for _, valid := range validStatuses {
		if status == valid {
			return true
		}
	}
	return false
}

type ProjectService struct {
	DB *gorm.DB // 数据库实例
}

// NewProjectService 构造函数，初始化 ProjectService
func NewProjectService(db *gorm.DB) *ProjectService {
	if db == nil {
		panic("Database connection must not be nil")
	}
	return &ProjectService{DB: db}
}

// CreateProject 创建新项目
func (s *ProjectService) CreateProject(project *models.Project) error {
	if !IsValidStatus(project.Status) {
		return errors.New("invalid project status")
	}
	if err := s.DB.Create(project).Error; err != nil {
		return err
	}
	return nil
}

// UpdateProject 更新项目内容
func (s *ProjectService) UpdateProject(id uint, updates map[string]interface{}) (models.Project, error) {
	var project models.Project

	// 查找项目
	if err := s.DB.First(&project, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.Project{}, errors.New("project not found")
		}
		return models.Project{}, err
	}

	// 验证状态更新
	if status, ok := updates["status"].(models.ProjectStatus); ok {
		if !IsValidStatus(status) {
			return models.Project{}, errors.New("invalid project status")
		}
	}

	// 更新项目
	if err := s.DB.Model(&project).Updates(updates).Error; err != nil {
		return models.Project{}, err
	}

	return project, nil
}

// ListProjects 获取创建者的项目列表
func (s *ProjectService) ListProjects(creatorID string) ([]models.Project, error) {
	var projects []models.Project

	// 查询项目列表
	if err := s.DB.Select("id, name").Where("creator_id = ?", creatorID).Find(&projects).Error; err != nil {
		return nil, err
	}

	return projects, nil
}

// GetProjectDetails 获取项目详情
func (s *ProjectService) GetProjectDetails(id uint) (models.Project, error) {
	var project models.Project

	// 查询项目详情
	if err := s.DB.First(&project, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.Project{}, errors.New("project not found")
		}
		return models.Project{}, err
	}

	return project, nil
}

// CountProjectsByCreator 返回创建者的项目数量
func (s *ProjectService) CountProjectsByCreator(creatorID string) (int64, error) {
	var count int64
	if err := s.DB.Model(&models.Project{}).Where("creator_id = ?", creatorID).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
