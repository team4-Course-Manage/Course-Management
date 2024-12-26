package services

import (
	"Course-Management/app/models"

	"gorm.io/gorm"
)

type TaskService struct {
	DB *gorm.DB
}

func NewTaskService(db *gorm.DB) *TaskService {
	return &TaskService{
		DB: db,
	}
}

// SaveTask 保存任务方法
func (ts *TaskService) SaveTask(task models.Task) (models.Task, error) {
	result := ts.DB.Create(&task)
	if result.Error != nil {
		return models.Task{}, result.Error
	}
	return task, nil
}

// GetTaskById 根据任务ID获取任务方法
func (ts *TaskService) GetTaskById(taskId int) (models.Task, error) {
	var task models.Task
	result := ts.DB.Where("task_id =?", taskId).First(&task)
	return task, result.Error
}

// GetTasksByProjectName 根据项目名称查询任务列表方法
func (ts *TaskService) GetTasksByProjectName(projectName string) ([]models.Task, error) {
	var tasks []models.Task
	result := ts.DB.Where("project_name =?", projectName).Find(&tasks)
	return tasks, result.Error
}

// UpdateTask 修改任务信息方法
func (ts *TaskService) UpdateTask(task models.Task) (models.Task, error) {
	var existingTask models.Task
	result := ts.DB.Where("task_id =?", task.TaskId).First(&existingTask)
	if result.Error == nil {
		existingTask.TaskName = task.TaskName
		existingTask.ProjectId = task.ProjectId
		existingTask.Subject = task.Subject
		existingTask.TaskDescription = task.TaskDescription
		existingTask.ReceiverId = task.ReceiverId
		existingTask.PrincipalId = task.PrincipalId
		existingTask.Priority = task.Priority
		existingTask.TaskDate = task.TaskDate
		existingTask.TaskStatus = task.TaskStatus

		result = ts.DB.Save(&existingTask)
		if result.Error != nil {
			return models.Task{}, result.Error
		}
		return existingTask, nil
	}
	return models.Task{}, gorm.ErrRecordNotFound
}

// DeleteTask 删除任务方法
func (ts *TaskService) DeleteTask(taskId int) error {
	result := ts.DB.Delete(&models.Task{}, taskId)
	return result.Error
}
