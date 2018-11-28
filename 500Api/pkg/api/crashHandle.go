package api

import (
	"fmt"
	"net/http"
)

func crashHandle(w http.ResponseWriter, r *http.Request, m string) {

	w.WriteHeader(http.StatusInternalServerError)

	json := fmt.Sprintf(`{ "error" : "%s" , code : 500 }`, m)

	w.Write([]byte(json))

}
