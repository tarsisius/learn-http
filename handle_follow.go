package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/tarsisius/learn-http/db"
)

func (apiConfig *ApiConfig) handleFollow(w http.ResponseWriter, r *http.Request, user db.User) {
	decoder := json.NewDecoder(r.Body)

	type Parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}

	params := Parameters{}

	err := decoder.Decode(&params)
	if err != nil {
		responseWithError(w, 400, "Invalid request payload")
		return
	}

	follow, err := apiConfig.DB.CreateFollow(r.Context(), db.CreateFollowParams{
		ID:     uuid.New(),
		UserID: user.ID,
		FeedID: params.FeedID,
	})
	if err != nil {
		responseWithError(w, 500, "Something went wrong when create follow")
		return
	}

	respondWithJSON(w, 201, follow)
}

func (apiConfig *ApiConfig) handleGetFollows(w http.ResponseWriter, r *http.Request, user db.User) {
	follows, err := apiConfig.DB.GetFollows(r.Context(), user.ID)
	if err != nil {
		responseWithError(w, 500, "Something went wrong when get follows")
		return
	}

	respondWithJSON(w, 201, follows)
}

func (apiConfig *ApiConfig) handleDeleteFollow(w http.ResponseWriter, r *http.Request, user db.User) {
	followIDstr := chi.URLParam(r, "ID")

	followID, err := uuid.Parse(followIDstr)
	if err != nil {
		responseWithError(w, 400, "Error when parsing follow id")
		return
	}

	err = apiConfig.DB.DeleteFollow(r.Context(), db.DeleteFollowParams{
		ID:     followID,
		UserID: user.ID,
	})
	if err != nil {
		responseWithError(w, 500, "Something went wrong when delete follow")
		return
	}

	respondWithJSON(w, 200, map[string]string{"message": "success"})
}
