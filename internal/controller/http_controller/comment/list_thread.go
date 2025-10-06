package comment

import (
	"encoding/json"
	"net/http"
	"posts_commets_service/internal/domain/models"
	"strconv"
)

func (c *Controller) ListThread(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	ctx := r.Context()
	var last *models.CommentID

	postID, err := strconv.ParseInt(q.Get("post_id"), 10, 64)
	if err != nil {
		http.Error(w, "bad post_id", http.StatusBadRequest)
		return
	}
	rootID, err := strconv.ParseInt(q.Get("root_id"), 10, 64)
	if err != nil {
		http.Error(w, "bad root_id", http.StatusBadRequest)
		return
	}
	limit, _ := strconv.Atoi(q.Get("limit"))
	if limit <= 0 {
		limit = 20
	}

	lastStr := q.Get("last_id")
	if lastStr != "" {
		var id int64
		id, err = strconv.ParseInt(lastStr, 10, 64)
		if err != nil {
			http.Error(w, "bad request, last_id", http.StatusBadRequest)
			return
		}
		tmp := models.CommentID(id)
		last = &tmp
	}

	comments, next, err := c.uc.ListThreads(ctx, models.PostID(postID), models.CommentID(rootID), limit, last)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]any{"comments": comments, "next_id": next})

}
