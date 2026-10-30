package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"slices"
	"strings"
	"time"

	"github.com/KamogeloMahlake/chirpy/internal/database"
	"github.com/google/uuid"
)

func (cfg *apiConfig) handleCreateChirp(w http.ResponseWriter, r *http.Request) {
	type params struct {
		Body string `json:"body"`
		UserId uuid.UUID `json:"user_id"`
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
		CleanedBody string `json:"cleaned_body"`
	}

	bannedWords := []string{"kerfuffle", "sharbert", "fornax"}
	bodySlice := strings.Fields(p.Body)

	for i, field := range bodySlice {
		if slices.Contains(bannedWords, field) {
			bodySlice[i] = "****"
		}
	}

	user, err := cfg.db.GetUser(r.Context(), p.UserId)
	if err != nil {
		respondWithError(w, 404, fmt.Sprint("User doesn't exist"))
		return
	}

	chirp, err := cfg.db.CreateChirp(r.Context(), database.CreateChirpParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Body: strings.Join(bodySlice, " "),
		UserID: user.ID,
	})
	if err != nil {
		respondWithError(w, 500, fmt.Sprintf("Error creating chirp: %s", err))
		return
	}

	data, err := json.Marshal(chirp)
	if err != nil {
		respondWithError(w, 500, fmt.Sprintf("Error encoding: %s\n", err))
		return
	}
	respondWithJSON(w, 200, data)

}

func (cfg *apiConfig) handleGetAllChirps(w http.ResponseWriter, r *http.Request) {
	chirps, err := cfg.db.GetAllChirps(r.Context())
	if err != nil {
		respondWithError(w, 500, fmt.Sprintf("Unable to get chirps: %s", err))
		return
	}
	
	data, err := json.Marshal(chirps)
	if err != nil {
		respondWithError(w, 500, fmt.Sprintf("Error encoding: %s\n", err))
		return
	}
	respondWithJSON(w, 200, data)
}

func (cfg *apiConfig) handleGetChirp( w http.ResponseWriter, r *http.Request) {
	chirp, err := cfg.db.GetChirp(r.Context(), uuid.MustParse(r.PathValue("chirpID")))
	if err != nil {
		respondWithError(w, 404, fmt.Sprintf("Error Unable to found chirp: %s", err))
		return
	}

	data, err := json.Marshal(chirp)
	if err != nil {
		respondWithError(w, 500, fmt.Sprintf("Error encoding: %s\n", err))
		return
	}
	respondWithJSON(w, 200, data)
}
