package handlers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("1234567890"))

func CreateSession(c *gin.Context, sessionName string, key string, value interface{}) error {
	session, err := store.Get(c.Request, sessionName)
	if err != nil {
		return err
	}
	session.Values[key] = value
	return session.Save(c.Request, c.Writer)
}

func GetSessionValue(c *gin.Context, sessionName string, key string) interface{} {
	session, _ := store.Get(c.Request, sessionName)
	fmt.Println(session)
	return session.Values[key]
}

func DestroySession(c *gin.Context, sessionName string) error {
	session, err := store.Get(c.Request, sessionName)
	if err != nil {
		return err
	}
	session.Options.MaxAge = -1
	return session.Save(c.Request, c.Writer)
}
