package main

import (
	"TWFjaWVqLVJvc2lhaw-/handlers"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreate(t *testing.T) {
	var body = []byte(`{"url": "https://httpbin.org/range/15", "interval": 2}`)
	req, err := http.NewRequest("POST", "/", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.CreateRequest)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	respMap := make(map[string]interface{})
	err = json.Unmarshal([]byte(rr.Body.String()), &respMap)
	if err != nil {
		t.Errorf("Error while parsing json")
	}

	expected := "https://httpbin.org/range/15"
	if respMap["url"] != expected {
		t.Errorf("handler returned unexpected url: got %v want %v",
			respMap["url"], expected)
	}
}