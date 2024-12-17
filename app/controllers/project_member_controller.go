package controllers

import (
	"course_management/app/models"
	"course_management/app/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ProjectMemberController 控制器，处理项目组成员相关的 HTTP 请求
type ProjectMemberController struct {
	ProjectMemberService *services.ProjectMemberService
}

// 获取项目成员列表
func (pmc *ProjectMemberController) GetProjectMembers(c *gin.Context) {
	projectID := c.Param("project_id")

	// 调用服务层
	members, err := pmc.ProjectMemberService.GetProjectMembers(projectID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"project_id": projectID,
		"members":    members,
	})
}

// 添加项目成员
func (pmc *ProjectMemberController) AddProjectMember(c *gin.Context) {
	projectID := c.Param("project_id")

	// 解析请求体
	var requestBody struct {
		MemberID string `json:"member_id"`
	}
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// 调用服务层
	if err := pmc.ProjectMemberService.AddProjectMember(projectID, requestBody.MemberID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":    "Member added successfully",
		"project_id": projectID,
		"member_id":  requestBody.MemberID,
	})
}
