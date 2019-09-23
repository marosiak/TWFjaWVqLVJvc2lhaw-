package handlers

import (
	"TWFjaWVqLVJvc2lhaw-/database"
	"encoding/json"
	"net/http"
)

func RequestsList(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(database.Requests)
}