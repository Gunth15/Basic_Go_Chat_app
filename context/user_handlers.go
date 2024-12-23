package context

import (
	"log"
	"net/http"
)

type FormData struct {
	Url      string
	IsNew    bool
	IsUpdate bool
}

func (ctxt *Ctxt) GetProfile(w http.ResponseWriter, r *http.Request) {
	ctxt.Tmpls["profile"].ExecuteTemplate(w, "profile.html", nil)
}

func (ctxt *Ctxt) GetSignup(w http.ResponseWriter, r *http.Request) {
	ctxt.Tmpls["signup"].ExecuteTemplate(w, "signup.html", nil)
}

func (ctxt *Ctxt) PostSignup(w http.ResponseWriter, r *http.Request) {
}

func (ctxt *Ctxt) GetLogin(w http.ResponseWriter, r *http.Request) {
	ctxt.Tmpls["login"].ExecuteTemplate(w, "login.html", nil)
}

func (ctxt *Ctxt) PostLogin(w http.ResponseWriter, r *http.Request) {
}

func (ctxt *Ctxt) GetUpdateLogin(w http.ResponseWriter, r *http.Request) {
	initform := FormData{
		IsNew:    false,
		IsUpdate: true,
		Url:      "/update",
	}
	err := ctxt.Tmpls["update"].ExecuteTemplate(w, "update.html", initform)
	if err != nil {
		log.Fatal(err)
	}
}

func (ctxt *Ctxt) PutUpdateLogin(w http.ResponseWriter, r *http.Request) {
}
