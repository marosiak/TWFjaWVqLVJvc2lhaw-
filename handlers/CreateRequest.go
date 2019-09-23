package handlers

import (
	"TWFjaWVqLVJvc2lhaw-/database"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)



func test(interval time.Duration) {
	database.GlobalShit = database.GlobalShit +1
	time.Sleep(interval * time.Second)
	go test(interval)
}

func CreateRequest(w http.ResponseWriter, r *http.Request) {
	var newRequest database.Request
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Something went wrong")
	}
	_ = json.Unmarshal(reqBody, &newRequest)
	newRequest.ID = len(database.Requests)+1
	var newHistory database.History
	newHistory.Response = "dupa"
	newHistory.CreatedAt = 22222
	newHistory.Duration = 0.3
	newRequest.History = append(newRequest.History, newHistory)
	database.Requests = append(database.Requests, newRequest)

	w.WriteHeader(http.StatusOK) // In my opinion it should return StatusCreated (201), but the spec says to return 200

	_ = json.NewEncoder(w).Encode(newRequest)


	//go test(2)


	// Well, the spec says to return {"id": x}, but I'm returning url and interval as well,
	// as it wont take much more resources, but it could be helpfully for frontend devs
}