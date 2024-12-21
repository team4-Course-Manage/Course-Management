package services

import (
	"Course-Management/app/models"
	"errors"

	"gorm.io/gorm"
)

type GroupService struct {
	DB *gorm.DB
}

func NewGroupService(db *gorm.DB) *GroupService {
	return &GroupService{
		DB: db,
	}
}

// CreateGroup 创建组的方法
func (gs *GroupService) CreateGroup(group models.Group) (models.Group, error) {
	result := gs.DB.Create(&group)
	if result.Error != nil {
		return models.Group{}, result.Error
	}
	return group, nil
}

// GetGroupList 获取组列表的方法
func (gs *GroupService) GetGroupList() ([]models.Group, error) {
	var groups []models.Group
	result := gs.DB.Find(&groups)
	return groups, result.Error
}

// JoinGroup 申请加入组的方法，判断人数是否满员等逻辑
func (gs *GroupService) JoinGroup(groupId int, studentId int) error {
	var group models.Group
	result := gs.DB.Where("group_id =?", groupId).First(&group)
	if result.Error != nil {
		return result.Error
	}

	if group.CurrentMembers >= group.MaxMembers {
		return errors.New("组已满员，无法加入")
	}

	// 这里假设存在Student结构体和对应的方法来更新学生表中的所属组id，此处仅示意更新逻辑
	// 比如有个UpdateStudentGroup方法用于更新学生所属组信息，需根据实际情况完善
	err := gs.updateStudentGroup(studentId, groupId)
	if err != nil {
		return err
	}

	group.CurrentMembers++
	result = gs.DB.Save(&group)
	return result.Error
}

// GetGroupMembersByGroupId 根据组id返回组成员信息的方法
func (gs *GroupService) GetGroupMembersByGroupId(groupId int) ([]string, error) {
	// 这里假设通过关联查询等方式从数据库获取组内成员的学生姓名列表
	// 实际可能需要根据数据库表结构和关联关系来完善具体的查询语句
	// 以下是简单示意，返回空列表和nil表示无错误，需根据实际情况调整
	return []string{}, nil
}

func (gs *GroupService) updateStudentGroup(studentId int, groupId int) error {
	// 这里假设存在Student结构体，使用gorm更新学生表中对应学生的所属组id字段
	// 例如：
	// var student models.Student
	// result := gs.DB.Model(&student).Where("student_id =?", studentId).Update("group_id", groupId)
	// return result.Error
	return nil
}

func (gs *GroupService) CreateGroupAndJoin(group models.Group, student models.Student) (*models.Group, error) {
	// 先创建组
	createdGroup, err := gs.CreateGroup(group)
	if err != nil {
		return nil, err
	}
	// 将创建组的学生加入该组，更新学生表中对应的组ID字段
	err = gs.updateStudentGroup(student.StudentID, group.GroupId)
	if err != nil {
		return nil, err
	}
	return &createdGroup, nil
}
