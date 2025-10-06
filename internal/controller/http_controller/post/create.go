package post

import (
	"encoding/json"
	"net/http"
	"posts_commets_service/internal/domain/models"
)

type reqCreatePost struct {
	UserID       int64  `json:"user_id"`
	Title        string `json:"title"`
	Body         string `json:"body"`
	CommentBlock bool   `json:"comment_block"`
}

func (c *Controller) Create(w http.ResponseWriter, r *http.Request) {
	var req reqCreatePost
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	if req.Title == "" || req.Body == "" {
		http.Error(w, "title and body is required", http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	post, err := c.uc.Create(ctx, models.UserID(req.UserID), req.Title, req.Body, req.CommentBlock)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(post)
}
