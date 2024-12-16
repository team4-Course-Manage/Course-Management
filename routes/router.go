package routes

import (
	"Course-Management/app/controllers"
	"Course-Management/app/services"
	"Course-Management/config"
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
		loginService := services.NewLoginService(config.DB1, config.DB2)
		authController := controllers.NewAuthController(loginService)

		login.POST("/login", authController.LoginHandler)
		login.POST("/login/github", authController.LoginWithGithubHandler)
	}

	// 学生信息
	stuInfo := route.Group("/stuInfo")
	{
		stuInfoService := services.NewStuInfoService(config.DB1)
		stuInfoController := controllers.NewStuInfoController(stuInfoService)
		stuInfo.POST("/getStuInfoByName", stuInfoController.GetStudentInfoByName)
		stuInfo.POST("/getStuInfoByID", stuInfoController.GetStudentInfoByID)
	}

	// 公告
	announce := route.Group("/announce")
	{
		announceController := controllers.NewAnnouncementController()
		announce.POST("/post", announceController.PostAnnouncement)
		announce.POST("/addAnnouncement", announceController.PostAnnouncement)
		announce.POST("/getAnnouncements", announceController.GetAnnouncements)
		announce.POST("/getAnnouncementByID", announceController.GetAnnouncementByID)
		announce.POST("/updateAnnouncement", announceController.UpdateAnnouncement)
		announce.POST("/deleteAnnouncement", announceController.DeleteAnnouncement)

	}

	// 启动
	route.Run(":8080")
}
