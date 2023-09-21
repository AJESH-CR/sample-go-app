package main

import (
	"sample-go-app/src/services"
	
	"github.com/gin-gonic/gin"
)


func main() {
	r := gin.Default()
	r.GET("/templates", services.ListTemplates)
	r.POST("/templates", services.PostTemplate)

	r.Run()
}