package application

import (
	"context"

	domain "gitlab.com/hospital-web/hospital-backend/leves/domain"
)

func (u *LeveService) GetLeves(ctx context.Context) ([]domain.Leve, error) {
	return u.leve_repo.GetLeves(ctx)
}
