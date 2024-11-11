package controllers

import (
	"Course-Management/app/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AnnouncementController struct{}

func NewAnnouncementController() *AnnouncementController {
	return &AnnouncementController{}
}

func (ac *AnnouncementController) PostAnnouncement(c *gin.Context) {
	var newAnnouncement services.Announcement
	if err := c.ShouldBindJSON(&newAnnouncement); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdAnnouncement, err := services.AddAnnouncement(newAnnouncement)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save announcement"})
		return
	}

	c.JSON(http.StatusCreated, createdAnnouncement)
}

func (ac *AnnouncementController) GetAnnouncements(c *gin.Context) {
	announcements, err := services.GetAnnouncements()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve announcements"})
		return
	}

	c.JSON(http.StatusOK, announcements)
}

func (ac *AnnouncementController) GetAnnouncementByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	announcement, err := services.GetAnnouncementByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Announcement not found"})
		return
	}

	c.JSON(http.StatusOK, announcement)
}

func (ac *AnnouncementController) UpdateAnnouncement(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var updatedAnnouncement services.Announcement
	if err := c.ShouldBindJSON(&updatedAnnouncement); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	announcement, err := services.UpdateAnnouncement(id, updatedAnnouncement)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Announcement not found"})
		return
	}

	c.JSON(http.StatusOK, announcement)
}

func (ac *AnnouncementController) DeleteAnnouncement(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := services.DeleteAnnouncement(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Announcement not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Announcement deleted"})
}
