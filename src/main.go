package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var templates []Template = make([]Template, 0)

type Level struct {
	ID int `json:"level"`
	Approvers []int `json:"approvers"`
	Cond string `json:"cond"`
}

type Template struct {
	AppId int `json:"appId"`
	Levels []Level `json:"levels"`
}

func ListTemplates(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, templates)
}

func PostTemplate(c *gin.Context) {
	var template Template
	if err := c.BindJSON(&template); err != nil {
		return
	}
	
	templates = append(templates, template)
	c.IndentedJSON(http.StatusCreated, template)
}

func main() {
	r := gin.Default()
	r.GET("/templates", ListTemplates)
	r.POST("/templates", PostTemplate)

	r.Run()
}