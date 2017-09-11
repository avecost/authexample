package main

import (
	"net/http"
	"fmt"
	"html/template"
	"github.com/avecost/authexample/sessions"
	"github.com/avecost/authexample/views"
	"github.com/avecost/authexample/config"
	"github.com/avecost/authexample/password"
)

var Users []config.AppUser

func loginHandler(w http.ResponseWriter, r *http.Request) {
	session, err := sessions.Store.Get(r, "session")
	tpl, _ := template.ParseFiles("template/login.gtpl")
	if err != nil {
		tpl.Execute(w, nil)
	} else {
		isLoggedIn := session.Values["loggedIn"]
		if isLoggedIn != true {
			if r.Method == "POST" {
				if password.ValidUser(r.FormValue("username"), r.FormValue("password"), Users) {
					session.Values["loggedIn"] = true
					session.Values["username"] = r.FormValue("username")
					session.Save(r, w)
					http.Redirect(w, r, "/", 302)
					return
				} else {
					http.Redirect(w, r, "/login", 302)
				}
			} else if r.Method == "GET" {
				tpl.Execute(w, nil)
			}
		} else {
			http.Redirect(w, r, "/", 302)
		}
	}
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	session, err := sessions.Store.Get(r, "session")
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	session.Options.MaxAge = -1
	session.Values["loggedIn"] = false
	session.Values["username"] = ""
	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	http.Redirect(w, r, "/", 302)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if sessions.IsLoggedIn(r) {
		tpl, err := template.ParseFiles("template/home.gtpl")
		if err != nil {
			fmt.Println("Error parsing template")
		}
		tpl.Execute(w, nil)
	} else {
		http.Redirect(w, r, "/login", 302)
	}
}

func selectedHandler(w http.ResponseWriter, r *http.Request) {
	if views.UserAllowedURL(Users, sessions.Username(r), r.URL.Path) {
		tpl, err := template.ParseFiles("template/selected.gtpl")
		if err != nil {
			fmt.Println("Error parsing template")
		}
		tpl.Execute(w, nil)
	} else {
		http.Redirect(w, r, "/", 302)
	}
}

func main() {
	Users = config.LoadConfiguration("./users.json")

	mux := http.NewServeMux()
	mux.HandleFunc("/login", loginHandler)
	mux.HandleFunc("/logout", logoutHandler)
	mux.HandleFunc("/selected", views.RequiresLogin(selectedHandler))
	mux.HandleFunc("/", views.RequiresLogin(homeHandler))

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println(err)
	}
}
