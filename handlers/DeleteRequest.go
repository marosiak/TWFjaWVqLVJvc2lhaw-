package handlers

import (
	"TWFjaWVqLVJvc2lhaw-/database"
	"encoding/json"
	. "github.com/go-chi/chi"
	"net/http"
	"strconv"
)

func DeleteRequest(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(URLParam(r, "requestId"))
	for i, request := range database.Requests {
		if request.ID == id {
			database.Requests = append(database.Requests[:i], database.Requests[i+1:]...)
			m := make(map[string]int)
			m["id"] = id
			json.NewEncoder(w).Encode(m)
			return
		}
	}
	http.Error(w, http.StatusText(404), 404)
}