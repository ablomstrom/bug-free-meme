package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
	// userID := c.Param("userId")
	// user, err := userExist(userID)

	// // Early return if user dont exists
	// if err != nil {
	// 	c.HTML(http.StatusOK, "404.tmpl", gin.H{})
	// 	return
	// }

	// c.HTML(http.StatusOK, "user.tmpl", gin.H{
	// 	"username": user.Username,
	// })

}

func SignupHandler(c *gin.Context) {

	var newUser UserSignup
	// Checks request is valod json?
	if err := c.ShouldBind(&newUser); err != nil {
		c.HTML(http.StatusOK, "invalid_json.tmpl", gin.H{"Users": Users})
		return
	}

	user, _ := userExist(newUser.Username)

	if user != nil {
		c.HTML(http.StatusOK, "users_already_exists.tmpl", gin.H{"Username": user.Username})
		return
	}

	Users = append(Users, User{newUser.Username, newUser.Password, uuid.New().String()})

	c.HTML(http.StatusCreated, "user_added.tmpl", gin.H{"NewUserName": newUser.Username})
}

func LoginHandler(c *gin.Context) {

	var user UserSignup
	// Checks request is valod json?
	if err := c.ShouldBind(&user); err != nil {
		c.HTML(http.StatusOK, "invalid_json.tmpl", gin.H{"Users": Users})
		return
	}
	logged_in_user, err := LoginUser(user)

	if err != nil {

		c.HTML(http.StatusOK, "error.tmpl", gin.H{"Error": err.Error()})
		return
	}

	session, _ := store.Get(c.Request, "goran-session")
	session.Values["authenticated"] = true
	session.Values["username"] = logged_in_user.Username // Store the username in the session
	session.Save(c.Request, c.Writer)

	c.HTML(http.StatusOK, "user_profile.tmpl", gin.H{"Username": logged_in_user.Username})
}

func LoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "login_page.tmpl", gin.H{})
}

func DashboardHandler(c *gin.Context) {

	session, err := store.Get(c.Request, "goran-session")

	if err != nil {
		// Handle error - perhaps redirect to an error page or login page
		c.Redirect(http.StatusFound, "/auth/login")
		return
	}

	authenticated, ok := session.Values["authenticated"].(bool)

	if !ok || !authenticated {
		// User is not authenticated
		c.HTML(http.StatusOK, "dashboard.tmpl", gin.H{"LoggedIn": false})
	} else {

		username, _ := session.Values["username"]
		// User is authenticated
		c.HTML(http.StatusOK, "dashboard.tmpl", gin.H{"LoggedIn": true, "Username": username})

	}

}

func LogoutHandler(c *gin.Context) {
	session, err := store.Get(c.Request, "goran-session")
	if err != nil {
		// Optional: Handle error
	}
	session.Values["authenticated"] = false
	session.Values["username"] = "" // Clear username or other user-specific data
	session.Save(c.Request, c.Writer)

	// Redirect to home page or login page after logout
	c.Redirect(http.StatusFound, "/d")
}
