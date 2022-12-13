package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitUserRoutes(route *gin.Engine) {

	groupRoute := route.Group("/api/v1/users")
	groupRoute.GET("/list", listUser)
}

func listUser(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "{a: 1}")
}
