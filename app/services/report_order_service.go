package services

import (
	"Course-Management/app/models"
	"fmt"
	"gorm.io/gorm"
	"errors"
)

type ReportService struct {
	DB *gorm.DB
}

// NewReportService 构造函数
func NewReportService(db *gorm.DB) *ReportService {
	return &ReportService{
		DB: db,
	}
}

// InitializeReportOrders 初始化报告顺序，预插入10周，每周5个顺序
func (s *ReportService) InitializeReportOrders() error {
	// 先检查是否已存在数据，避免重复插入
	var existingOrders []models.ReportOrder
	if err := s.DB.Find(&existingOrders).Error; err != nil {
		return fmt.Errorf("检查现有报告顺序时出错: %v", err)
	}

	// 如果已有数据，则不再插入
	if len(existingOrders) > 0 {
		return nil
	}

	// 插入10周，每周5个顺序
	for week := 1; week <= 10; week++ {
		for order := 1; order <= 5; order++ {
			reportOrder := models.ReportOrder{
				Week:  uint(week),
				Order: uint(order),
			}
			if err := s.DB.Create(&reportOrder).Error; err != nil {
				return fmt.Errorf("插入报告顺序时出错: %v", err)
			}
		}
	}

	return nil
}

// 获取某一周的汇报顺序
func (s *ReportService) GetWeeklyOrder(week uint) ([]models.ReportOrder, error) {
	var orders []models.ReportOrder
	err := s.DB.Preload("Student").Where("week = ?", week).Find(&orders).Error
	if err != nil {
		return nil, fmt.Errorf("failed to fetch weekly orders: %v", err)
	}
	return orders, nil
}

// 学生选择汇报顺序
func (s *ReportService) ChooseOrder(week, order uint, studentID string) error {
	var existingOrder models.ReportOrder
	if err := s.DB.Where("week = ? AND `order` = ?", week, order).First(&existingOrder).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 如果不存在该顺序，则初始化
			existingOrder = models.ReportOrder{
				Week:  week,
				Order: order,
			}
		} else {
			return fmt.Errorf("failed to query order: %v", err)
		}
	}

	// 检查是否已被占用
	if existingOrder.StudentID != nil {
		return errors.New("order already selected")
	}

	// 分配学生到顺序
	existingOrder.StudentID = &studentID
	return s.DB.Save(&existingOrder).Error
}
