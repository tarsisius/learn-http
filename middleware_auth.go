package main

import (
	"net/http"

	"github.com/tarsisius/learn-http/db"
)

type AuthedHandler func(http.ResponseWriter, *http.Request, db.User)

func (apiConfig *ApiConfig) middlewareAuth(handler AuthedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := getApiKey(r.Header)
		if err != nil {
			responseWithError(w, 403, err.Error())
			return
		}

		user, err := apiConfig.DB.GetUserByApiKey(r.Context(), apiKey)
		if err != nil {
			responseWithError(w, 400, "Something went wrong when get user")
			return
		}

		handler(w, r, user)
	}
}
