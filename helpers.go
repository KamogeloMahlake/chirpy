package main

import (
	"fmt"
	"log"
	"net/http"
)

func respondWithError(w http.ResponseWriter, code int, msg string) {
	log.Println(msg)
	w.WriteHeader(code)
}

func respondWithJSON(w http.ResponseWriter, code int, payload []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(payload)
	w.Write([]byte(fmt.Sprint("\n")))
}
