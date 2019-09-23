package handlers

import (
	"TWFjaWVqLVJvc2lhaw-/database"
	"encoding/json"
	. "github.com/go-chi/chi"
	"net/http"
	"strconv"
)

func RequestDetail(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(URLParam(r, "requestId"))

	request, error := database.GetById(id)
	if error != nil {
		http.Error(w, http.StatusText(404), 404)
		return
	}

	json.NewEncoder(w).Encode(request)
}