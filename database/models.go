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

type UpdateUser struct {
	ID               int
	Username         string
	NewPassword      string
	NewPasswordAgain string
	OldPassword      string
}

type Chat struct {
	ID       int
	Sender   int
	Receiver int
	body     string
}
