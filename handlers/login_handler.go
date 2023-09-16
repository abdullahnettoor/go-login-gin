package handlers

import (
	"fmt"
	"login-gin/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	utils.ClearCache(c)

	userSession := GetSessionValue(c, "login-session", "username")
	if userSession != nil {
		fmt.Println("Redirected to Home")
		c.HTML(http.StatusSeeOther, "index.html", userSession)
		return
	}

	fmt.Println("Login Loaded")
	c.HTML(http.StatusOK, "login.html", nil)
}

func LoginPost(c *gin.Context) {
	utils.ClearCache(c)

	userSession := GetSessionValue(c, "login-session", "username")
	if userSession != nil {
		fmt.Println("Redirected to Home")
		c.HTML(http.StatusSeeOther, "index.html", userSession)
		return
	}

	c.Request.ParseForm()
	email := c.Request.FormValue("email")
	pwd := c.Request.Form.Get("password")

	user, exist := GetUser(email)
	if exist != nil {
		c.HTML(http.StatusBadRequest, "login.html", exist)
		return
	}

	if user.Email != email {
		c.HTML(http.StatusBadRequest, "login.html", "Invalid Email")
		return
	} else if user.Password != pwd {
		c.HTML(http.StatusBadRequest, "login.html", "Wrong Password")
		return
	}

	err := CreateSession(c, "login-session", "username", email)
	if err != nil {
		c.Error(err)
		return
	}
	c.HTML(http.StatusOK, "index.html", user)
	c.Redirect(http.StatusOK, "/")

	fmt.Println(c.Request)
}
