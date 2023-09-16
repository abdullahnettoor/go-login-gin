package main

import (
	"fmt"
	// "login-gin/handlers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello World")

	r := gin.Default()
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"greet": "Welcome Friend",
		})
	})

	// handlers.Login()

	r.Run(":8080")

}
