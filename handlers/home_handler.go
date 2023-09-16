package handlers

import (
	"fmt"
	"login-gin/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	utils.ClearCache(c)

	userSession := GetSessionValue(c, "login-session", "username")
	if userSession != nil {
		fmt.Println("Home Loaded", userSession)
		c.HTML(http.StatusSeeOther, "index.html", userSession)
		return
	}

	fmt.Println("Redirected to Login")
	c.HTML(http.StatusOK, "login.html", nil)

}
