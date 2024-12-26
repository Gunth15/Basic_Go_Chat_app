package context

import (
	"log"
	"net/http"

	"github.com/chat_app/database"

	"github.com/a-h/templ"
	"github.com/chat_app/templates"
)

// Initializes a new router for User related routes
func NewUserMux(preface_url string, ctxt *Ctxt) *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("GET /signup", templ.Handler(templates.Signup(preface_url+"signup")))
	mux.HandleFunc("POST /signup", ctxt.PostSignup)
	mux.Handle("GET /login", templ.Handler(templates.Login(preface_url+"login")))
	mux.HandleFunc("POST /login", ctxt.PostLogin)
	mux.Handle("GET /update", templ.Handler(templates.Update(preface_url+"update")))
	mux.HandleFunc("POST /update", ctxt.PostUpdateLogin)
	return mux
}

func (ctxt *Ctxt) PostSignup(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Print(err)
		http.Error(w, "Invalid form", http.StatusBadRequest)
		return
	}
	form := r.PostForm
	if !form.Has("username") || !form.Has("password") || !form.Has("password_again") {
		log.Print(err)
		http.Error(w, "Invalid form data", http.StatusUnprocessableEntity)
		return
	}
	signup := database.NewUser{
		Username:      form.Get("username"),
		Password:      form.Get("password"),
		PasswordAgain: form.Get("password_again"),
	}

	new_user, err := ctxt.Users.New(ctxt.Db, signup)
	if err != nil {
		log.Print(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	log.Printf("Added user: %s", new_user.Username)
	http.Redirect(w, r, "/user/profile", http.StatusAccepted)
}

func (ctxt *Ctxt) PostLogin(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Print(err)
		http.Error(w, "Invalid form", http.StatusBadRequest)
		return
	}
	form := r.PostForm
	if !form.Has("username") || !form.Has("password") {
		log.Print(err)
		http.Error(w, "Invalid form data", http.StatusUnprocessableEntity)
	}
	login := database.LoginUser{
		Username: form.Get("username"),
		Password: form.Get("password"),
	}

	user, err := ctxt.Users.Login(ctxt.Db, login)
	if err != nil {
		log.Print(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	log.Printf("User %v logged in", user.Username)
	http.Redirect(w, r, "/user/profile", http.StatusAccepted)
}

func (ctxt *Ctxt) PostUpdateLogin(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Print(err)
		http.Error(w, "Invalid form", http.StatusBadRequest)
		return
	}
	form := r.PostForm
	if !form.Has("username") || !form.Has("password") || !form.Has("new_password") || !form.Has("new_password_again") {
		log.Print(err)
		http.Error(w, "Invalid form data", http.StatusUnprocessableEntity)
	}
	update := database.UpdateUser{
		Username:         form.Get("username"),
		OldPassword:      form.Get("password"),
		NewPassword:      form.Get("new_password"),
		NewPasswordAgain: form.Get("new_password_again"),
	}

	user, err := ctxt.Users.Update(ctxt.Db, &update)
	if err != nil {
		log.Print(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	log.Printf("User %s updated login info", user.Username)
	http.Redirect(w, r, "/user/profile", http.StatusAccepted)
}
