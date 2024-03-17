package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Deepjyoti-Sarmah/RSS_Aggregator/internal/database"
)

func (apiCfg *apiConfig) handlerPostsGet(w http.ResponseWriter, r *http.Request, user database.User)  {
	limitStr := r.URL.Query().Get("limit")
	limit := 10
	if specifiedLimit, err := strconv.Atoi(limitStr); err == nil {
		limit = specifiedLimit
	}

	posts, err := apiCfg.DB.GetPostsForUser(r.Context(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit: int32(limit),
	})

	if err != nil {
		responseWithError(w, http.StatusInternalServerError, fmt.Sprintf("Couldn't get posts for user %v:", err))
		return
	}

	responseWithJSON(w, http.StatusOK, databasePostsToPosts(posts))
}
