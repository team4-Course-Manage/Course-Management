package services

import (
	"errors"
	"fmt"
	"mime/multipart"
	"strings"

	"github.com/xuri/excelize/v2"

	"Course-Management/app/models"

	"gorm.io/gorm"
)



type ImportService struct {
	DB *gorm.DB
}
// NewImportService 是 ImportService 的构造函数
func NewImportService(db *gorm.DB) *ImportService {
	return &ImportService{
		DB: db, // 注入数据库连接
	}
}
// 从 Excel 文件中导入学生
func (s *ImportService) ImportStudentsFromFile(file *multipart.FileHeader) (int, error) {
	// 打开文件
	f, err := file.Open()
	if err != nil {
		return 0, fmt.Errorf("failed to open file: %v", err)
	}
	defer f.Close()

	// 解析 Excel 文件
	excel, err := excelize.OpenReader(f)
	if err != nil {
		return 0, fmt.Errorf("failed to read Excel file: %v", err)
	}

	// 获取第一张表
	sheetName := excel.GetSheetName(0)
	rows, err := excel.GetRows(sheetName)
	if err != nil || len(rows) == 0 {
		return 0, errors.New("invalid Excel file format")
	}

	// 逐行解析
	var students []models.Student
	for _, row := range rows[1:] { // 跳过表头
		if len(row) < 2 {
			continue // 跳过无效行
		}

		studentID := strings.TrimSpace(row[0])
		name := strings.TrimSpace(row[1])
		institute := ""
		if len(row) > 2 {
			institute = strings.TrimSpace(row[2])
		}

		// 默认密码为明文
		password := "123456"

		// 添加到学生列表
		students = append(students, models.Student{
			StudentID: studentID,
			Name:      name,
			Institute: institute,
			Password:  password,
		})
	}

	// 插入数据库
	if err := s.DB.Create(&students).Error; err != nil {
		return 0, fmt.Errorf("failed to insert students: %v", err)
	}

	return len(students), nil
}
