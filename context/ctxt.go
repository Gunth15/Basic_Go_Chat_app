package context

import (
	"database/sql"
	"html/template"

	"github.com/chat_app/database"

	_ "github.com/mattn/go-sqlite3"
)

// The environment for the application which includes the tempkating engine and the database connection
// Also controls queries for both models
// All methods are handlers for a http route
type Ctxt struct {
	Db    *sql.DB
	Tmpls map[string]*template.Template
	Users interface {
		Get(*sql.DB, int) (database.User, error)
		New(*sql.DB, database.NewUser) (database.User, error)
		All(*sql.DB) ([]database.User, error)
		Update(*sql.DB, *database.UpdateUser) (database.User, error)
	}
	Chats interface {
		New(*sql.DB, database.NewChat) (database.Chat, error)
		UserAll(*sql.DB, int, int) ([]database.Chat, error)
	}
}

// Implements the query interfaces for CTXT
type (
	QueryUsers struct{}
	QueryChats struct{}
)

func (qc *QueryChats) New(db *sql.DB, new_chat database.NewChat) (database.Chat, error) {
	return database.New_chat(new_chat, db)
}

func (qc *QueryChats) UserAll(db *sql.DB, sender, receiver int) ([]database.Chat, error) {
	return database.Get_chats(sender, receiver, db)
}

func (qu *QueryUsers) Get(db *sql.DB, id int) (database.User, error) {
	return database.Get_user(db, id)
}

func (qu *QueryUsers) New(db *sql.DB, new database.NewUser) (database.User, error) {
	return database.New_user(db, new)
}

func (qu *QueryUsers) All(db *sql.DB) ([]database.User, error) {
	return database.Get_all_users(db)
}

func (qu *QueryUsers) Update(db *sql.DB, update *database.UpdateUser) (database.User, error) {
	return database.Update_user(db, update)
}
