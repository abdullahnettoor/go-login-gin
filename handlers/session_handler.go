package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
)

// `store` is a global variable that is used to store all sessions.
var store = sessions.NewCookieStore([]byte("1234567890"))

// `CreateSession` creates a new session.
func CreateSession(c *gin.Context, sessionName string, key string, value interface{}) error {
	// Get the session for the given session name.
	session, err := store.Get(c.Request, sessionName)
	if err != nil {
		return err
	}

	// Set the value of the given key in the session.
	session.Values[key] = value

	// Save the session.
	return session.Save(c.Request, c.Writer)
}

// `GetSessionValue` gets the value of a given key from the session.
func GetSessionValue(c *gin.Context, sessionName string, key string) interface{} {
	// Get the session for the given session name.
	session, _ := store.Get(c.Request, sessionName)

	// Return the value of the given key from the session.
	return session.Values[key]
}

// `DestroySession` destroys the session.
func DestroySession(c *gin.Context, sessionName string) error {
	// Get the session for the given session name.
	session, err := store.Get(c.Request, sessionName)
	if err != nil {
		return err
	}

	// Set the maximum age of the session to -1, which destroys the session.
	session.Options.MaxAge = -1
	// Save the session.
	return session.Save(c.Request, c.Writer)
}
