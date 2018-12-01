package api

import "net/http"

// GETSecureRes handles GET HTTP requests
// for secureRes
func GETSecureRes(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Please implement me."))

}
