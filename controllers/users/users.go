package users

import (
	"github.com/gin-gonic/gin"
	user "github.com/mefefirat/golang-rest-api/models/users"
	"net/http"
)

func checkError(err error, c *gin.Context) {
	if err != nil {
		panic(err)
	}
}

func List(c *gin.Context) {

	users, err := user.List()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "users": users})

}

func Create(c *gin.Context) {

	username := c.PostForm("username")
	if username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Please enter username"})
		return
	}

	id, err := user.Create(username)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "insert_id": id})
}
