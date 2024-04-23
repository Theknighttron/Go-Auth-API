package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithError(w http.ResponseWriter, code int, msg string) {
	// check if the error code is server status code
	if code > 499 {
		log.Println("Responding with 500 error: ", msg)
	}

	type errResponse struct {
		Error string `json:"error"`
	}

	respondWithJson(w, code, errResponse{
		Error: msg,
	})

}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	// converting data to json formatt
	data, err := json.Marshal(payload)

	if err != nil {
		log.Printf("Failed to marshal JSON response %v", payload)
		w.WriteHeader(500)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}
