package users

import (
	"github.com/gin-gonic/gin"
	model "github.com/mefefirat/golang-rest-api/models/users"
	"github.com/mefefirat/golang-rest-api/models/users/entry"
	"net/http"
)

func List(c *gin.Context) {

	users, err := model.List()

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

	email := c.PostForm("email")
	if email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Please enter email"})
		return
	}

	first_name := c.PostForm("first_name")
	if first_name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Please enter first name"})
		return
	}

	last_name := c.PostForm("last_name")
	if last_name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Please enter last name"})
		return
	}

	user := entry.User{
		UserName:  username,
		Email:     email,
		FirstName: first_name,
		LastName:  last_name,
	}

	id, err := model.Create(&user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "insert_id": id})
}
