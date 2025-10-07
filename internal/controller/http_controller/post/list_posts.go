package post

import (
	"encoding/json"
	"net/http"
	"posts_commets_service/internal/domain/models"
	"strconv"
)

func (c *Controller) ListPosts(w http.ResponseWriter, r *http.Request) {
	limitStr := r.URL.Query().Get("limit")
	lastStr := r.URL.Query().Get("last_id")

	limit, err := strconv.Atoi(limitStr)
	if limitStr != "" && err != nil {
		http.Error(w, "limit must be a number", http.StatusBadRequest)
	}

	var lastID *models.PostID
	if lastStr != "" {
		id, err := strconv.ParseInt(lastStr, 10, 64)
		if err != nil {
			http.Error(w, "last_id must be a number", http.StatusBadRequest)
			return
		}
		tmp := models.PostID(id)
		lastID = &tmp
	}

	posts, next, err := c.uc.ListPosts(r.Context(), limit, lastID)
	if err != nil {
		http.Error(w, "list error", http.StatusBadRequest)
		return
	}
	resp := map[string]interface{}{
		"posts":   posts,
		"next_id": next,
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(resp)
}
