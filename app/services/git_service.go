package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type GitService struct {
	BaseURL string
	Token   string
}

func NewGitService() *GitService {
	return &GitService{
		BaseURL: os.Getenv("GITEA_BASE_URL"),
		Token:   os.Getenv("GITEA_API_TOKEN"),
	}
}

// 创建仓库
func (s *GitService) CreateRepository(repoName string) (string, error) {
	url := fmt.Sprintf("%s/api/v1/user/repos", s.BaseURL)
	data := map[string]interface{}{
		"name": repoName,
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "token "+s.Token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return "", fmt.Errorf("failed to create repository: %s", resp.Status)
	}

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	return result["html_url"].(string), nil
}

// 获取仓库信息
func (s *GitService) GetRepository(repoName string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/api/v1/repos/%s", s.BaseURL, repoName)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "token "+s.Token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get repository: %s", resp.Status)
	}

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result, nil
}

// 删除仓库
func (s *GitService) DeleteRepository(repoName string) error {
	url := fmt.Sprintf("%s/api/v1/repos/%s", s.BaseURL, repoName)

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "token "+s.Token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("failed to delete repository: %s", resp.Status)
	}

	return nil
}

// 列出用户的所有仓库
func (s *GitService) ListRepositories() ([]map[string]interface{}, error) {
	url := fmt.Sprintf("%s/api/v1/user/repos", s.BaseURL)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "token "+s.Token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to list repositories: %s", resp.Status)
	}

	var result []map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result, nil
}

// 添加协作者
func (s *GitService) AddCollaborator(repoName, collaborator string) error {
	url := fmt.Sprintf("%s/api/v1/repos/%s/collaborators/%s", s.BaseURL, repoName, collaborator)

	req, err := http.NewRequest("PUT", url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "token "+s.Token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("failed to add collaborator: %s", resp.Status)
	}

	return nil
}

// 获取仓库的提交记录
func (s *GitService) ListCommits(repoName string) ([]map[string]interface{}, error) {
	url := fmt.Sprintf("%s/api/v1/repos/%s/commits", s.BaseURL, repoName)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "token "+s.Token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to list commits: %s", resp.Status)
	}

	var result []map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result, nil
}
