package main

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
)

var Users = []User{{"Goran", "asd132", "8acb57cf-c8a3-4dcb-8ea6-41b87a3b5651"}, {"Pelle", "hunter2", "9851bcea-3b50-4fa7-a7cf-c626ca94cc3d "}}

var store = sessions.NewCookieStore([]byte("goranssecret324234dosijnfsij3489hfn349n3erfiouwernfioedjnv93e4rn9onfoguijndgodeng9odfj"))

func main() {

	r := gin.Default()
	setupRoutes(r)
	r.Run()
}

func userExist(name string) (*User, error) {
	for _, user := range Users {
		if user.Username == name {
			return &user, nil // Found the user
		}
	}
	return nil, errors.New("User not found")
	// User not found
}

func LoginUser(user UserSignup) (*User, error) {
	for _, user_try := range Users {
		if user_try.Username == user.Username && user_try.Password == user.Password {
			return &user_try, nil // Found the user
		}
	}
	return nil, errors.New("Wrong username or password")
	// User not found
}
