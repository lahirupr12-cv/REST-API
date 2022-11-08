package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lahirupr12-cv/gin_restapi/config"
	"github.com/lahirupr12-cv/gin_restapi/routes"
)

func main() {

	//configure router
	router := gin.New()
	routes.UserRoutes(router)
	router.Run(":8081")

	//database connection
	config.Connect()
}
