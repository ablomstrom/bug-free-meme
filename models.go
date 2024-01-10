package main

type User struct {
	Username string `form:"username"`
	Password string `form:"password"`
	UserId   string
}

type UserSignup struct {
	Username string `form:"username"`
	Password string `form:"password"`
}
