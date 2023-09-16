package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Signup(c *gin.Context) {
	fmt.Println("Signup Loaded")
	c.HTML(http.StatusOK, "signup.html", nil)
}

func SignupPost(c *gin.Context) {

}
