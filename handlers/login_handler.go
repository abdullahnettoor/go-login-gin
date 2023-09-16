package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {

	fmt.Println(c.Request)

	fmt.Println("Login Loaded")
	c.HTML(http.StatusOK, "login.html", nil)
}

func LoginPost(c *gin.Context) {
	fmt.Println(c.Request)

}
