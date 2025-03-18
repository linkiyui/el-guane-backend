package application

import (
	"context"

	domain "gitlab.com/hospital-web/hospital-backend/graves/domain"
)

func (u *GraveService) GetGraves(ctx context.Context) ([]domain.Grave, error) {
	return u.grave_repo.GetGraves(ctx)
}
