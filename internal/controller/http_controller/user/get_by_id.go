package user

import (
	"encoding/json"
	"net/http"
	"posts_comments_service/internal/controller"
	"posts_comments_service/internal/domain/models"
	"strconv"
)

// GetBYID обрабатывает HTTP запрос на получение пользователя ID
func (c *Controller) GetByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	idQuery := r.URL.Query().Get("id")
	if idQuery == "" {
		controller.WriteError(w, controller.ErrInvalidInput)
		return
	}
	id, err := strconv.ParseInt(idQuery, 10, 64)
	if err != nil || id <= 0 {
		controller.WriteError(w, controller.ErrInvalidInput)
		return
	}

	user, err := c.uc.GetByID(ctx, models.UserID(id))
	if err != nil {
		controller.WriteError(w, controller.ErrNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(user)
}
