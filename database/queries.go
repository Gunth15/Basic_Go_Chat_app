package database

import (
	"database/sql"
	"errors"

	"golang.org/x/crypto/bcrypt"

	_ "github.com/mattn/go-sqlite3"
)

// gets user from database.
// produces error if the user does not exist or invalid credentials.
func get_user(db *sql.DB, id uint) (User, error) {
	var user User

	stmnt, err := db.Prepare("SELECT id, username, password FROM Users WHERE id=?;")
	if err != nil {
		return user, err
	}

	defer stmnt.Close()

	err = stmnt.QueryRow(id).Scan(&user.ID, &user.Username, &user.Password)

	return user, err
}

// Adds new user to databse
// Throws error if passwords do not match
func new_user(db *sql.DB, new_user NewUser) (User, error) {
	var user User

	if new_user.Password != new_user.PasswordAgain {
		return user, errors.New("Passwords do not match")
	}

	// Salt password returns errr if password is too short or too long to hash
	hashed_passwrd, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return user, err
	}

	stmnt, err := db.Prepare("INSERT INTO Users(Username,Password) VALUES (?,?) RETURNING id;")
	if err != nil {
		return user, err
	}

	defer stmnt.Close()

	err = stmnt.QueryRow(&new_user.Username, &new_user.Password).Scan(&user.ID)
	user.Password = string(hashed_passwrd)

	return user, nil
}

// updates the users login infromation
// checks if user is valid aswell
func update_user(db *sql.DB, update_user *UpdateUser) (User, error) {
	var user User

	stmnt, err := db.Prepare("SELECT username, password, id FROM Users WHERE id=?;")
	if err != nil {
		return user, err
	}

	defer stmnt.Close()

	err = stmnt.QueryRow(&update_user.ID).Scan(&user.Username, &user.Password, &user.ID)
	if err != nil {
		return user, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(update_user.OldPassword))
	if err != nil {
		return user, err
	}

	query := "UPDATE User SET "

	if user.Username != update_user.Username {
		query += ("username = " + update_user.Username)
	}
	if user.Password != update_user.OldPassword {
		query += ("password =  " + update_user.NewPassword)
	}

	query += "WHERE id = ? RETURNING *;"

	stmnt, err = db.Prepare(query)
	if err != nil {
		return user, err
	}

	defer stmnt.Close()

	err = stmnt.QueryRow().Scan(&user.Username, &user.Password, &user.ID)
	return user, err
}

// Gets user by username and password
func login_user() User

// Gets all the users frrom database
func get_all_users() []User

// Gets all chats for a combination of the sender's id and the receiver's id.
func get_chats() []Chat

// creates a new chat
func new_chat() Chat
