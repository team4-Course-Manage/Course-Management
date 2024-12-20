package main

import (
	"log"
	"os"

	"Course-Management/config"
	"Course-Management/routes"
	"Course_Management/app/controllers"
	"Course_Management/app/services"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// 加载环境变量
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	//连接数据库
	config.ConnectDatabase()
	// 初始化数据库
	db := config.InitDB()
	defer db.Close()
	//连接oauth服务
	config.LoadOAuthConfig()

	//初始化router
	r := gin.Default()

	projectMemberService := &services.ProjectMemberService{DB: db}
	projectMemberController := &controllers.ProjectMemberController{ProjectMemberService: projectMemberService}
	// 路由定义
	r.GET("/api/project/:project_id/members", projectMemberController.GetProjectMembers)
	r.POST("/api/project/:project_id/add_member", projectMemberController.AddProjectMember)
	routes.Init(r)

	// 开启服务
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Starting server on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
