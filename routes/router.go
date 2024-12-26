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

	// 登录
	login := r.Group("/")
	{
		loginService := services.NewLoginService(config.DB1, config.DB2)
		authController := controllers.NewAuthController(loginService)

		login.POST("/login", authController.LoginHandler)
		login.POST("/login/gitea", authController.LoginWithGiteaHandler)
		login.GET("/callback", authController.CallbackHandler)
	}

	// 批量导入学生的路由
	importGroup := r.Group("/import")
	{
		importService := services.NewImportService(config.DB1)
		importController := controllers.NewImportController(importService)

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

	// 项目相关路由
	project := r.Group("/project")
	{
		projectService := services.NewProjectService(config.DB1)
		projectController := controllers.NewProjectController(projectService)

		project.POST("/create", projectController.CreateProject)
		project.PUT("/update", projectController.UpdateProject)
		project.GET("/list", projectController.ListProjects)
		project.GET("/details", projectController.GetProjectDetails)
		project.GET("/count", projectController.CountProjects)

		// 项目成员相关路由
		projectMemberService := services.NewProjectMemberService(config.DB1)
		projectMemberController := controllers.NewProjectMemberController(projectMemberService)

		project.POST("/:project_id/add_member", projectMemberController.AddMember)
		project.GET("/:project_id/get_members", projectMemberController.GetMembers)
	}

	//小组相关路由
	group := r.Group("/group")
	{
		groupService := services.NewGroupService(config.DB1)
		groupController := controllers.NewGroupController(groupService)

		group.POST("/add_group", groupController.CreateGroupHandler)
		group.GET("/get_group", groupController.GetGroupListHandler)
		group.POST("/join_group", groupController.JoinGroupHandler)
		group.GET("/check_message", groupController.GetGroupMembersHandler)
	}
	// 学生信息
	stuInfo := r.Group("/stuInfo")
	{
		stuInfoService := services.NewStuInfoService(config.DB1)
		stuInfoController := controllers.NewStuInfoController(stuInfoService)
		stuInfo.GET("/getAllStudents", stuInfoController.GetAllStudents)
		stuInfo.GET("/getStuInfoByName", stuInfoController.GetStudentInfoByName)
		stuInfo.GET("/getStuInfoByID", stuInfoController.GetStudentInfoByID)
	}

	// 公告
	announce := r.Group("/announcement")
	{
		announceController := controllers.NewAnnouncementController()
		announce.POST("/post", announceController.PostAnnouncement)
		announce.POST("/get", announceController.GetAnnouncements)
		announce.POST("/getbyid/:id", announceController.GetAnnouncementByID)
		announce.POST("/update/:id", announceController.UpdateAnnouncement)
		announce.POST("/delete/:id", announceController.DeleteAnnouncement)
	}

	// Git 服务
	gitService := services.NewGitService()
	gitController := controllers.NewGitController(gitService)
	gitGroup := r.Group("/git")
	{
		gitGroup.POST("/createRepo", gitController.CreateRepository)
		gitGroup.GET("/repos/:repo_name", gitController.GetRepository)
		gitGroup.DELETE("/repos/:repo_name", gitController.DeleteRepository)
		gitGroup.GET("/repos", gitController.ListRepositories)
		gitGroup.PUT("/repos/:repo_name/collaborators", gitController.AddCollaborator)
		gitGroup.GET("/repos/:repo_name/commits", gitController.ListCommits)
	}
}
