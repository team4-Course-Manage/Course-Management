package main

import (
	"log"
	"os"

	"Course-Management/config"
	"Course-Management/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// 加载环境变量
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// 连接数据库
	config.ConnectDatabase()
	// 初始化数据库
	// db := config.InitDB()
	// defer db.Close()
	// 连接 OAuth 服务
	config.LoadOAuthConfig()

	// 初始化 router
	r := gin.Default()

	// 配置 CORS 中间件
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:5173"}
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	corsConfig.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization"}
	r.Use(cors.New(corsConfig))

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
