package handlers

import (
	"fmt"
	"login-gin/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Logout(c *gin.Context) {

	err := DestroySession(c, "login-session")
	if err != nil {
		fmt.Println("Error occured:", err)
		return
	}
	utils.ClearCache(c)
	c.HTML(http.StatusOK, "login.html", "Successfully Logged Out")
}
