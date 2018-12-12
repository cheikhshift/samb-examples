package api

import (
	"encoding/json"
	"net/http"

	"gopkg.in/mgo.v2/bson"
)

// jsonPanic handles recovery of requests.
func jsonPanic(w http.ResponseWriter, r *http.Request, m string) {

	res := bson.M{"error": m}

	enc := json.NewEncoder(w)
	err := enc.Encode(&res)

	if err != nil {
		panic(err)
	}

}
