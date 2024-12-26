package controllers

import (
	"net/http"
	"strconv"

	"Course-Management/app/models"
	"Course-Management/app/services"

	"github.com/gin-gonic/gin"
)

type ProjectController struct {
	Service *services.ProjectService
}

// 构造函数
func NewProjectController(service *services.ProjectService) *ProjectController {
	return &ProjectController{Service: service}
}

// CreateProject 创建项目
func (c *ProjectController) CreateProject(ctx *gin.Context) {
	var input struct {
		Name        string               `json:"name" binding:"required"`
		Description string               `json:"description"`
		Status      models.ProjectStatus `json:"status" binding:"required"` // 确保是 models.ProjectStatus 类型
		CreatorID   string               `json:"creator_id" binding:"required"`
	}

	// 绑定输入数据
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 验证状态是否合法
	if !services.IsValidStatus(input.Status) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project status"})
		return
	}

	// 创建项目
	project := models.Project{
		ID:          input.CreatorID,
		Name:        input.Name,
		Description: input.Description,
		Status:      input.Status,
		CreatorID:   input.CreatorID,
	}

	if err := c.Service.CreateProject(&project); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create project"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Project created successfully", "project": project})
}

// UpdateProject 更新项目
func (c *ProjectController) UpdateProject(ctx *gin.Context) {
	var input struct {
		ID          uint                 `json:"id" binding:"required"`
		Name        string               `json:"name"`
		Description string               `json:"description"`
		Status      models.ProjectStatus `json:"status"`
	}

	// 绑定输入数据
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 验证状态是否合法
	if input.Status != models.ProjectStatus(0) && !services.IsValidStatus(input.Status) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project status"})
		return
	}

	// 准备更新数据
	updates := map[string]interface{}{}

	if input.Name != "" {
		updates["name"] = input.Name
	}
	if input.Description != "" {
		updates["description"] = input.Description
	}
	if input.Status != models.ProjectStatus(0) {
		updates["status"] = input.Status
	}

	project, err := c.Service.UpdateProject(input.ID, updates)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Project updated successfully", "project": project})
}

// ListProjects 获取项目列表
func (c *ProjectController) ListProjects(ctx *gin.Context) {
	creatorID := ctx.Query("creator_id")
	if creatorID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Creator ID is required"})
		return
	}

	projects, err := c.Service.ListProjects(creatorID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch projects"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"projects": projects})
}

// GetProjectDetails 获取项目详情
func (c *ProjectController) GetProjectDetails(ctx *gin.Context) {
	idStr := ctx.Query("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project ID"})
		return
	}

	project, err := c.Service.GetProjectDetails(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"project": project})
}

// CountProjects 获取创建者的项目数量
func (c *ProjectController) CountProjects(ctx *gin.Context) {
	creatorID := ctx.Query("creator_id")
	if creatorID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Creator ID is required"})
		return
	}

	count, err := c.Service.CountProjectsByCreator(creatorID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch project count"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"creator_id": creatorID, "project_count": count})
}
