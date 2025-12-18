package main

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/Ayushmangit/api/internal/database"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) HandlerUpdateUser(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	if idStr == "" {
		RespondWithError(w, http.StatusBadRequest, "Id is required")
		return
	}
	id, err := uuid.Parse(idStr)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "invalid uuid")
		return
	}

	type parameters struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}
	params := parameters{}
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		RespondWithError(w, http.StatusBadRequest, "invalid json body")
		return
	}
	user, err := apiCfg.DB.UpdateUser(r.Context(), database.UpdateUserParams{
		ID:    id,
		Name:  params.Name,
		Email: params.Email,
	})
	if err != nil {
		if err == sql.ErrNoRows {
			RespondWithError(w, http.StatusNotFound, "user not found")
			return
		}
		RespondWithError(w, http.StatusInternalServerError, "could not update user")
		return
	}

	RespondWithJSON(w, http.StatusOK, databaseUserToUser(user))
}
