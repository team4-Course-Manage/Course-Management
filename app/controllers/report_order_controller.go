package controllers

import (
	"Course-Management/app/services"
	"net/http"
	"github.com/gin-gonic/gin"
	"strconv"
)

type ReportController struct {
	ReportService *services.ReportService
}

// NewReportController 构造函数
func NewReportController(reportService *services.ReportService) *ReportController {
	return &ReportController{
		ReportService: reportService,
	}
}

// 初始化报告顺序数据
func (c *ReportController) InitializeReportOrders(ctx *gin.Context) {
	err := c.ReportService.InitializeReportOrders()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "初始化报告顺序失败", "details": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "报告顺序初始化成功"})
}

// 获取某周的汇报顺序
func (c *ReportController) GetWeeklyOrder(ctx *gin.Context) {
	weekStr := ctx.Query("week")
	week, err := strconv.Atoi(weekStr)
	if err != nil || week <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid week parameter"})
		return
	}

	orders, err := c.ReportService.GetWeeklyOrder(uint(week))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch weekly orders", "details": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, orders)
}

// 学生选择汇报顺序
func (c *ReportController) ChooseOrder(ctx *gin.Context) {
	var req struct {
		Week      uint   `json:"week" binding:"required"`
		Order     uint   `json:"order" binding:"required"`
		StudentID string `json:"student_id" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data", "details": err.Error()})
		return
	}

	err := c.ReportService.ChooseOrder(req.Week, req.Order, req.StudentID)
	if err != nil {
		ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Order selected successfully"})
}
