package comment

import (
	"encoding/json"
	"net/http"
	"posts_commets_service/internal/domain/models"
)

type reqCreateReply struct {
	PostID          int64  `json:"post_id"`
	UserID          int64  `json:"user_id"`
	ParentCommentID int64  `json:"parent_comment_id"`
	Text            string `json:"text"`
}

func (c *Controller) CreateReply(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req reqCreateReply

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	if req.Text == "" {
		http.Error(w, "text is required", http.StatusBadRequest)
		return
	}

	comment, err := c.uc.CreateReplyComment(ctx, models.PostID(req.PostID), models.UserID(req.UserID), models.CommentID(req.ParentCommentID), req.Text)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(comment)

}
