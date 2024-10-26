package main

import (
	"Course-Management/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 注册路由
	routes.Init(r)

	// 启动服务器
	r.Run(":8080")
}
