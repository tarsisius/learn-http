package main

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/tarsisius/learn-http/db"
)

func (apiConfig *ApiConfig) handleCreateUser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	type Parameters struct {
		Name string `json:"name"`
	}

	params := Parameters{}

	err := decoder.Decode(&params)
	if err != nil {
		responseWithError(w, 400, "Invalid request payload")
		return
	}

	user, err := apiConfig.DB.CreateUser(r.Context(), db.CreateUserParams{
		ID:   uuid.New(),
		Name: params.Name,
	})
	if err != nil {
		responseWithError(w, 500, "Something went wrong when create user")
		return
	}

	respondWithJSON(w, 201, user)
}

func (apiConfig *ApiConfig) handleGetUser(w http.ResponseWriter, r *http.Request, user db.User) {
	respondWithJSON(w, 200, user)
}
