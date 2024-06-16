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

func (u User) GetName() string {
	return u.Username
}
func (u User) GetEmail() string {
	return u.Email
}
func (u User) GetID() int64 {
	return u.Id
}

func (u User) GetIsAdmin() bool {
	return u.IsAdmin
}

func (u User) GetIsActive() bool {
	return u.IsActive
}
