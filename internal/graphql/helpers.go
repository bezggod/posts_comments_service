package graphql

import "posts_commets_service/internal/domain/models"

func boolOrFalse(p *bool) bool {
	if p == nil {
		return false
	}
	return *p
}

func toIntPtr(id *models.PostID) *int {
	if id == nil {
		return nil
	}
	v := int(*id)
	return &v
}
