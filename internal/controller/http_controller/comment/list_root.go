package comment

import (
	"encoding/json"
	"net/http"
	"posts_commets_service/internal/domain/models"
	"strconv"
)

func (c *Controller) ListRoots(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	q := r.URL.Query()
	var last *models.CommentID

	postID, err := strconv.ParseInt(q.Get("post_id"), 10, 64)
	if err != nil || postID == 0 {
		http.Error(w, "bad request, post_id", http.StatusBadRequest)
		return
	}

	limit, err := strconv.Atoi(q.Get("limit"))
	if limit <= 0 {
		limit = 20
	}

	lastStr := q.Get("last_id")
	if lastStr != "" {
		var id int64
		id, err = strconv.ParseInt(lastStr, 10, 64)
		if err != nil {
			http.Error(w, "bad request, last_id", http.StatusBadRequest)
		}
		tmp := models.CommentID(id)
		last = &tmp
	}

	comments, next, err := c.uc.ListRoots(ctx, models.PostID(postID), limit, last)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]any{"comments": comments, "next_id": next})
}
