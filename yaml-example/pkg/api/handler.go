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

	if strings.Contains(r.URL.Path, "/Foo") {

		if strings.Contains(r.URL.Path, "/Foo/Bar") && r.Method == "GET" {

			w.Write([]byte("Hello World"))
			return
		}
	}
	if strings.Contains(r.URL.Path, "/hello") && r.Method == "GET" {

		w.Write([]byte("Hello World"))
		return
	}
	if strings.Contains(r.URL.Path, "/hello_POST") && r.Method == "POST" {
		println("Request to Hello_post")
		w.Write([]byte("Hello World"))
		return
	}
}
