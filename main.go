package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/a-h/templ"
	"github.com/chat_app/context"
	"github.com/chat_app/middleware"
	"github.com/chat_app/templates"
	"github.com/joho/godotenv"

	_ "github.com/mattn/go-sqlite3"
)

/*
Need these routes:
- landing page/ GET
- logout button / Remove cookie
- profile page/ GET(should be able to see current chat rooms)
- page so users can find each other(Look to create new rooms)
- rooms(places where users commmunicate to each other using web sockets; when users online, open ws for sending. If the other user is online, open ws for recieving.)

Involve cookies:
How does one use cookies?
*/

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open("sqlite3", os.Getenv("DATABASE"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	ctxt := context.Ctxt{
		Db:     db,
		Secret: []byte(os.Getenv("SECRET_KEY")),
		Users:  &context.QueryUsers{},
		Chats:  &context.QueryChats{},
	}

	main_mux := http.NewServeMux()
	main_mux.Handle("/user/", http.StripPrefix("/user", context.NewUserMux("/user/", &ctxt)))
	main_mux.Handle("GET /{$}", templ.Handler(templates.Landing()))

	server := &http.Server{
		Addr:    os.Getenv("PORT"),
		Handler: middleware.Logger(main_mux),
	}

	log.Printf("Running on localhost%s", server.Addr)
	defer log.Print("Shutting Down.......")
	log.Fatal(server.ListenAndServe())
}
