package post

import (
	"encoding/json"
	"net/http"
	"posts_comments_service/internal/controller"
	"posts_comments_service/internal/domain/models"
)

type reqCommentBlock struct {
	PostID int64 `json:"post_id"`
	Off    bool  `json:"off"`
}

func (c *Controller) CommentBlock(w http.ResponseWriter, r *http.Request) {
	var req reqCommentBlock
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		controller.WriteError(w, controller.ErrInvalidInput)
		return
	}
	if req.PostID <= 0 {
		controller.WriteError(w, controller.ErrInvalidInput)
		return
	}

	err := c.uc.CommentBlock(r.Context(), models.PostID(req.PostID), req.Off)
	if err != nil {
		controller.WriteError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(`{"status":"ok"}`))
}
