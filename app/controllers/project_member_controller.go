package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"Course-Management/app/services"
)

type ProjectMemberController struct {
	Service *services.ProjectMemberService
}

// NewProjectMemberController 构造函数
func NewProjectMemberController(service *services.ProjectMemberService) *ProjectMemberController {
	return &ProjectMemberController{Service: service}
}

// AddMember 添加成员到项目
func (c *ProjectMemberController) AddMember(ctx *gin.Context) {
	var input struct {
		ProjectID string `json:"project_id" binding:"required"`
		MemberID  string `json:"member_id" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.Service.AddMember(input.ProjectID, input.MemberID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Member added successfully"})
}

// GetMembers 获取项目成员列表
func (c *ProjectMemberController) GetMembers(ctx *gin.Context) {
	projectID := ctx.Param("project_id")
	if projectID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Project ID is required"})
		return
	}

	memberIDs, err := c.Service.GetMembers(projectID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch project members"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"members": memberIDs})
}
