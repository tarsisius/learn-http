package main

import "net/http"

func handleError(w http.ResponseWriter, r *http.Request) {
	responseWithError(w, 400, "Something went wrong")
}

func handleHealth(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, 200, map[string]string{})
}
