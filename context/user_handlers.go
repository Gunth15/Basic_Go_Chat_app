package context

import (
	"net/http"
)

func (ctxt *Ctxt) GetProfile(w http.ResponseWriter, r *http.Request) {
	ctxt.Tmpl.ExecuteTemplate(w, "profile.html", nil)
}

func (ctxt *Ctxt) GetSignup(w http.ResponseWriter, r *http.Request) {
	ctxt.Tmpl.ExecuteTemplate(w, "signup.html", nil)
}

func (ctxt *Ctxt) PostSignup(w http.ResponseWriter, r *http.Request) {
}

func (ctxt *Ctxt) GetLogin(w http.ResponseWriter, r *http.Request) {
	ctxt.Tmpl.ExecuteTemplate(w, "login.html", nil)
}

func (ctxt *Ctxt) PostLogin(w http.ResponseWriter, r *http.Request) {
}

func (ctxt *Ctxt) GetUpdateLogin(w http.ResponseWriter, r *http.Request) {
	ctxt.Tmpl.ExecuteTemplate(w, "update.html", nil)
}

func (ctxt *Ctxt) PutUpdateLogin(w http.ResponseWriter, r *http.Request) {
}
