package post

import (
	"encoding/json"
	"net/http"
	"posts_comments_service/internal/controller"
	"posts_comments_service/internal/domain/models"
	"strconv"
)

// ListPosts обрабатывает HTTP запрос на получение списка постов с пагинацией
func (c *Controller) ListPosts(w http.ResponseWriter, r *http.Request) {
	limitStr := r.URL.Query().Get("limit")
	lastStr := r.URL.Query().Get("last_id")

	limit, err := strconv.Atoi(limitStr)
	if limitStr != "" && err != nil {
		controller.WriteError(w, controller.ErrInvalidInput)
		return
	}

	var lastID *models.PostID
	if lastStr != "" {
		id, err := strconv.ParseInt(lastStr, 10, 64)
		if err != nil {
			controller.WriteError(w, controller.ErrInvalidInput)
			return
		}
		tmp := models.PostID(id)
		lastID = &tmp
	}

	posts, next, err := c.uc.ListPosts(r.Context(), limit, lastID)
	if err != nil {
		controller.WriteError(w, err)
		return
	}
	resp := map[string]interface{}{
		"posts":   posts,
		"next_id": next,
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(resp)
}
