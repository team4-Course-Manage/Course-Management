package controllers

import (
	"Course-Management/app/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	LoginService *services.LoginService
}

func NewAuthController(loginService *services.LoginService) *AuthController {
	return &AuthController{
		LoginService: loginService,
	}
}

// 使用 UserID 和 Password 登录
func (ac *AuthController) LoginHandler(c *gin.Context) {
	var credentials struct {
		UserID   string `json:"userID"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	loginResult, err := ac.LoginService.Login(credentials.UserID, credentials.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if loginResult.Success {
		c.JSON(http.StatusOK, gin.H{
			"message": "Login successful",
			"role":    loginResult.Role,
		})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
	}
}

// 使用 Gitea 登录
func (c *AuthController) LoginWithGiteaHandler(ctx *gin.Context) {
	var giteaRequest struct {
		Code string `json:"code" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&giteaRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data", "details": err.Error()})
		return
	}

	userInfo, err := c.LoginService.LoginWithGitea(giteaRequest.Code)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, userInfo)
}

// 回调处理
func (ac *AuthController) CallbackHandler(c *gin.Context) {
	ac.LoginService.HandleCallback(c.Writer, c.Request)
}
