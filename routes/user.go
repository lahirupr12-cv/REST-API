package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lahirupr12-cv/gin_restapi/controller"
)

func UserRoutes(router *gin.Engine) {
	router.GET("/", controller.UserController)
}
