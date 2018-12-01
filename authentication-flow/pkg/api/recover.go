package api

import "net/http"

// Function used to get
// error message related
// to panic.
func catchPanic(w http.ResponseWriter, r *http.Request) {

	if n := recover(); n != nil {

		FooBar(w, r, n.(string))

	}
}
