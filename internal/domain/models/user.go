package models

import "time"

type User struct {
	Id        int64
	Username  string
	Email     string
	IsActive  bool
	IsAdmin   bool
	CreatedAt time.Time
	Password  string
}
