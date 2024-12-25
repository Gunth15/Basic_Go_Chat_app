package context

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/chat_app/templates"
)

func NewUserMux(preface_url string) *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("GET /signup", templ.Handler(templates.Signup(preface_url+"/signup")))
	mux.Handle("GET /login", templ.Handler(templates.Login(preface_url+"/login")))
	mux.Handle("GET /update", templ.Handler(templates.Update(preface_url+"/update")))
	return mux
}

func (ctxt *Ctxt) PostSignup(w http.ResponseWriter, r *http.Request) {
}

func (ctxt *Ctxt) PostLogin(w http.ResponseWriter, r *http.Request) {
}

func (ctxt *Ctxt) PutUpdateLogin(w http.ResponseWriter, r *http.Request) {
}
