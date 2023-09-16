package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	fmt.Println("Home Loaded")
	c.HTML(http.StatusOK, "index.html", nil)
}
