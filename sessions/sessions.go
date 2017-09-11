package sessions

import (
	"github.com/gorilla/sessions"
	"net/http"
	"fmt"
)

var Store = sessions.NewCookieStore([]byte("AVECOST"))

func IsLoggedIn(r *http.Request) bool {
	session, err := Store.Get(r, "session")
	if err != nil {
		fmt.Println("Session error: ", err)
	}
	if session.Values["loggedIn"] == true {
		return true
	}
	return false
}

func Username(r *http.Request) string {
	session, err := Store.Get(r, "session")
	if err != nil {
		fmt.Println("Session error: ", err)
	}
	return session.Values["username"].(string)
}
