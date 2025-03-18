package postgresrepository

import (
	"context"

	"gitlab.com/hospital-web/hospital-backend/doctor/domain"
	"gorm.io/gorm"
)

type DoctorPostgresRepository struct {
	db *gorm.DB
}

func NewDoctorPostgresRepository(repo *gorm.DB) *DoctorPostgresRepository {
	return &DoctorPostgresRepository{
		db: repo,
	}
}

func (u *DoctorPostgresRepository) CreateDoctor(ctx context.Context, doctor domain.Doctor) error {
	err := u.db.Create(&doctor).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *DoctorPostgresRepository) GetAllDoctors(ctx context.Context) ([]domain.Doctor, error) {

	var doctors []domain.Doctor
	err := u.db.Find(&doctors).Error
	if err != nil {
		return nil, err
	}
	return doctors, nil
}
