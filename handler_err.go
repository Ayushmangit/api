package main

import "net/http"

func HandlerError(w http.ResponseWriter, r *http.Request) {
	RespondWithError(w, 400, "internal server error")
}
