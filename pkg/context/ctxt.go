package context

import (
	"database/sql"

	"github.com/chat_app/pkg/database"
)

// Ctxt contains interfaces for database querying.
// Also includes handlers for accesing databse
type Ctxt struct {
	Db     *sql.DB
	Secret []byte
	Users  interface {
		Get(*sql.DB, int) (database.User, error)
		New(*sql.DB, database.NewUser) (database.User, error)
		All(*sql.DB) ([]database.User, error)
		Update(*sql.DB, *database.UpdateUser) (database.User, error)
		Login(*sql.DB, database.LoginUser) (database.User, error)
	}
	Chats interface {
		New(*sql.DB, database.NewChat) (database.Chat, error)
		UserAll(*sql.DB, int, int) ([]database.Chat, error)
	}
}

// QueryUsers satisfies the User interface and abstract queries form tha databse.
// QueryChats satisfies the Chats interface and abstract queries form tha databse.
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

func (qu *QueryUsers) Login(db *sql.DB, login database.LoginUser) (database.User, error) {
	return database.Login_user(login, db)
}

func (qu *QueryUsers) All(db *sql.DB) ([]database.User, error) {
	return database.Get_all_users(db)
}

func (qu *QueryUsers) Update(db *sql.DB, update *database.UpdateUser) (database.User, error) {
	return database.Update_user(db, update)
}
