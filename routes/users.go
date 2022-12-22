package routes

import (
	"github.com/gin-gonic/gin"
	userController "github.com/mefefirat/golang-rest-api/controllers/users"
)

func InitRoutes(route *gin.Engine) {

	groupRoute := route.Group("/api/v1/user")
	groupRoute.GET("/list", userController.List)
	groupRoute.POST("/add", userController.Create)
}
