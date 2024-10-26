package routes

import (
	"Course-Management/app/controllers"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	fmt.Println("路由已加载")
	// 初始化
	router := gin.Default()

	// 登录注册
	register := router.Group("/")
	{
		register.POST("/register", controllers.Register.PhoneRegister)
		register.POST("/register", controllers.Register.EmailRegister)
	}

	login := router.Group("/")
	{
		login.POST("/login", controllers.Login.Login)
		// login.GET("/captcha", controllers.Login.Captcha)
		//login.POST("/updatePwd", controllers.Login.UpdatePwd)
	}

	// 启动
	router.Run(":8080")
}
