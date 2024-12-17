package services

import (
	"course_management/app/models"
	"database/sql"
	"errors"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"gorm.io/gorm"
)

// ProjectMemberService 提供项目组成员的业务逻辑
type ProjectMemberService struct {
	DB *sql.DB
}

// 获取项目成员列表
func (pms *ProjectMemberService) GetProjectMembers(projectID string) ([]models.ProjectMember, error) {
	query := "SELECT project_id, member_id FROM project_member WHERE project_id = ?"

	rows, err := pms.DB.Query(query, projectID)
	if err != nil {
		return nil, errors.New("database query failed")
	}
	defer rows.Close()

	var members []models.ProjectMember
	for rows.Next() {
		var member models.ProjectMember
		if err := rows.Scan(&member.ProjectID, &member.MemberID); err != nil {
			return nil, errors.New("failed to parse query result")
		}
		members = append(members, member)
	}

	return members, nil
}

// 添加项目成员
func (pms *ProjectMemberService) AddProjectMember(projectID, memberID string) error {
	insertQuery := "INSERT INTO project_member (project_id, member_id) VALUES (?, ?)"
	_, err := pms.DB.Exec(insertQuery, projectID, memberID)
	if err != nil {
		return errors.New("failed to insert data into database")
	}
	return nil
}
