package application

import (
	"context"

	domain "gitlab.com/hospital-web/hospital-backend/graves/domain"
)

type IGraveService interface {
	CreateGrave(ctx context.Context, grave domain.Grave) error
}

type GraveService struct {
	grave_repo domain.IGraveRepository
}

func NewGraveService(grave_repo domain.IGraveRepository) *GraveService {
	return &GraveService{
		grave_repo: grave_repo,
	}
}

func (s *GraveService) CreateGrave(ctx context.Context, grave domain.Grave) error {
	return s.grave_repo.CreateGrave(ctx, grave)
}
