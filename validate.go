package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func handleValidation(w http.ResponseWriter, r *http.Request) {
	type params struct {
		Body string `json:"body"`
	}

	p := params{}
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		respondWithError(w, 500, fmt.Sprintf("Error decoding parameters: %s", err))
		return
	}
	if len(p.Body) > 140 {
		respondWithError(w, 400, "Chirp is too long")
		return
	}
	type returnVal struct {
		Valid bool `json:"valid"`
	}

	v := returnVal{
		Valid: true,
	}
	data, err := json.Marshal(v)
	if err != nil {
		respondWithError(w, 500, fmt.Sprintf("Error encoding: %s"))
		return
	}
	respondWithJSON(w, 200, data)

}
