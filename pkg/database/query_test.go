package database

import (
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB = start_db()

func start_db() *sql.DB {
	db, err := sql.Open("sqlite3", "./chat_app.db")
	if err != nil {
		panic(err)
	}
	return db
}

func TestUsers(t *testing.T) {
	t.Run("Add User", func(t *testing.T) { test_add_user(t) })
	t.Run("Get user", func(t *testing.T) { test_get_user(t) })
	t.Run("Get all users", func(t *testing.T) { test_get_all_users(t) })
	t.Run("Update users", func(t *testing.T) { test_get_all_users(t) })
	t.Run("Login user", func(t *testing.T) { test_login(t) })
}

func TestChats(t *testing.T) {
	t.Run("Add chat", func(t *testing.T) { test_add_chat(t) })
	t.Run("Get chats", func(t *testing.T) { test_get_chats(t) })
}

func test_add_user(t *testing.T) {
	newuser := NewUser{
		Username:      "Help",
		Password:      "Password",
		PasswordAgain: "Password",
	}

	user, err := New_user(db, newuser)
	if err != nil {
		t.Fatal(err)
		return
	}

	if user.Username != newuser.Username {
		t.Fatal("Usernames do not match")
		return
	}
	if user.Password != newuser.Password {
		t.Fatal("Passwords do not match")
		return
	}
}

func test_get_user(t *testing.T) {
	user, err := Get_user(db, 1)
	if err != nil {
		t.Fatal(err)
		return
	}

	if user.ID != 1 {
		t.Fatalf("ID does not match ID: %d", user.ID)
		return
	}
}

func test_get_all_users(t *testing.T) {
	_, err := Get_all_users(db)
	if err != nil {
		t.Fatal(err)
		return
	}
}

func test_update_user(t *testing.T) {
	update_user := UpdateUser{
		Username:    "Helps",
		NewPassword: "blah",
		OldPassword: "Password",
		ID:          1,
	}

	user, err := Update_user(db, &update_user)
	if err != nil {
		t.Fatal(err)
		return
	}

	if user.Username != update_user.Username {
		t.Fatal("Username did not update")
		return
	}
	if user.Password != update_user.NewPassword {
		t.Fatal("Password did not update")
		return
	}
}

func test_login(t *testing.T) {
	login := LoginUser{
		Username: "Helps",
		Password: "Blah",
	}

	user, err := Login_user(login, db)
	if err != nil {
		t.Fatal(err)
		return
	}

	if user.Username != login.Username {
		t.Fatal("Username does not match login")
		return
	}

	if user.Password != login.Password {
		t.Fatal("Password does not match login")
		return
	}
}

func test_add_chat(t *testing.T) {
	new_chat := NewChat{
		Sender:   1,
		Receiver: 1,
		Body:     "This is a test, do not send messages to yourself",
	}

	chat, err := New_chat(new_chat, db)
	if err != nil {
		t.Fatal(err)
		return
	}

	if new_chat.Sender != chat.Sender {
		t.Fatal("Sender aint right")
		return
	}

	if new_chat.Receiver != chat.Receiver {
		t.Fatal("Receiver aint right")
		return
	}

	if new_chat.Body != chat.Body {
		t.Fatal("Receiver aint right")
		return
	}
}

func test_get_chats(t *testing.T) {
	_, err := Get_chats(1, 1, db)
	if err != nil {
		t.Fatal(err)
		return
	}
}
