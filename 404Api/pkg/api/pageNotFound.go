package api

import (
	"net/http"
)

func pageNotFound(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`{"error":"Res. not found.", "code": 404 }`))
}
