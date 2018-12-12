package api

import (
	"context"
	"encoding/json"
	"time"

	"github.com/cheikhshift/samb-examples/crud-example/pkg/globals"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"github.com/mongodb/mongo-go-driver/mongo"

	"net/http"

	"gopkg.in/mgo.v2/bson"
)

var ShoesNotFound = "{\"error\" : \"Not supported\", \"code\" : 405 }"

// HandleShoes Routes PUT,PATCH,DELETE,POST and GET
// requests for corresponding
// handler Shoes.
func HandleShoes(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPut:
		PutShoes(w, r)
	case http.MethodPatch:
		PatchShoes(w, r)
	case http.MethodDelete:
		DeleteShoes(w, r)
	case http.MethodPost:
		PostShoes(w, r)
	case http.MethodGet:
		GetShoes(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		m := bson.M{"error": "Method not supported.", "code": 405}
		err := WriteAsJson(w, &m)
		if err != nil {
			panic(err)
		}
	}
}

// PutShoes handles Put HTTP requests
// for Shoes
func PutShoes(w http.ResponseWriter, r *http.Request) {

	var req bson.M

	objID, err := primitive.ObjectIDFromHex(r.URL.Path)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		panic(err.Error())
	}

	ParseBody(r, w, &req)

	delete(req, "_id")

	_ = getShoeCollection().FindOneAndUpdate(
		getContext(),
		bson.M{"_id": objID},
		bson.M{"$set": req},
	)

	var jsonResult = bson.M{"message": "resource updated!"}

	err = WriteAsJson(w, &jsonResult)

	if err != nil {
		panic(err)
	}

}

// PatchShoes handles Patch HTTP requests
// for Shoes
func PatchShoes(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Please implement me."))

}

// DeleteShoes handles Delete HTTP requests
// for Shoes
func DeleteShoes(w http.ResponseWriter, r *http.Request) {

	objID, err := primitive.ObjectIDFromHex(r.URL.Path)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		panic(err.Error())
	}

	res, err := getShoeCollection().DeleteOne(
		getContext(),
		bson.M{"_id": objID},
	)

	if err != nil {
		panic(err)
	}

	if res.DeletedCount == 0 {
		w.WriteHeader(http.StatusNotFound)
		panic("Specified item not found.")
	}

	err = WriteAsJson(w, res)

	if err != nil {
		panic(err)
	}
}

// PostShoes handles Post HTTP requests
// for Shoes
func PostShoes(w http.ResponseWriter, r *http.Request) {

	var req bson.M

	ParseBody(r, w, &req)

	mongoRes, err := getShoeCollection().InsertOne(getContext(), req)

	req["_id"] = mongoRes.InsertedID

	err = WriteAsJson(w, &req)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		panic(err)
	}

}

// GetShoes handles Get HTTP requests
// for Shoes
func GetShoes(w http.ResponseWriter, r *http.Request) {

	ctx := getContext()

	cur, err := getShoeCollection().Find(ctx, nil)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		panic(err.Error())
	}

	defer cur.Close(ctx)

	var results []bson.M

	for cur.Next(ctx) {
		var result bson.M
		err := cur.Decode(&result)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			panic(err.Error())
		}
		results = append(results, result)
	}

	if err := cur.Err(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		panic(err.Error())
	}

	err = WriteAsJson(w, &results)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		panic(err.Error())
	}

}

func WriteAsJson(w http.ResponseWriter, i interface{}) error {
	w.WriteHeader(http.StatusOK)
	enc := json.NewEncoder(w)
	err := enc.Encode(i)
	return err
}

func getContext() context.Context {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	return ctx
}

func getShoeCollection() *mongo.Collection {
	return globals.Client.Database("testing").Collection("shoes")
}

func ParseBody(r *http.Request, w http.ResponseWriter, i interface{}) {

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(i)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		// Important to panic using string.
		panic(err.Error())
	}

}
