package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithError(w http.ResponseWriter, code int, msg string, logErr error) {
	if logErr != nil {
		log.Println(logErr)
	}
	if code > 499 {
		log.Printf("Responding with 5XX error: %s", msg)
	}
	type errorResponse struct {
		Error string `json:"error"`
	}
	respondWithJSON(w, code, errorResponse{
		Error: msg,
	})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
    w.Header().Set("Content-Type", "application/json")
    // Set status code BEFORE writing the body
    w.WriteHeader(code)
    
    dat, err := json.Marshal(payload)
    if err != nil {
        log.Printf("Error marshalling JSON: %s", err)
        // Too late to change status code here
        return
    }
    
    _, err = w.Write(dat)
    if err != nil {
        log.Printf("Error writing payload: %s", err)
        // Too late to change status code here
        return
    }
}
