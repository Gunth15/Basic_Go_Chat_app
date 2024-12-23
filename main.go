package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"

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

	logint, err := template.ParseFiles("./templates/base/base.html", "./templates/user/form.html", "./templates/user/login.html")
	if err != nil {
		log.Fatal(err)
	}

	updatet, err := template.ParseFiles("./templates/base/base.html", "./templates/user/form.html", "./templates/user/update.html")
	if err != nil {
		log.Fatal(err)
	}

	landingt, err := template.ParseFiles("./templates/base/base.html", "./templates/base/landing.html")
	if err != nil {
		log.Fatal(err)
	}

	profilet, err := template.ParseFiles("./templates/base/base.html", "./templates/user/profile.html")
	if err != nil {
		log.Fatal(err)
	}

	signupt, err := template.ParseFiles("./templates/base/base.html", "./templates/user/form.html", "./templates/user/signup.html")
	if err != nil {
		log.Fatal(err)
	}

	ctxt := context.Ctxt{
		Db: db,
		Tmpls: map[string]*template.Template{
			"update":  updatet,
			"landing": landingt,
			"login":   logint,
			"profile": profilet,
			"signup":  signupt,
		},
	}

	http.HandleFunc("GET /login/", ctxt.GetLogin)
	http.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		err := ctxt.Tmpls["landing"].ExecuteTemplate(w, "landing.html", nil)
		if err != nil {
			log.Print(err)
		}
	})
	http.HandleFunc("GET /signup/", ctxt.GetSignup)
	http.HandleFunc("GET /update/", ctxt.GetUpdateLogin)
	http.HandleFunc("GET /profile/", ctxt.GetProfile)

	port := ":8080"
	log.Printf("Running on localhost%s", port)
	defer log.Print("Shutting Down")
	log.Fatal(http.ListenAndServe(port, nil))
}
