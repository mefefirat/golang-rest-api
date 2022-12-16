package users

import (
	"github.com/gin-gonic/gin"
	model "github.com/mefefirat/golang-rest-api/models/users"
	"net/http"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func ListUser(c *gin.Context) {

	//w.Header().Set("Content-Type", "application/json")
	users, err := model.List()
	checkError(err)

	c.IndentedJSON(http.StatusOK, users)
}
