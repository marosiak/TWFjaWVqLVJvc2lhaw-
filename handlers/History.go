package handlers

import (
	"TWFjaWVqLVJvc2lhaw-/database"
	"encoding/json"
	. "github.com/go-chi/chi"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

var tr = &http.Transport{
	IdleConnTimeout:    60 * time.Second,
} // It may be created somewhere globally if I'd use it in more places than just this file

func FetchData(id int) {
	// Well, I've been trying to just pass Request object in argument
	// instead of iterating eveytime, but I need index to swap the data (history array)
	// I think using ORM instead of dummy database would be actually easier and better optimized, but it's a bit too late
	start := time.Now()

	for i, request := range database.Requests {
		if request.ID == id {
			var newHistory database.History

			newHistory.CreatedAt = time.Now().Unix()
			//newHistory.Response = null
			client := &http.Client{Transport: tr}
			resp, err := client.Get(request.Url)
			if err == nil {
				bodyBytes, err2 := ioutil.ReadAll(resp.Body)
				if err2 == nil {
					newHistory.Response = string(bodyBytes)
				}
			}

			//newHistory.Duration = float64(time.Now().Unix() - newHistory.CreatedAt)
			// ^ That was my first attempt which actually worked pretty good, but it wasn't precise as unix time doesnt give miliseconds

			newHistory.Duration = float32(time.Since(start).Seconds())
			request.History = append(request.History, newHistory)
			database.Requests[i] = request
			time.Sleep(request.Interval * time.Second)
			go FetchData(request.ID)
		}
	}
}

func RequestHistory(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(URLParam(r, "requestId"))
	for _, request := range database.Requests {
		if request.ID == id {
			json.NewEncoder(w).Encode(request.History)
			return
		}
	}
	http.Error(w, http.StatusText(404), 404)
}
