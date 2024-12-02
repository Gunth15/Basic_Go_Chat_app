package database

import (
	"database/sql"
	"errors"

	"golang.org/x/crypto/bcrypt"

	_ "github.com/mattn/go-sqlite3"
)

// gets user from database.
// produces error if the user does not exist or invalid credentials.
func Get_user(db *sql.DB, id uint) (User, error) {
	var user User

	stmnt, err := db.Prepare("SELECT id, username, password FROM Users WHERE id=?;")
	if err != nil {
		return user, err
	}

	defer stmnt.Close()

	err = stmnt.QueryRow(&id).Scan(&user.ID, &user.Username, &user.Password)

	return user, err
}

// Adds new user to databse
// Throws error if passwords do not match
func New_user(db *sql.DB, new_user NewUser) (User, error) {
	var user User

	if new_user.Password != new_user.PasswordAgain {
		return user, errors.New("Passwords do not match given %s and returned %s")
	}

	// Salt password returns errr if password is too short or too long to hash
	hashed_passwrd, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return user, err
	}

	user.Password = string(hashed_passwrd)

	stmnt, err := db.Prepare("INSERT INTO Users(Username,Password) VALUES (?,?) RETURNING id;")
	if err != nil {
		return user, err
	}

	defer stmnt.Close()

	if err = stmnt.QueryRow(&new_user.Username, &user.Password).Scan(&user.ID); err != nil {
		return user, err
	}

	return user, nil
}

// updates the users login infromation
// checks if user is valid aswell
func Update_user(db *sql.DB, update_user *UpdateUser) (User, error) {
	var user User

	stmnt, err := db.Prepare("SELECT (username, password, id) FROM Users WHERE id=?;")
	if err != nil {
		return user, err
	}

	defer stmnt.Close()

	err = stmnt.QueryRow(&update_user.ID).Scan(&user.Username, &user.Password, &user.ID)
	if err != nil {
		return user, err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(update_user.OldPassword)); err != nil {
		return user, err
	}

	query := "UPDATE User SET "

	if user.Username != update_user.Username {
		query += ("username = " + update_user.Username)
	}

	if user.Password != update_user.NewPassword {
		new_hash, err := bcrypt.GenerateFromPassword([]byte(update_user.NewPassword), bcrypt.DefaultCost)
		if err != nil {
			return user, err
		}
		query += ("password = " + string(new_hash))
	}

	query += "WHERE id = ? RETURNING *;"

	stmnt, err = db.Prepare(query)
	if err != nil {
		return user, err
	}

	defer stmnt.Close()

	err = stmnt.QueryRow().Scan(&user.Username, &user.Password, &user.ID)
	return user, nil
}

// Gets user by username and password
func Login_user(login LoginUser, db *sql.DB) (User, error) {
	var user User

	stmnt, err := db.Prepare("SELECT username, password, id FROM Users WHERE username = ?;")
	if err != nil {
		return user, err
	}

	defer stmnt.Close()

	err = stmnt.QueryRow(&login.Username).Scan(&user.Username, &user.Password, &user.ID)
	if err != nil {
		return user, err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Password)); err != nil {
		return user, err
	}

	return user, nil
}

// Gets all the users from database
func Get_all_users(db *sql.DB) ([]User, error) {
	users := make([]User, 0)

	rows, err := db.Query("SELECT username, password, id FROM Users;")
	if err != nil {
		return users, err
	}

	defer rows.Close()
	for rows.Next() {
		var user User

		if err = rows.Scan(&user.Username, &user.Password, &user.ID); err != nil {
			return users, err
		}

		users = append(users, user)
	}

	return users, nil
}

// Gets all chats for a combination of the sender's id and the receiver's id.
func Get_chats(sender, receiver int, db *sql.DB) ([]Chat, error) {
	chats := make([]Chat, 0)

	rows, err := db.Query("SELECT id, sender, receiver, body FROM Chats WHERE receiver=? and sender=?;", &receiver, &sender)
	if err != nil {
		return chats, err
	}

	defer rows.Close()
	for rows.Next() {
		var chat Chat

		if err = rows.Scan(&chat.ID, &chat.Sender, &chat.Receiver, &chat.Body); err != nil {
			return chats, err
		}

		chats = append(chats, chat)
	}

	return chats, nil
}

// creates a new chat
func New_chat(new_chat NewChat, db *sql.DB) (Chat, error) {
	var chat Chat

	err := db.QueryRow("INSERT INTO chats(sender, receiver, body) VALUES (?,?,?) RETURNING *;", &new_chat.Sender, &new_chat.Receiver, &new_chat.Body).Scan(&chat.ID, &chat.Sender, &chat.Receiver, &chat.Body)
	if err != nil {
		return chat, err
	}

	return chat, nil
}
