package handlers

import (
	"TWFjaWVqLVJvc2lhaw-/database"
	"encoding/json"
	"io/ioutil"
	"net/http"
)


func CreateRequest(w http.ResponseWriter, r *http.Request) {

	var newRequest database.Request
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		if len(reqBody) >= 1000*1000 {
			http.Error(w, "Ponad 1 MB danych...'", http.StatusRequestEntityTooLarge) // I'd propably write the error in english
		}
	}
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(reqBody, &newRequest)
	newRequest.ID = len(database.Requests)
	database.Requests = append(database.Requests, newRequest)

	FetchData(newRequest.ID)
	w.WriteHeader(http.StatusOK) // In my opinion it should return StatusCreated (201), but the spec says to return 200

	_ = json.NewEncoder(w).Encode(newRequest)

	// Well, the spec says to return {"id": x}, but I'm returning url and interval as well,
	// as it wont take much more resources, but it could be helpfully for frontend devs
}