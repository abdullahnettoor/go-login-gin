package handlers

import (
	"fmt"
	"login-gin/models"
	// "github.com/gin-gonic/gin"
)

var Users = map[string]models.User{}

func GetUser(key string) (u models.User, err any) {
	u, exists := Users[key]
	if !exists {
		err = "User don't exist"
		return u, err
	}
	err = nil
	return u, err
}

func CreateUser(name, email, pwd string) (err any) {
	_, exists := Users[email]
	if exists {
		fmt.Println("User already exists:", Users[email])
		return "User already exists"
	}
	Users[email] = models.User{
		Name:     name,
		Email:    email,
		Password: pwd,
	}
	fmt.Println("User Created:", Users[email])
	fmt.Println(Users)
	return nil
}
