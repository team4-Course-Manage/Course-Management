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
func (c *AuthController) LoginHandler(ctx *gin.Context) {
	var loginRequest struct {
		UserID   string `json:"user_id" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&loginRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data", "details": err.Error()})
		return
	}

	userInfo, err := c.LoginService.Login(loginRequest.UserID, loginRequest.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, userInfo)
}

// 使用 GitHub 登录
func (c *AuthController) LoginWithGithubHandler(ctx *gin.Context) {
	var githubRequest struct {
		Code string `json:"code" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&githubRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data", "details": err.Error()})
		return
	}

	userInfo, err := c.LoginService.LoginWithGithub(githubRequest.Code)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, userInfo)
}
