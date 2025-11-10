package comment

import (
	"encoding/json"
	"net/http"
	"posts_comments_service/internal/controller"
	"posts_comments_service/internal/domain/models"
)

type reqCreateRoot struct {
	PostID int64  `json:"post_id"`
	UserID int64  `json:"user_id"`
	Text   string `json:"text"`
}

// CreateRoot обрабатывает HTTP запрос на создание ответа на комментарий
func (c *Controller) CreateRoot(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req reqCreateRoot

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		controller.WriteError(w, controller.ErrInvalidInput)
		return
	}
	if err := controller.ValidateComment(req.Text); err != nil {
		controller.WriteError(w, controller.ErrInvalidInput)
		return
	}
	if req.PostID <= 0 || req.UserID <= 0 {
		controller.WriteError(w, controller.ErrInvalidInput)
	}

	comment, err := c.uc.CreateRoot(ctx, models.PostID(req.PostID), models.UserID(req.UserID), req.Text)
	if err != nil {
		controller.WriteError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(comment)

}
