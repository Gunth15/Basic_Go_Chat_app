// Package main runs chat_app server.
package main

import (
	"database/sql"
	"encoding/gob"
	"log"
	"net/http"
	"os"

	"github.com/a-h/templ"
	"github.com/chat_app/pkg/context"
	"github.com/chat_app/pkg/database"
	"github.com/chat_app/pkg/middleware"
	"github.com/chat_app/web/templates"
	"github.com/joho/godotenv"

	_ "github.com/mattn/go-sqlite3"
)

/*
Need these routes:
- logout button / Remove cookie
- profile page/ GET(should be able to see current chat rooms)
- page so users can find each other(Look to create new rooms)
- rooms(places where users commmunicate to each other using web sockets; when users online, open ws for sending. If the other user is online, open ws for recieving.)
*/

func main() {
	gob.Register(&database.User{})
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open("sqlite3", os.Getenv("DATABASE"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	secret := []byte(os.Getenv("SECRET_KEY"))

	ctxt := context.Ctxt{
		Db:     db,
		Secret: secret,
		Users:  &context.QueryUsers{},
		Chats:  &context.QueryChats{},
	}

	main_mux := http.NewServeMux()
	main_mux.Handle("/user/", http.StripPrefix("/user", context.NewUserMux("/user/", secret, &ctxt)))
	main_mux.Handle("GET /{$}", templ.Handler(templates.Landing()))
	main_mux.Handle("/", http.FileServer(http.Dir("./web/static/")))

	server := &http.Server{
		Addr:    os.Getenv("PORT"),
		Handler: middleware.Logger(main_mux),
	}
	log.Printf("Running on localhost%s", server.Addr)
	defer log.Print("Shutting Down.......")
	log.Fatal(server.ListenAndServe())
}
