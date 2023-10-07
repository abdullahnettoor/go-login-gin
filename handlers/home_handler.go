package handlers

import (
	"fmt"
	"login-gin/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	// Clear cache of browser
	utils.ClearCache(c)

	// Check if user logged in
	userSession := GetSessionValue(c, "login-session", "username")
	if userSession != nil {
		// Load Home
		c.HTML(http.StatusSeeOther, "index.html", userSession)
		return
	}

	// Redirect to Login
	fmt.Println("Redirected to Login")
	c.HTML(http.StatusOK, "login.html", nil)

}
