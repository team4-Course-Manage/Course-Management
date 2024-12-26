package controllers

import (
	"Course-Management/app/models"
	"Course-Management/app/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GroupController struct {
	groupService *services.GroupService
}

func NewGroupController(groupService *services.GroupService) *GroupController {
	return &GroupController{
		groupService: groupService,
	}
}

// CreateGroupHandler 处理创建组的HTTP请求，并自动将创建者加入组
func (gc *GroupController) CreateGroupHandler(c *gin.Context) {
	var group models.Group
	if err := c.ShouldBindJSON(&group); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 从请求中获取创建组的学生ID和姓名，这里假设前端传入的JSON数据里包含这两个字段
	var studentInfo models.Student
	if err := c.ShouldBindJSON(&studentInfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 调用服务层方法创建组，并获取创建后的组信息以及可能出现的错误
	createdGroup, err := gc.groupService.CreateGroupAndJoin(group, studentInfo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "创建组失败"})
		return
	}

	c.JSON(http.StatusCreated, createdGroup)
}

// GetGroupListHandler 处理获取组列表的HTTP请求
func (gc *GroupController) GetGroupListHandler(c *gin.Context) {
	groups, err := gc.groupService.GetGroupList()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取组列表失败"})
		return
	}

	c.JSON(http.StatusOK, groups)
}

// JoinGroupHandler 处理申请加入组的HTTP请求
// group_controller.go
func (gc *GroupController) JoinGroupHandler(c *gin.Context) {
	groupId, err := strconv.Atoi(c.Query("groupId"))
	studentIdStr := c.Query("studentId") // 直接获取字符串形式的 studentId
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数错误"})
		return
	}

	err = gc.groupService.JoinGroup(groupId, studentIdStr)
	if err != nil {
		if err.Error() == "组已满员，无法加入" {
			c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "加入组失败"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "成功加入组"})
}

// GetGroupMembersHandler 处理根据组id返回组成员信息的HTTP请求
func (gc *GroupController) GetGroupMembersHandler(c *gin.Context) {
	groupId, err := strconv.Atoi(c.Param("groupId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的组id"})
		return
	}

	members, err := gc.groupService.GetGroupMembersByGroupId(groupId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取组成员信息失败"})
		return
	}

	c.JSON(http.StatusOK, members)
}
