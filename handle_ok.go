package main

import "net/http"

func handleOk(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, 200, map[string]string{})
}
