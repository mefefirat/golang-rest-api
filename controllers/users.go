package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func listUser(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "{a: 1}")
}
