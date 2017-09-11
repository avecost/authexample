package views

import (
	"net/http"
	"github.com/avecost/authexample/sessions"
	"github.com/avecost/authexample/config"
	"fmt"
)

func RequiresLogin(handler func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if !sessions.IsLoggedIn(r) {
			http.Redirect(w, r, "/login", 302)
			return
		}
		handler(w, r)
	}
}

func UserAllowedURL(slcUsers []config.AppUser, sUser, sUrl string) bool {
	for _, v := range slcUsers {
		if v.User == sUser {
			for _, url := range v.Url {
				if url == sUrl {
					return true
				}
			}
		}
	}
	return false
}