package database

import (
	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	ID       int
	Username string
	Password string
}

type NewUser struct {
	Username      string
	Password      string
	PasswordAgain string
}

type LoginUser struct {
	Username string
	Password string
}

type UpdateUser struct {
	ID               int
	Username         string
	NewPassword      string
	NewPasswordAgain string
	OldPassword      string
}

type NewChat struct {
	Sender   int
	Receiver int
	Body     string
}

type Chat struct {
	ID       int    `json:"id"`
	Sender   int    `json:"sender"`
	Receiver int    `json:"receiver"`
	Body     string `json:"body"`
}
