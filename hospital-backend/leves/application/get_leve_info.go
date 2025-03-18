package application

import (
	"context"

	"gitlab.com/hospital-web/hospital-backend/leves/domain"
)

func (u *LeveService) GetLeveInfo(ctx context.Context, consulta_id string) (domain.LeveToDomain, error) {
	return u.leve_repo.GetLeveInfo(ctx, consulta_id)
}
