package comment

import (
	"encoding/json"
	"net/http"
	"posts_commets_service/internal/domain/models"
	"strconv"
)

func (c *Controller) GetByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	idQuery := r.URL.Query().Get("id")
	if idQuery == "" {
		http.Error(w, "id is required", http.StatusBadRequest)
	}
	id, err := strconv.ParseInt(idQuery, 10, 64)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
	}

	user, err := c.uc.GetByID(ctx, models.CommentID(id))
	if err != nil {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(user)
}
