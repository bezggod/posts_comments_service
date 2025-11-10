package comment

import (
	"encoding/json"
	"net/http"
	"posts_comments_service/internal/controller"
	"posts_comments_service/internal/domain/models"
)

type reqCreateReply struct {
	PostID          int64  `json:"post_id"`
	UserID          int64  `json:"user_id"`
	ParentCommentID int64  `json:"parent_comment_id"`
	Text            string `json:"text"`
}

// CreateReply обрабатывает HTTP запрос на создание родительского комментария к посту
func (c *Controller) CreateReply(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req reqCreateReply

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		controller.WriteError(w, controller.ErrInvalidInput)
		return
	}
	if err := controller.ValidateComment(req.Text); err != nil {
		controller.WriteError(w, controller.ErrInvalidInput)
		return
	}
	if req.PostID <= 0 || req.UserID <= 0 || req.ParentCommentID < 0 {
		controller.WriteError(w, controller.ErrInvalidInput)
		return
	}

	comment, err := c.uc.CreateReplyComment(ctx, models.PostID(req.PostID), models.UserID(req.UserID), models.CommentID(req.ParentCommentID), req.Text)
	if err != nil {
		controller.WriteError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(comment)

}
