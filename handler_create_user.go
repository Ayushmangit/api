package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Ayushmangit/api/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) HandlerCreateUser(w http.ResponseWriter, r *http.Request) {

	type parameters struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	params := parameters{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&params); err != nil {
		RespondWithError(w, 400, fmt.Sprintln("Error parsing json:", err))
		return
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		Name:      params.Name,
		Email:     params.Email,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	})

	if err != nil {
		RespondWithError(w, 400, fmt.Sprintln("internal server error couldn't create user:", err))
		return
	}

	RespondWithJSON(w, 200, databaseUserToUser(user))
}
