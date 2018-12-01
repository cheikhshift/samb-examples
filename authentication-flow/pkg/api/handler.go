//package api contains your web app's handler definitions.
// GENERATED CODE, DO NOT EDIT!!
package api

import (
	"net/http"
	"strings"
)

// Handles routing of application.
func Handler(w http.ResponseWriter, r *http.Request) {

	defer catchPanic(w, r)

	if strings.Contains(r.URL.Path, "/") {
		authenticate(w,
			r)

		if strings.Contains(r.URL.Path, "/Foo") && r.Method == "GET" {

			GETSecureRes(w, r)
			return
		}
	}
}
