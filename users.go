package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/KamogeloMahlake/chirpy/internal/database"
	"github.com/google/uuid"
)

func (cfg *apiConfig) handleCreateUser(w http.ResponseWriter, r *http.Request) {
	type params struct {
		Email string `json:"email"`
	}
  p := params{}
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		respondWithError(w, 500, fmt.Sprintf("Error decoding parameters: %s", err))
		return
	}
	user, err := cfg.db.CreateUser(r.Context(), database.CreateUserParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Email: p.Email,
	})
	if err != nil {
		respondWithError(w, 500, fmt.Sprintf("Error saving user: %s", err))
		return
	}

	data, err := json.Marshal(user)
	if err != nil {
		respondWithError(w, 500, fmt.Sprintf("Error encoding: %s", err))
		return
	}
	respondWithJSON(w, 201, data)

}
