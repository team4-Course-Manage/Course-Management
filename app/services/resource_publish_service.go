package services

import (
	"Course-Management/app/models"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

// ResourceService 管理资源发布相关的操作
type ResourceService struct {
	DB *gorm.DB
}

// NewResourceService 创建一个新的 ResourceService 实例
func NewResourceService(db *gorm.DB) *ResourceService {
	return &ResourceService{
		DB: db,
	}
}

// PublishResource 处理资源的发布逻辑
func (s *ResourceService) PublishResource(courseID, title, description, url, publisherID string) (*models.Resource, error) {
	if s.DB == nil {
		return nil, errors.New("DB is not initialized")
	}

	// 校验资源是否已存在（根据课程ID和标题）
	var existingResource models.Resource
	if err := s.DB.Where("course_ID = ? AND title = ?", courseID, title).First(&existingResource).Error; err == nil {
		return nil, errors.New("resource with this title already exists for the course")
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	// 创建新的资源
	resource := models.Resource{
		CourseID:    courseID,
		Title:       title,
		Description: description,
		URL:         url,
		PublisherID: publisherID,
	}

	if err := s.DB.Create(&resource).Error; err != nil {
		return nil, fmt.Errorf("failed to publish resource: %w", err)
	}

	return &resource, nil
}

// GetResourceByID 根据资源ID查询单个资源
func (s *ResourceService) GetResourceByID(resID string) (*models.Resource, error) {
	if s.DB == nil {
		return nil, errors.New("DB is not initialized")
	}

	var resource models.Resource
	if err := s.DB.First(&resource, resID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("resource not found")
		}
		return nil, fmt.Errorf("failed to fetch resource: %w", err)
	}

	return &resource, nil
}

// GetResourcesByCourse 根据课程ID获取所有资源
func (s *ResourceService) GetResourcesByCourse(courseID string) ([]models.Resource, error) {
	if s.DB == nil {
		return nil, errors.New("DB is not initialized")
	}

	var resources []models.Resource
	if err := s.DB.Where("course_ID = ?", courseID).Find(&resources).Error; err != nil {
		return nil, fmt.Errorf("failed to fetch resources for course: %w", err)
	}

	return resources, nil
}
