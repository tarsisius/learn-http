package main

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/tarsisius/learn-http/db"
)

func (apiConfig *ApiConfig) handleCreateFeed(w http.ResponseWriter, r *http.Request, user db.User) {
	decoder := json.NewDecoder(r.Body)

	type Parameters struct {
		Name   string    `json:"name"`
		Url    string    `json:"url"`
		UserID uuid.UUID `json:"user_id"`
	}

	params := Parameters{}

	err := decoder.Decode(&params)
	if err != nil {
		responseWithError(w, 400, "Invalid request payload")
		return
	}

	feed, err := apiConfig.DB.CreateFeed(r.Context(), db.CreateFeedParams{
		ID:     uuid.New(),
		Name:   params.Name,
		Url:    params.Url,
		UserID: user.ID,
	})
	if err != nil {
		responseWithError(w, 500, "Something went wrong when create feed")
		return
	}

	respondWithJSON(w, 201, feed)
}

func (apiConfig *ApiConfig) handleGetFeeds(w http.ResponseWriter, r *http.Request) {
	feeds, err := apiConfig.DB.GetFeeds(r.Context())
	if err != nil {
		responseWithError(w, 500, "Something went wrong when getting feeds")
		return
	}

	respondWithJSON(w, 201, feeds)
}
