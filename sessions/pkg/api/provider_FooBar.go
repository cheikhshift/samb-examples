package api

import (
	"net/http"

	"github.com/gorilla/sessions"
)

// Ripped for Gorrilla sample on Github
// Note: Don't store your key in your source code.
var store = sessions.NewCookieStore([]byte("customKey"))

// FooBar does...
func providerFooBar(w http.ResponseWriter, r *http.Request) *string {

	session, _ := store.Get(r, "session-name")
	// Set some session values.
	if _, ok := session.Values["UserID"]; !ok {
		session.Values["UserID"] = "Hello From a provider function."
		
	}

	// Save it before we write to the response/return from the handler.
	session.Save(r, w)

	result := session.Values["UserID"].(string) 

	w.Write(
		[]byte(
			result,
		),
	)

	return &result
}
