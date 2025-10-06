package user

import (
	"encoding/json"
	"net/http"
)

type reqCreateUser struct {
	Name string `json:"name"`
}

func (c *Controller) Create(w http.ResponseWriter, r *http.Request) {
	var req reqCreateUser
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	if req.Name == "" {
		http.Error(w, "name is required", http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	user, err := c.uc.Create(ctx, req.Name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(user)
}
