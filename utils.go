package main

import (
	"encoding/json"
	"net/http"
)

func sendError(w http.ResponseWriter, statusCode int, err string) {
	response := map[string]string{"error": err}
	sendResponse(w, statusCode, response)
}

func sendResponse(w http.ResponseWriter, statusCode int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		Log.Errorf("Unable to marshal to JSON, review the payload")
		return
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(response)
}
