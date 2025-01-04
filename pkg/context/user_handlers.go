// Package context inludes handlers for chat_app and Ctxt struct.
//
// Ctxt stuct holds database connection and other resources  handlers may use.
package context

import (
	"log"
	"net/http"

	"github.com/chat_app/pkg/database"
	"github.com/chat_app/pkg/middleware"

	"github.com/a-h/templ"
	"github.com/chat_app/pkg/cookies"
	"github.com/chat_app/web/templates"
)

// NewUserMux initializes a new router for User related routes.
func NewUserMux(preface_url string, secret []byte, ctxt *Ctxt) *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("GET /signup/", templ.Handler(templates.Signup(preface_url+"signup/")))
	mux.HandleFunc("POST /signup/", ctxt.post_signup)
	mux.Handle("GET /login/", templ.Handler(templates.Login(preface_url+"login/")))
	mux.HandleFunc("POST /login/", ctxt.post_login)
	mux.Handle("GET /update/", templ.Handler(templates.Update(preface_url+"update/")))
	mux.HandleFunc("POST /update/", ctxt.post_update_login)
	mux.Handle("GET /profile/", middleware.IsAuth(http.HandlerFunc(ctxt.profile), secret))
	return mux
}

// PostSignup adds a new user to the server.
func (ctxt *Ctxt) post_signup(w http.ResponseWriter, r *http.Request) {
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
	http.Redirect(w, r, "/user/profile/", http.StatusSeeOther)
}

// PostLogin verifies that a user exist in the databse.
func (ctxt *Ctxt) post_login(w http.ResponseWriter, r *http.Request) {
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
		return
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

	if err = cookies.Set(w, r, user, ctxt.Secret); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/user/profile/", http.StatusSeeOther)
}

// PostUpdateLogin updates a user's login information.
func (ctxt *Ctxt) post_update_login(w http.ResponseWriter, r *http.Request) {
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

	if err = cookies.Set(w, r, user, ctxt.Secret); err != nil {
		log.Print(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	http.Redirect(w, r, "/user/profile/", http.StatusSeeOther)
}

// Profile grabs authorization information fomr the context and fills the template with said info
func (ctxt *Ctxt) profile(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value(middleware.UserKey).(database.User)
	if !ok {
		log.Print("Error: No login information provided")
		http.Error(w, "Unauthorized Request", http.StatusBadRequest)
		return
	}

	db_user, err := ctxt.Users.Get(ctxt.Db, user.ID)
	if err != nil {
		log.Printf("Error: Bad login information for user %s", user.Username)
		http.Error(w, "Unauthorized Request", http.StatusBadRequest)
		return
	}

	if user.Username != db_user.Username {
		log.Printf("Error: Bad login information for user %s", user.Username)
		http.Error(w, "Unauthorized Request", http.StatusBadRequest)
		return
	}

	err = templates.Profile(user).Render(r.Context(), w)
	if err != nil {
		log.Println(err)
		http.Error(w, "Page unavailable", http.StatusUnauthorized)
		return
	}
}
