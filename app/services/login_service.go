package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"

	"Course-Management/app/models"

	"gorm.io/gorm"
)

type LoginService struct {
	DB1 *gorm.DB
	DB2 *gorm.DB
}

func NewLoginService(db1, db2 *gorm.DB) *LoginService {
	if db1 == nil || db2 == nil {
		panic("Database connections (DB1, DB2) must not be nil")
	}
	return &LoginService{DB1: db1, DB2: db2}
}

type UserInfo struct {
	AccessToken string `json:"access_token,omitempty"`
	Name        string `json:"name"`
	Institute   string `json:"institute"`
}

// 登录（使用 UserID 和 Password）
func (s *LoginService) Login(userID, password string) (UserInfo, error) {
	// 确保数据库连接已经初始化
	if s.DB1 == nil {
		return UserInfo{}, errors.New("DB1 is not initialized")
	}

	var student models.Student
	var teacher models.Teacher

	// 检查 Student 表
	if err := s.DB1.Where("student_id = ? AND password = ?", userID, password).First(&student).Error; err == nil {
		return UserInfo{Name: student.Name, Institute: student.Institute}, nil
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return UserInfo{}, fmt.Errorf("database query error: %v", err)
	}

	// 检查 Teacher 表
	if err := s.DB1.Where("teacher_id = ? AND password = ?", userID, password).First(&teacher).Error; err == nil {
		return UserInfo{Name: teacher.Name, Institute: teacher.Institute}, nil
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return UserInfo{}, fmt.Errorf("database query error: %v", err)
	}

	// 如果用户信息不存在
	return UserInfo{}, errors.New("invalid user ID or password")
}

// 使用 GitHub OAuth 登录
func (s *LoginService) LoginWithGithub(code string) (UserInfo, error) {
	clientID := os.Getenv("OAUTH_CLIENT_ID")
	clientSecret := os.Getenv("OAUTH_CLIENT_SECRET")

	// 确保环境变量已加载
	if clientID == "" || clientSecret == "" {
		return UserInfo{}, errors.New("client ID or secret not configured")
	}

	// 通过 GitHub 获取 access token
	tokenURL := "https://github.com/login/oauth/access_token"
	payload := map[string]string{
		"client_id":     clientID,
		"client_secret": clientSecret,
		"code":          code,
	}
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return UserInfo{}, fmt.Errorf("failed to marshal JSON: %w", err)
	}

	req, err := http.NewRequest("POST", tokenURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return UserInfo{}, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return UserInfo{}, fmt.Errorf("failed to send request to GitHub: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return UserInfo{}, fmt.Errorf("GitHub responded with status: %d", resp.StatusCode)
	}

	var tokenResponse struct {
		AccessToken string `json:"access_token"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&tokenResponse); err != nil {
		return UserInfo{}, fmt.Errorf("failed to parse token response: %w", err)
	}

	if tokenResponse.AccessToken == "" {
		return UserInfo{}, errors.New("received empty access token")
	}

	// 使用 access token 获取用户信息
	userInfo, err := s.getGithubUserInfo(tokenResponse.AccessToken)
	if err != nil {
		return UserInfo{}, fmt.Errorf("failed to fetch user info: %w", err)
	}

	return userInfo, nil
}

// 通过 Access Token 获取 GitHub 用户信息
func (s *LoginService) getGithubUserInfo(accessToken string) (UserInfo, error) {
	userURL := "https://api.github.com/user"
	req, err := http.NewRequest("GET", userURL, nil)
	if err != nil {
		return UserInfo{}, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+accessToken)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return UserInfo{}, fmt.Errorf("failed to send request to GitHub: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return UserInfo{}, fmt.Errorf("GitHub responded with status: %d", resp.StatusCode)
	}

	var githubUser struct {
		Name    string `json:"name"`
		Company string `json:"company"`
		Email   string `json:"email"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&githubUser); err != nil {
		return UserInfo{}, fmt.Errorf("failed to parse user info: %w", err)
	}

	return UserInfo{
		Name:      githubUser.Name,
		Institute: githubUser.Company,
	}, nil
}
