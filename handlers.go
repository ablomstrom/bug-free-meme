package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShowAddUserForm(c *gin.Context) {
	c.HTML(http.StatusOK, "add_user.tmpl", gin.H{"Users": Users})
}

func GetUserList(c *gin.Context) {
	c.HTML(http.StatusOK, "users.tmpl", gin.H{"Users": Users})
}

func AddUser(c *gin.Context) {
	var newUser User
	if err := c.ShouldBind(&newUser); err != nil {
		c.HTML(http.StatusBadRequest, "error.tmpl", gin.H{"error": err.Error()})
		return
	}

	Users = append(Users, newUser)

	c.HTML(http.StatusOK, "users.tmpl", gin.H{"Users": Users})

}

func GetUsetById(c *gin.Context) {
	userID := c.Param("userId")
	user, err := userExist(userID)

	// Early return if user dont exists
	if err != nil {
		c.HTML(http.StatusOK, "404.tmpl", gin.H{})
		return
	}

	c.HTML(http.StatusOK, "user.tmpl", gin.H{
		"username": user.Username,
	})

}
