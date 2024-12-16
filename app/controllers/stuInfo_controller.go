package controllers

import (
	"net/http"
	"strconv"

	"Course-Management/app/services"

	"github.com/gin-gonic/gin"
)

type StuInfoController struct {
	StuInfoService *services.StuInfoService
}

// NewStuInfoController 初始化 StuInfoController
func NewStuInfoController(stuInfoService *services.StuInfoService) *StuInfoController {
	return &StuInfoController{
		StuInfoService: stuInfoService,
	}
}

// GetAllStudents 获取所有学生信息
func (c *StuInfoController) GetAllStudents(ctx *gin.Context) {
	students, err := c.StuInfoService.GetAllStudents()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"students": students})
}

// GetStudentInfoByID 通过 StudentID 获取学生信息
func (c *StuInfoController) GetStudentInfoByID(ctx *gin.Context) {
	studentIDStr := ctx.Query("student_id")
	if studentIDStr == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "student_id is required"})
		return
	}

	studentID, err := strconv.Atoi(studentIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid student_id"})
		return
	}

	studentInfo, err := c.StuInfoService.GetStudentByID(studentID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"student_info": studentInfo})
}

// GetStudentInfoByName 通过名字获取学生信息
func (c *StuInfoController) GetStudentInfoByName(ctx *gin.Context) {
	name := ctx.Query("name")
	if name == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "name is required"})
		return
	}

	studentInfo, err := c.StuInfoService.GetStudentByName(name)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"student_info": studentInfo})
}
