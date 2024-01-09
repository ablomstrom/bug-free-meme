package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Username string `form:"username"`
	UserId   string `form:"userid"`
}

var Users = []User{{"Goran", "123"}, {"Pelle", "111"}}

func main() {

	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "My App",
		})
	})

	r.GET("/u/:userId", GetUsetById)

	r.GET("/add-user", ShowAddUserForm)
	r.GET("/user-list", GetUserList)
	r.POST("/add-user", AddUser)
	r.Run()
}

func userExist(id string) (*User, error) {
	for _, user := range Users {
		if user.UserId == id {
			return &user, nil // Found the user
		}
	}
	return nil, errors.New("User not found")
	// User not found
}

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
