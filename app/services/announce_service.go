package services

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
)

type Announcement struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

// 文件夹路径
const announcementsDir = "data/announcements"
const announcementsFile = "announcements.json"

var (
	mutex = &sync.Mutex{}
)

// AddAnnouncement 创建新公告并添加到文件
func AddAnnouncement(newAnnouncement Announcement) (Announcement, error) {
	mutex.Lock()
	defer mutex.Unlock()

	// 确保文件夹存在
	if err := ensureDirectory(announcementsDir); err != nil {
		return Announcement{}, err
	}

	announcements, err := loadAnnouncements()
	if err != nil {
		return Announcement{}, err
	}

	newAnnouncement.ID = len(announcements) + 1
	announcements = append(announcements, newAnnouncement)

	if err := saveAnnouncements(announcements); err != nil {
		return Announcement{}, err
	}

	return newAnnouncement, nil
}

// GetAnnouncements 获取所有公告
func GetAnnouncements() ([]Announcement, error) {
	mutex.Lock()
	defer mutex.Unlock()

	return loadAnnouncements()
}

// GetAnnouncementByID 根据 ID 获取公告详细内容
func GetAnnouncementByID(id int) (Announcement, error) {
	mutex.Lock()
	defer mutex.Unlock()

	announcements, err := loadAnnouncements()
	if err != nil {
		return Announcement{}, err
	}

	for _, a := range announcements {
		if a.ID == id {
			return a, nil
		}
	}

	return Announcement{}, errors.New("announcement not found")
}

// UpdateAnnouncement 更新公告
func UpdateAnnouncement(id int, updated Announcement) (Announcement, error) {
	mutex.Lock()
	defer mutex.Unlock()

	announcements, err := loadAnnouncements()
	if err != nil {
		return Announcement{}, err
	}

	for i, a := range announcements {
		if a.ID == id {
			announcements[i].Title = updated.Title
			announcements[i].Content = updated.Content

			if err := saveAnnouncements(announcements); err != nil {
				return Announcement{}, err
			}

			return announcements[i], nil
		}
	}

	return Announcement{}, errors.New("announcement not found")
}

// DeleteAnnouncement 删除公告
func DeleteAnnouncement(id int) error {
	mutex.Lock()
	defer mutex.Unlock()

	announcements, err := loadAnnouncements()
	if err != nil {
		return err
	}

	for i, a := range announcements {
		if a.ID == id {
			announcements = append(announcements[:i], announcements[i+1:]...)
			return saveAnnouncements(announcements)
		}
	}

	return errors.New("announcement not found")
}

// 内部函数：加载公告列表
func loadAnnouncements() ([]Announcement, error) {
	filePath := filepath.Join(announcementsDir, announcementsFile)

	if _, err := os.Stat(filePath); errors.Is(err, os.ErrNotExist) {
		return []Announcement{}, nil
	}

	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var announcements []Announcement
	if err := json.Unmarshal(data, &announcements); err != nil {
		return nil, err
	}

	return announcements, nil
}

// 内部函数：保存公告列表
func saveAnnouncements(announcements []Announcement) error {
	filePath := filepath.Join(announcementsDir, announcementsFile)

	data, err := json.Marshal(announcements)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filePath, data, 0644)
}

// 内部函数：确保目录存在
func ensureDirectory(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return os.MkdirAll(dir, 0755)
	}
	return nil
}
