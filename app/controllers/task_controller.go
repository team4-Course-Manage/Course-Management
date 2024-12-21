package controllers

import (
	"Course-Management/app/services"
	"net/http"

	"Course-Management/app/models" // 引入models包，假设正确的路径是这样
	"github.com/gin-gonic/gin"
)

type TaskController struct {
	taskService services.TaskService
}

func NewTaskController(taskService services.TaskService) *TaskController {
	return &TaskController{
		taskService: taskService,
	}
}

func (tc *TaskController) CreateTask(c *gin.Context) {
	var task models.Task // 将这里的类型修改为models.Task，符合期望的任务数据结构类型
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	savedTask, err := tc.taskService.SaveTask(task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save task"})
		return
	}

	c.JSON(http.StatusCreated, savedTask)
}
