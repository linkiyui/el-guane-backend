package postgresrepository

import (
	"context"

	"gitlab.com/hospital-web/hospital-backend/graves/domain"
	"gorm.io/gorm"
)

type GravePostgresRepository struct {
	db *gorm.DB
}

func NewGravePostgresRepository(repo *gorm.DB) *GravePostgresRepository {
	return &GravePostgresRepository{
		db: repo,
	}
}

func (u *GravePostgresRepository) CreateGrave(ctx context.Context, grave domain.Grave) error {
	err := u.db.Create(&grave).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *GravePostgresRepository) GetGraves(ctx context.Context) ([]domain.Grave, error) {

	var graves []domain.Grave
	err := u.db.Find(&graves).Error
	if err != nil {
		return nil, err
	}
	return graves, nil
}

func (u *GravePostgresRepository) GetGravesHighPressure(ctx context.Context) (int64, error) {

	var totalGraves int64
	err := u.db.Model(&domain.Grave{}).Where("press_min > ? AND press_max > ?", 90, 130).Count(&totalGraves).Error
	if err != nil {
		return 0, err
	}
	return totalGraves, nil
}
