// Course-Management/app/services/stuInfo_service.go
package services

import (
	"errors"
	"fmt"

	"Course-Management/app/models"
	"gorm.io/gorm"
)

type StuInfoService struct {
	DB1 *gorm.DB
}

func NewStuInfoService(db1 *gorm.DB) *StuInfoService {
	return &StuInfoService{DB1: db1}
}

// GetStudentByID 根据 StudentID 获取学生信息
func (s *StuInfoService) GetStudentByID(studentID int) (*models.Student, error) {
	// 确保数据库连接已经初始化
	if s.DB1 == nil {
		return nil, errors.New("DB1 is not initialized")
	}

	var student models.Student

	// 检查 Student 表
	if err := s.DB1.Where("student_id = ?", studentID).First(&student).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("student not found")
		}
		return nil, fmt.Errorf("database query error: %v", err)
	}

	return &student, nil
}

// GetStudentByName 根据 Name 获取学生信息列表
func (s *StuInfoService) GetStudentByName(name string) ([]*models.Student, error) {
	// 确保数据库连接已经初始化
	if s.DB1 == nil {
		return nil, errors.New("DB1 is not initialized")
	}

	var students []*models.Student

	// 检查 Student 表
	if err := s.DB1.Where("name = ?", name).Find(&students).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("students not found")
		}
		return nil, fmt.Errorf("database query error: %v", err)
	}

	return students, nil
}
