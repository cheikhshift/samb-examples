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

	if basePath := "/shoes/"; strings.Contains(r.URL.Path, basePath) {

		tools.ShortenPath(basePath, r)
		w.Header().Set("Content-Type", "application/json")
		println(r.Method + " request to " + r.URL.Path)
		HandleShoes(w, r)
		return
	}
}
