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
	// 批量导入学生的路由
	importGroup := r.Group("/import")
	{
		// 创建导入服务
		importService := services.NewImportService(config.DB1)
		// 创建导入控制器
		importController := controllers.NewImportController(importService)

		// 注册批量导入学生的路由
		importGroup.POST("/students", importController.ImportStudents)
	}
	// 汇报顺序功能
	report := r.Group("/report")
	{
		reportService := services.NewReportService(config.DB1)
		reportController := controllers.NewReportController(reportService)

		report.GET("/weekly", reportController.GetWeeklyOrder)
		report.POST("/choose", reportController.ChooseOrder)
	}

	Project := r.Group("/project")
	{
		projectService := services.NewProjectService(config.DB1)
		projectController := controllers.NewProjectController(projectService)


		Project.POST("/create", projectController.CreateProject)
		Project.PUT("/update", projectController.UpdateProject)
		Project.GET("/list", projectController.ListProjects)
		Project.GET("/details", projectController.GetProjectDetails)
		Project.GET("/count", projectController.CountProjects)
	}


	// 启动
	route.Run(":8080")
}
