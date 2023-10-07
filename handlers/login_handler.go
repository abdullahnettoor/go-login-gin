package handlers

import (
	"fmt"
	"login-gin/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	// Clear cache of browser
	utils.ClearCache(c)

	// Check if user already logged in
	userSession := GetSessionValue(c, "login-session", "username")
	if userSession != nil {
		// Redirect to Home
		fmt.Println("Redirected to Home")
		c.HTML(http.StatusSeeOther, "index.html", userSession)
		return
	}

	// Load Login
	fmt.Println("Login Loaded")
	c.HTML(http.StatusOK, "login.html", nil)
}

func LoginPost(c *gin.Context) {
	// Clear Cache of browser
	utils.ClearCache(c)

	// Check if user already logged in
	userSession := GetSessionValue(c, "login-session", "username")
	if userSession != nil {
		// Redirect to Home
		fmt.Println("Redirected to Home")
		c.HTML(http.StatusSeeOther, "index.html", userSession)
		return
	}

	// Retrieve data from Login Form
	c.Request.ParseForm()
	email := c.Request.FormValue("email")
	pwd := c.Request.Form.Get("password")

	// Check if user exists
	user, exist := GetUser(email)
	if exist != nil {
		// if not, Redirect to login
		c.HTML(http.StatusBadRequest, "login.html", exist)
		return
	}

	// Check if entered email and password is valid
	if user.Email != email {
		c.HTML(http.StatusBadRequest, "login.html", "Invalid Email")
		return
	} else if user.Password != pwd {
		c.HTML(http.StatusBadRequest, "login.html", "Wrong Password")
		return
	}

	// Create Session for the user
	err := CreateSession(c, "login-session", "username", email)
	if err != nil {
		c.Error(err)
		return
	}
	// Load Home
	c.HTML(http.StatusOK, "index.html", user)
	c.Redirect(http.StatusOK, "/")
}
