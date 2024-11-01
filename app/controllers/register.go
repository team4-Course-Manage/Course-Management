package controllers

import (
	"Course-Management/app/models"
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

	// 创建用户对象
	user := &models.User{Phone: phone}

	// 生成 TOTP
	secret, qrCode, err := services.GenerateTOTP(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate TOTP key"})
		return
	}

	// 将 TOTP 密钥保存到用户对象
	user.TOTPSecret = secret

	// 保存用户信息到数据库
	if err := services.SaveUser(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Phone registered successfully",
		"secret":  secret,
		"qrCode":  qrCode,
	})
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

	// 创建用户对象
	user := &models.User{Email: email}

	// 生成 TOTP
	secret, qrCode, err := services.GenerateTOTP(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate TOTP key"})
		return
	}

	// 将 TOTP 密钥保存到用户对象
	user.TOTPSecret = secret

	// 保存用户信息到数据库
	if err := services.SaveUser(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Email registered successfully",
		"secret":  secret,
		"qrCode":  qrCode,
	})
}

var Register = RegisterController{}
