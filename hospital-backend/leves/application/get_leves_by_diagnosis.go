package application

import (
	"context"

	domain "gitlab.com/hospital-web/hospital-backend/leves/domain"
)

func (u *LeveService) GetLevesByDiagnosis(ctx context.Context, diagnosis string) ([]domain.Leve_diagnosis, error) {
	return u.leve_repo.GetLevesByDiagnosis(ctx, diagnosis)
}
