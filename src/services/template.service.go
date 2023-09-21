package services

import (
	"net/http"
	"sample-go-app/src/dto"

	"github.com/gin-gonic/gin"
)

var templates []dto.Template = make([]dto.Template, 0)

func ListTemplates(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, templates)
}

func PostTemplate(c *gin.Context) {
	var template dto.Template
	if err := c.BindJSON(&template); err != nil {
		return
	}
	
	templates = append(templates, template)
	c.IndentedJSON(http.StatusCreated, template)
}