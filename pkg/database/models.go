package database

import (
	_ "github.com/mattn/go-sqlite3"
)

// User table struct.
type User struct {
	ID       int
	Username string
	Password string
}

// NewUser table struct
type NewUser struct {
	Username      string
	Password      string
	PasswordAgain string
}

// LoginUser table struct
type LoginUser struct {
	Username string
	Password string
}

// UpdateUser table struct
type UpdateUser struct {
	ID               int
	Username         string
	NewPassword      string
	NewPasswordAgain string
	OldPassword      string
}

// NewChat table struct
type NewChat struct {
	Sender   int
	Receiver int
	Body     string
}

// Chat table struct
type Chat struct {
	ID       int    `json:"id"`
	Sender   int    `json:"sender"`
	Receiver int    `json:"receiver"`
	Body     string `json:"body"`
}
