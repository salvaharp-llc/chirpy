package main

import (
	"database/sql"
	"net/http"

	"github.com/salvaharp-llc/chirpy/internal/auth"
)

func (cfg *apiConfig) handlerRevoke(w http.ResponseWriter, r *http.Request) {
	refreshToken, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Could not get token from header", err)
		return
	}

	_, err = cfg.db.RevokeRefreshToken(r.Context(), refreshToken)
	if err == sql.ErrNoRows {
		respondWithError(w, http.StatusBadRequest, "Could not find refresh token", err)
		return
	}
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Could not revoke the refresh token", err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
