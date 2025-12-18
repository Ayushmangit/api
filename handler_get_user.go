package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/Ayushmangit/api/internal/auth"
)

func (apiCfg *apiConfig) HandlerGetUser(w http.ResponseWriter, r *http.Request) {
	apiKey, err := auth.GetApiKey(r.Header)
	if err != nil {
		RespondWithError(w, 403, fmt.Sprintln("Auth error:", err))
		return
	}

	user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)
	if err != nil {
		if err == sql.ErrNoRows {
			RespondWithError(w, http.StatusBadRequest, "no user wiht id")
			return
		}
		RespondWithError(w, http.StatusBadRequest, "failed to retrieve user")
	}

	RespondWithJSON(w, http.StatusOK, databaseUserToUser(user))
}
