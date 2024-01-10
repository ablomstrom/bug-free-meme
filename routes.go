package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func setupRoutes(router *gin.Engine) {
	router.LoadHTMLGlob("templates/*")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "My App",
		})
	})

	router.GET("/u/:userId", GetUsetById)

	router.GET("/add-user", ShowAddUserForm)
	router.GET("/user-list", GetUserList)
	router.GET("/d", DashboardHandler)

	user := router.Group("/u")
	{
		user.POST("/", SignupHandler)
	}

	api := router.Group("/auth")
	{
		api.POST("/login", LoginHandler)
		api.POST("/logout", LogoutHandler)
		api.GET("/logout", LogoutHandler)

		api.GET("/login", LoginPage)
	}
}
