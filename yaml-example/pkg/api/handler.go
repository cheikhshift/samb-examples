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

	if basePath := "/Foogazi"; strings.Contains(r.URL.Path, basePath) && r.Method == "GET" {

		tools.ShortenPath(basePath, r)

		w.Write([]byte("Hello"))
		return
	}
	if basePath := "/Foo"; strings.Contains(r.URL.Path, basePath) {

		if basePath := "/Foo/Bar"; strings.Contains(r.URL.Path, basePath) && r.Method == "GET" {

			tools.ShortenPath(basePath, r)

			w.Write([]byte(r.URL.Path))
			return
		}

		tools.ShortenPath(basePath, r)

		w.Write([]byte("Hello world"))
		return
	}
	if basePath := "/hello"; strings.Contains(r.URL.Path, basePath) && r.Method == "GET" {

		tools.ShortenPath(basePath, r)

		w.Write([]byte("Hello World"))
		return
	}
	if basePath := "/hello_POST"; strings.Contains(r.URL.Path, basePath) && r.Method == "POST" {

		tools.ShortenPath(basePath, r)
		println("Request to Hello_post")
		w.Write([]byte("Hello World"))
		return
	}
}
