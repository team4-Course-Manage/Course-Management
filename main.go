package main

import (
	"CMS\routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	routes.SetupRouter(r)
	r.Run(":8080")
}
