package api

import (
	"net/http"
)

// FooBar handles recovery of requests.
func FooBar(w http.ResponseWriter, r *http.Request, m string) {

	w.Write([]byte(m))

}

func authenticate(w http.ResponseWriter, r *http.Request) {

	if r.FormValue("token") == "" {
		w.WriteHeader(http.StatusBadRequest)
		// For the love of god,
		// do not use errors.New("...")
		// to panic a request
		panic("Bad request")
	}

	if r.FormValue("token") != "token" {
		w.WriteHeader(http.StatusUnauthorized)
		panic("Unauthorized request")
	}

}
