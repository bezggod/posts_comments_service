package post

import (
	"encoding/json"
	"net/http"
	"posts_comments_service/internal/controller"
	"posts_comments_service/internal/domain/models"
)

type reqCreatePost struct {
	UserID       int64  `json:"user_id"`
	Title        string `json:"title"`
	Body         string `json:"body"`
	CommentBlock bool   `json:"comment_block"`
}

// Create обрабатывает HTTP запрос на создание нового поста
func (c *Controller) Create(w http.ResponseWriter, r *http.Request) {
	var req reqCreatePost
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		controller.WriteError(w, controller.ErrInvalidInput)
		return
	}
	if err := controller.ValidatePost(req.Title, req.Body); err != nil {
		controller.WriteError(w, controller.ErrInvalidInput)
		return
	}

	ctx := r.Context()
	post, err := c.uc.Create(ctx, models.UserID(req.UserID), req.Title, req.Body, req.CommentBlock)
	if err != nil {
		controller.WriteError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(post)
}
