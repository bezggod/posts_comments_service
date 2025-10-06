package post

import (
	"encoding/json"
	"net/http"
	"posts_commets_service/internal/domain/models"
)

type reqCommentBlock struct {
	PostID int64 `json:"post_id"`
	Off    bool  `json:"off"`
}

func (c *Controller) CommentBlock(w http.ResponseWriter, r *http.Request) {
	var req reqCommentBlock
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	err := c.uc.CommentBlock(r.Context(), models.PostID(req.PostID), req.Off)
	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(`{"status":"ok"}`))
}
