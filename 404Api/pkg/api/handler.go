//package api contains your web app's handler definitions.
// GENERATED CODE, DO NOT EDIT!!
package api

import (
	"fmt"
	"net/http"
	"strings"
)

// Handles routing of application.
func Handler(w http.ResponseWriter, r *http.Request) {

	defer catchPanic(w, r)

	if strings.Contains(r.URL.Path, "/") {
		println("Hello")

		if strings.Contains(r.URL.Path, "/Foo") && r.Method == "GET" {
			println("Hello")
			fmt.Println("Hello")
			return
		}
		if strings.Contains(r.URL.Path, "/") {

			pageNotFound(w)
			return
		}
	}
}
