package routes

import (
	"Course-Management/app/controllers"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	fmt.Println("路由已加载")
	// 初始化
	route := gin.Default()

	// 登录
	login := route.Group("/")
	{
		authController := new(controllers.AuthController)
		login.POST("/login", authController.LoginHandler)
	}

	// 启动
	route.Run(":8080")
}
