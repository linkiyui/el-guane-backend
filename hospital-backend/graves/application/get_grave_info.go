package application

import (
	"context"

	"gitlab.com/hospital-web/hospital-backend/graves/domain"
)

func (u *GraveService) GetGraveInfo(ctx context.Context, consulta_id string) (domain.GraveToDomain, error) {
	return u.grave_repo.GetGraveInfo(ctx, consulta_id)
}
