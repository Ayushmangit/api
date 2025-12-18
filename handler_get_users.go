package main

import (
	"fmt"
	"net/http"
)

func (apiCfg *apiConfig) HandlerGetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := apiCfg.DB.GetAllUsers(r.Context())
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, fmt.Sprintln("failed to retrieve all users"))
		return
	}
	RespondWithJSON(w, http.StatusOK, databaseUsersToUsers(users))
}
