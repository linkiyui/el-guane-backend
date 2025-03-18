package application

import (
	"context"

	domain "gitlab.com/hospital-web/hospital-backend/leves/domain"
)

type ILeveService interface {
	CreateLeve(ctx context.Context, leve domain.Leve) error
}

type LeveService struct {
	leve_repo domain.ILeveRepository
}

func NewLeveService(leve_repo domain.ILeveRepository) *LeveService {
	return &LeveService{
		leve_repo: leve_repo,
	}
}

func (s *LeveService) CreateLeve(ctx context.Context, leve domain.Leve) error {
	return s.leve_repo.CreateLeve(ctx, leve)
}
