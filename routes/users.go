package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/mefefirat/golang-rest-api/controllers"
)

func InitUserRoutes(route *gin.Engine) {

	groupRoute := route.Group("/api/v1/users")
	groupRoute.GET("/list", controller.listUser)
}
