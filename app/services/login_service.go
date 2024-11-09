package services

import (
	"Course-Management/app/models"
	"Course-Management/config"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"gorm.io/gorm"
)

type LoginService struct {
	DB1 *gorm.DB
	DB2 *gorm.DB
}

func NewLoginService(db1, db2 *gorm.DB) *LoginService {
	return &LoginService{
		DB1: db1,
		DB2: db2,
	}
}

type UserInfo struct {
	AccessToken string
	Name        string
	Institute   string
}

func (s *LoginService) VerifyCredentials(userID, password string) (UserInfo, error) {
	if s.DB1 == nil {
		return UserInfo{}, errors.New("DB1 is not initialized")
	}
	if s.DB2 == nil {
		return UserInfo{}, errors.New("DB2 is not initialized")
	}
	var student models.Student
	var teacher models.Teacher

	// 检查 student 表
	if err := s.DB1.Where("student_id = ? AND password = ?", userID, password).First(&student).Error; err == nil {
		return UserInfo{Name: student.Name, Institute: student.Institute}, nil
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return UserInfo{}, err
	}

	// 检查 teacher 表
	if err := s.DB1.Where("teacher_id = ? AND password = ?", userID, password).First(&teacher).Error; err == nil {
		return UserInfo{Name: teacher.Name, Institute: teacher.Institute}, nil
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return UserInfo{}, err
	}

	return UserInfo{}, errors.New("invalid user ID or password")
}

func (s *LoginService) Login(userID, password string) (UserInfo, error) {
	userInfo, err := s.VerifyCredentials(userID, password)
	if err != nil {
		return UserInfo{}, fmt.Errorf("failed to verify credentials: %w", err)
	}

	payload := map[string]string{
		"grant_type":    "password",
		"username":      userID,
		"password":      password,
		"client_id":     config.OAuthSettings.ClientID,
		"client_secret": config.OAuthSettings.ClientSecret,
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return UserInfo{}, fmt.Errorf("failed to marshal request payload: %w", err)
	}

	resp, err := http.Post(config.OAuthSettings.TokenURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return UserInfo{}, fmt.Errorf("failed to contact OAuth server: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return UserInfo{}, errors.New("authentication failed at OAuth server")
	}

	var tokenResponse struct {
		AccessToken string `json:"access_token"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&tokenResponse); err != nil {
		return UserInfo{}, fmt.Errorf("failed to parse token response: %w", err)
	}

	userInfo.AccessToken = tokenResponse.AccessToken
	return userInfo, nil
}
