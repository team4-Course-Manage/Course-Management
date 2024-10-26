package controllers

import (
	"Course-Management/app/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type RegisterController struct{}

func (ctrl RegisterController) PhoneRegister(c *gin.Context) {
	phone := c.PostForm("phone")
	code := c.PostForm("code")

	if !services.IsValidPhone(phone) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid phone number"})
		return
	}

	if !services.VerifyPhoneCode(phone, code) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid verification code"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Phone registered successfully"})
}

func (ctrl RegisterController) EmailRegister(c *gin.Context) {
	email := c.PostForm("email")
	code := c.PostForm("code")

	if !services.IsValidEmail(email) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email address"})
		return
	}

	if !services.VerifyEmailCode(email, code) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid verification code"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Email registered successfully"})
}

var Register = RegisterController{}
