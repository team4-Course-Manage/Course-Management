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

func (c *AuthController) LoginHandler(ctx *gin.Context) {
	var loginRequest struct {
		UserID   string `json:"user_id" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&loginRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	token, err := c.LoginService.Login(loginRequest.UserID, loginRequest.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"access_token": token})
}
