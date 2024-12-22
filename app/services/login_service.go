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

type LoginResult struct {
	Success bool   `json:"success"`
	Role    string `json:"role,omitempty"`
}

// 登录（使用 UserID 和 Password）
func (s *LoginService) Login(userID, password string) (*LoginResult, error) {
	// 确保数据库连接已经初始化
	if s.DB1 == nil {
		return nil, errors.New("DB1 is not initialized")
	}

	var student models.Student
	var teacher models.Teacher

	// 检查 Student 表
	if err := s.DB1.Where("student_id = ? AND password = ?", userID, password).First(&student).Error; err == nil {
		return &LoginResult{
			Success: true,
			Role:    "student",
		}, nil
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("database query error: %v", err)
	}

	// 检查 Teacher 表
	if err := s.DB1.Where("teacher_id = ? AND password = ?", userID, password).First(&teacher).Error; err == nil {
		return &LoginResult{
			Success: true,
			Role:    "teacher",
		}, nil
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("database query error: %v", err)
	}

	// 如果用户信息不存在
	return &LoginResult{
		Success: false,
	}, nil
}

// 使用 Gitea OAuth 登录
func (s *LoginService) LoginWithGitea(code string) (bool, error) {
	clientID := os.Getenv("GITEA_CLIENT_ID")
	clientSecret := os.Getenv("GITEA_CLIENT_SECRET")
	baseURL := os.Getenv("GITEA_BASE_URL")

	// 确保环境变量已加载
	if clientID == "" || clientSecret == "" || baseURL == "" {
		return false, errors.New("client ID, secret, or base URL not configured")
	}

	// 通过 Gitea 获取 access token
	tokenURL := fmt.Sprintf("%s/login/oauth/access_token", baseURL)
	payload := map[string]string{
		"client_id":     clientID,
		"client_secret": clientSecret,
		"code":          code,
	}
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return false, fmt.Errorf("failed to marshal JSON: %w", err)
	}

	req, err := http.NewRequest("POST", tokenURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return false, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return false, fmt.Errorf("failed to send request to Gitea: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return false, fmt.Errorf("Gitea responded with status: %d", resp.StatusCode)
	}

	var tokenResponse struct {
		AccessToken string `json:"access_token"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&tokenResponse); err != nil {
		return false, fmt.Errorf("failed to parse token response: %w", err)
	}

	if tokenResponse.AccessToken == "" {
		return false, errors.New("received empty access token")
	}

	// 使用 access token 获取用户信息
	userInfo, err := s.getGiteaUserInfo(tokenResponse.AccessToken, baseURL)
	if err != nil {
		return false, fmt.Errorf("failed to fetch user info: %w", err)
	}

	// 处理 userInfo 的逻辑
	fmt.Printf("Logged in with Gitea user: %+v\n", userInfo)

	return true, nil
}

// 通过 Access Token 获取 Gitea 用户信息
func (s *LoginService) getGiteaUserInfo(accessToken, baseURL string) (UserInfo, error) {
	userURL := fmt.Sprintf("%s/api/v1/user", baseURL)
	req, err := http.NewRequest("GET", userURL, nil)
	if err != nil {
		return UserInfo{}, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Authorization", "token "+accessToken)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return UserInfo{}, fmt.Errorf("failed to send request to Gitea: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return UserInfo{}, fmt.Errorf("Gitea responded with status: %d", resp.StatusCode)
	}

	var giteaUser struct {
		Username string `json:"username"`
		FullName string `json:"full_name"`
		Company  string `json:"company"`
		Email    string `json:"email"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&giteaUser); err != nil {
		return UserInfo{}, fmt.Errorf("failed to parse user info: %w", err)
	}

	return UserInfo{
		Name:      giteaUser.FullName,
		Institute: giteaUser.Company,
	}, nil
}

// 处理回调逻辑
func (s *LoginService) HandleCallback(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var reqBody struct {
		Code string `json:"code"`
	}
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	success, err := s.LoginWithGitea(reqBody.Code)
	if err != nil {
		http.Error(w, fmt.Sprintf("Login failed: %v", err), http.StatusInternalServerError)
		return
	}

	response := map[string]bool{
		"success": success,
	}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
