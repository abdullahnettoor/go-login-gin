package models

import "time"

type Session struct {
	name   string
	expiry time.Time
}
