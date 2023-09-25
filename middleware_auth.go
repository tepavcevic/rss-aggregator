package main

import (
	"fmt"
	"net/http"

	"github.com/tepavcevic/rss-aggregator/internal/auth"
	"github.com/tepavcevic/rss-aggregator/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetApiKey(r.Header)
		if err != nil {
			respondWithError(w, http.StatusUnauthorized, fmt.Sprintf("Unauthorized: %v", err))
			return
		}

		user, err := apiCfg.DB.GetUserByApiKey(r.Context(), apiKey)
		if err != nil {
			respondWithError(w, http.StatusNotFound, "User not found")
			return
		}

		handler(w, r, user)
	}
}
