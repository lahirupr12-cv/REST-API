package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/lahirupr12-cv/gin_restapi/config"
	"github.com/lahirupr12-cv/gin_restapi/models"
)

func UserController(c *gin.Context) {
	users := []models.User{}
	config.DB.Find(&users)
	c.JSON(200, &users)
}
