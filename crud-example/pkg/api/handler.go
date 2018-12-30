//package api contains your web app's handler definitions.
// GENERATED CODE, DO NOT EDIT!!
package api

import (
	"net/http"
	"strings"

	"github.com/cheikhshift/samb/tools"
)

// Handles routing of application.
func Handler(w http.ResponseWriter, r *http.Request) {

	defer catchPanic(w, r)

	if basePath := "/shoes/"; r.URL.Path == basePath {
		w.Header().Set("Content-Type", "application/json")
		println(r.Method + " request to " + r.URL.Path)

		if basePath := "/shoes//shoes_sub/"; strings.Contains(r.URL.Path, basePath) {
			w.Header().Set("Content-Type", "application/json")
			println(r.Method + " request to " + r.URL.Path)

			tools.ShortenPath(basePath, r)
			HandleShoes(w, r)
			return
		}

		tools.ShortenPath(basePath, r)
		HandleShoes(w, r)
		return
	}
}
