package handlers

import (
	"TWFjaWVqLVJvc2lhaw-/database"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func CreateRequest(w http.ResponseWriter, r *http.Request) {
	var newRequest database.Request
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Something went wrong")
	}
	_ = json.Unmarshal(reqBody, &newRequest)
	newRequest.ID = len(database.Events)+1
	database.Events = append(database.Events, newRequest)

	w.WriteHeader(http.StatusOK) // In my opinion it should return StatusCreated (201), but the specs says to return 200

	_ = json.NewEncoder(w).Encode(newRequest)
}