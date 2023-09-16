package main

import (
	"encoding/gob"
	"fmt"
	"login-gin/handlers"
	"login-gin/models"

	"github.com/gin-gonic/gin"
)

var Users = map[string]models.User{}

func main() {
	gob.Register(models.User{})
	// Create Server
	r := gin.Default()

	// Load Html and other files
	r.LoadHTMLGlob("view/*.html")
	r.Static("/static", "./static")

	r.GET("/", handlers.Home)
	r.GET("/login", handlers.Login)
	r.POST("/login", handlers.LoginPost)
	r.GET("/signup", handlers.Signup)
	r.POST("/signup", handlers.SignupPost)
	r.POST("/logout", handlers.Logout)

	// Start Server
	fmt.Println("Server started: http://0.0.0.0:8080")
	r.Run(":8080")

}
