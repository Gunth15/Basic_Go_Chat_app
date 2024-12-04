package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/chat_app/context"

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
	db, err := sql.Open("sqlite3", "./database/chat_app.db")
	if err != nil {
		log.Fatal(err)
	}

	tmpl, err := template.ParseFS(os.DirFS("./templates/"), "*.html", "*/*.html")
	if err != nil {
		log.Fatal(err)
	}

	ctxt := context.Ctxt{
		Db:   db,
		Tmpl: tmpl,
	}

	http.HandleFunc("GET /login/", ctxt.GetLogin)
	http.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		ctxt.Tmpl.ExecuteTemplate(w, "landing.html", nil)
	})
	http.HandleFunc("GET /signup/", ctxt.GetSignup)
	http.HandleFunc("GET /update/", ctxt.GetUpdateLogin)
	http.HandleFunc("GET /profile/", ctxt.GetProfile)

	port := ":8080"
	log.Printf("Running on localhost%s", port)
	defer log.Print("Shutting Down")
	log.Fatal(http.ListenAndServe(port, nil))
}
