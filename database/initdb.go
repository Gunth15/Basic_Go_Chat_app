package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./chat_app.db")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	create_users := `
  CREATE TABLE Users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    username TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL
  );
  `

	create_chats := `
  CREATE TABLE Chats (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    sender INTEGER NOT NULL,
    receiver INTEGER NOT NULL,
    body TEXT NOT NULL,
    FOREIGN KEY (sender) REFERENCES Users(id),
    FOREIGN KEY (receiver) REFERENCES Users(id)
  );
  `

	_, err = db.Exec(create_users)
	if err != nil {
		log.Printf("%q: %s\n", err, create_users)
	}

	_, err = db.Exec(create_chats)
	if err != nil {
		log.Printf("%q: %s\n", err, create_chats)
	}
}
