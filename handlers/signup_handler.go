package handlers

import (
	"fmt"
	"login-gin/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Signup(c *gin.Context) {
	// Clear Cache
	utils.ClearCache(c)

	// Check if user is already logged in
	userSession := GetSessionValue(c, "login-session", "username")
	if userSession != nil {
		// Redirect to Home
		fmt.Println("Redirected to Home")
		c.HTML(http.StatusSeeOther, "index.html", userSession)
		return
	}

	// Load Signup
	fmt.Println("Signup Loaded")
	c.HTML(http.StatusOK, "signup.html", nil)
}

func SignupPost(c *gin.Context) {
	// Clear Cache
	utils.ClearCache(c)

	// Check if user is already logged in
	userSession := GetSessionValue(c, "login-session", "username")
	if userSession != nil {
		// Redirect to Home
		fmt.Println("Redirected to Home")
		c.HTML(http.StatusSeeOther, "index.html", userSession)
		return
	}

	// Retrieve data from Html Form
	c.Request.ParseForm()
	username := c.Request.FormValue("name")
	email := c.Request.FormValue("email")
	pwd := c.Request.Form.Get("password")
	confirmPwd := c.Request.Form.Get("confirm-password")

	// Check Passwords
	if pwd != confirmPwd {
		c.HTML(http.StatusBadRequest, "signup.html", "Password must match")
		return
	}

	// Create User
	exist := CreateUser(username, email, pwd)
	if exist != nil {
		c.HTML(http.StatusBadRequest, "signup.html", exist)
		return
	}

	// Create Session and Login User
	user, _ := GetUser(email)
	err := CreateSession(c, "login-session", "username", user)
	if err != nil {
		c.Error(err)
	}
	// Load home
	c.HTML(http.StatusOK, "index.html", user)
	c.Redirect(http.StatusOK, "/")
}
