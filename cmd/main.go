package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mefefirat/golang-rest-api/database"
	route "github.com/mefefirat/golang-rest-api/routes"
	"log"
)

func main() {

	router := Initialize()

	log.Fatal(router.Run(":8585"))
}

func Initialize() *gin.Engine {

	router := gin.Default()
	database.Connect()

	route.InitRoutes(router)

	return router
}
