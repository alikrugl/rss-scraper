package main

import (
	"net/http"
	"strconv"

	"github.com/alikrugl/rss-scraper/internal/database"
)

func (cfg *apiConfig) handlerPostsGet(w http.ResponseWriter, r *http.Request, user database.User) {
	limitStr := r.URL.Query().Get("limit")
	limit := 10
	if specifiedLimit, err := strconv.Atoi(limitStr); err == nil {
		limit = specifiedLimit
	}

	posts, err := cfg.DB.GetUserPosts(r.Context(), database.GetUserPostsParams{
		UserID: user.ID,
		Limit:  int32(limit),
	})

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't get posts")
		return
	}

	respondWithJSON(w, http.StatusOK, posts)
}
