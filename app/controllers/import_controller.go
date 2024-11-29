package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"Course-Management/app/services"
)

type ImportController struct {
	ImportService *services.ImportService // 注入服务
}
// NewImportController 是 ImportController 的构造函数
func NewImportController(importService *services.ImportService) *ImportController {
	return &ImportController{
		ImportService: importService, // 注入 ImportService
	}
}

// 批量导入学生的接口
func (c *ImportController) ImportStudents(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse file"})
		return
	}

	// 调用服务处理文件
	count, err := c.ImportService.ImportStudentsFromFile(file)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to import students", "details": err.Error()})
		return
	}

	// 返回成功结果
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Students imported successfully",
		"count":   count,
	})
}
