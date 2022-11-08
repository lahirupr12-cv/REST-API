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

	//database connection
	config.Connect()

	//config port
	router.Run(":8081")

}
