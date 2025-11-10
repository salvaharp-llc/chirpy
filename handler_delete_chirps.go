package main

import (
	"database/sql"
	"net/http"

	"github.com/google/uuid"
	"github.com/salvaharp-llc/chirpy/internal/auth"
)

func (cfg *apiConfig) handlerDeleteChirps(w http.ResponseWriter, r *http.Request) {
	chirpID, err := uuid.Parse(r.PathValue("chirpID"))
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid chirp ID", err)
		return
	}

	accessToken, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Could not get token from header", err)
		return
	}

	userID, err := auth.ValidateJWT(accessToken, cfg.JWTSecret)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Invalid token", err)
		return
	}

	dbChirp, err := cfg.db.GetChirp(r.Context(), chirpID)
	if err == sql.ErrNoRows {
		respondWithError(w, http.StatusNotFound, "Could not find chirp", err)
		return
	}
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Could not retrieve the chirp", err)
		return
	}

	if dbChirp.UserID != userID {
		respondWithError(w, http.StatusForbidden, "Not the author of the chirp", nil)
		return
	}

	err = cfg.db.DeleteChirp(r.Context(), dbChirp.ID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Could not delete the chirp", err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
