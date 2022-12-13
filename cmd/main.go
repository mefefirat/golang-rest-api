package main

import (
	"github.com/gin-gonic/gin"
	route "github.com/mefefirat/golang-rest-api/routes"
	"log"
)

func main() {

	router := Initialize()

	log.Fatal(router.Run(":8585"))
}

func Initialize() *gin.Engine {

	router := gin.Default()

	route.InitUserRoutes(router)

	return router
}
