package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) HandlerDestroyUser(w http.ResponseWriter, r *http.Request) {
	// if using query params
	//idStr := r.URL.Query().Get("id")
	idStr := chi.URLParam(r, "id")

	if idStr == "" {
		RespondWithError(w, http.StatusBadRequest, fmt.Sprintln("id is required"))
		return
	}

	id, err := uuid.Parse(idStr)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, fmt.Sprintln("invalid uuid format"))
		return
	}

	id, err = apiCfg.DB.DestroyUser(r.Context(), id)
	if err == sql.ErrNoRows {
		RespondWithError(w, http.StatusInternalServerError, fmt.Sprintln("No user with the given Id", err))
		return
	}

	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, fmt.Sprintln("internal server error couldn't delete user:", err))
		return
	}

	RespondWithJSON(w, 200, struct{}{})
}
