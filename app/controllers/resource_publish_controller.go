package controllers

import (
	"Course-Management/app/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ResourceController struct {
	ResourceService *services.ResourceService
}

// NewResourceController 创建一个新的 ResourceController 实例
func NewResourceController(resourceService *services.ResourceService) *ResourceController {
	return &ResourceController{
		ResourceService: resourceService,
	}
}

// PublishResource 处理资源发布请求
func (ctrl *ResourceController) PublishResource(c *gin.Context) {
	var input struct {
		CourseID    string `json:"course_id" binding:"required"`
		Title       string `json:"title" binding:"required"`
		Description string `json:"description"`
		URL         string `json:"url" binding:"required"`
		PublisherID string `json:"publisher_id" binding:"required"`
	}

	// 绑定请求体
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 调用服务发布资源
	resource, err := ctrl.ResourceService.PublishResource(input.CourseID, input.Title, input.Description, input.URL, input.PublisherID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Resource published successfully",
		"resource": resource,
	})
}

// GetResourceByID 根据资源ID查询资源
func (ctrl *ResourceController) GetResourceByID(c *gin.Context) {
	resID := c.Param("res_id")

	// 调用服务获取资源
	resource, err := ctrl.ResourceService.GetResourceByID(resID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"resource": resource,
	})
}

// GetResourcesByCourse 根据课程ID查询资源
func (ctrl *ResourceController) GetResourcesByCourse(c *gin.Context) {
	courseID := c.Param("course_id")

	// 调用服务获取资源
	resources, err := ctrl.ResourceService.GetResourcesByCourse(courseID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"resources": resources,
	})
}
