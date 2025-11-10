package user

import (
	"encoding/json"
	"net/http"
	"posts_comments_service/internal/controller"
)

type reqCreateUser struct {
	Name string `json:"name"`
}

// Create обрабатывает HTTP запрос на создание нового пользователя
func (c *Controller) Create(w http.ResponseWriter, r *http.Request) {
	var req reqCreateUser
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		controller.WriteError(w, controller.ErrInvalidInput)
		return
	}
	if err := controller.ValidateUser(req.Name); err != nil {
		controller.WriteError(w, controller.ErrInvalidInput)
		return
	}

	ctx := r.Context()
	user, err := c.uc.Create(ctx, req.Name)
	if err != nil {
		controller.WriteError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(user)
}
