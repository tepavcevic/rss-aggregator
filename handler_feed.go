package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/tepavcevic/rss-aggregator/internal/database"
)

func (apiCfg *apiConfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	feed, err := apiCfg.DB.AddUserFeed(r.Context(), database.AddUserFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Url:       params.Url,
		UserID:    user.ID,
	})
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Couldn't create feed: %s", err))
		return
	}

	respondWithJSON(w, http.StatusCreated, dbFeedToFeed(feed))
}

func (apiCfg *apiConfig) handlerGetAllFeeds(w http.ResponseWriter, r *http.Request) {
	feeds, err := apiCfg.DB.GetAllFeed(r.Context())
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Couldn't get feeds: %s", err))
		return
	}

	respondWithJSON(w, http.StatusOK, dbFeedsToFeeds(feeds))
}
