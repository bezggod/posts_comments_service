package post

import (
	"encoding/json"
	"net/http"
	"posts_commets_service/internal/domain/models"
	"strconv"
)

func (c *Controller) ListPosts(w http.ResponseWriter, r *http.Request) {
	limitStr := r.URL.Query().Get("limit")
	lastStr := r.URL.Query().Get("last")

	limit, _ := strconv.Atoi(limitStr)

	var lastID *models.PostID
	if lastStr != "" {
		id, _ := strconv.ParseInt(lastStr, 10, 64)
		tmp := models.PostID(id)
		lastID = &tmp
	}

	posts, next, err := c.uc.ListPosts(r.Context(), limit, lastID)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}
	resp := map[string]interface{}{
		"posts":   posts,
		"next_id": next,
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(resp)
}
