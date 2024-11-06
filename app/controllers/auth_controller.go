package controllers

import (
	"Course-Management/app/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthController struct {
	authService *services.AuthService
}

func NewAuthController(authService *services.AuthService) *AuthController {
	return &AuthController{authService: authService}
}

// @Summary 登录接口
// @Description 用户通过此接口登录系统并获取 token
// @Tags 用户
// @Accept json
// @Produce json
// @Param username body string true "用户名"
// @Param password body string true "密码"
// @Success 200 {object} Response "登录成功返回的响应"
// @Failure 401 {object} Response "认证失败返回的响应"
// @Router /login [post]
func (ctrl *AuthController) LoginHandler(ctx *gin.Context) {
	var loginRequest struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&loginRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := ctrl.authService.Login(loginRequest.Username, loginRequest.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"access_token": token,
		"message":      "Login successful",
	})
}
