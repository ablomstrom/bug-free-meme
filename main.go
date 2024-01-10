package main

import (
	"errors"

	"github.com/gin-gonic/gin"
)

var Users = []User{{"Goran", "123"}, {"Pelle", "111"}}

func main() {

	r := gin.Default()
	setupRoutes(r)
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
