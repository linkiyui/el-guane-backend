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

func (u *GravePostgresRepository) GetGraveInfo(ctx context.Context, consulta_id string) (domain.GraveToDomain, error) {

	var result domain.GraveToDomain
	err := u.db.Table("graves").
		Select("graves.*, consultas.*, doctors.name as doctor_name, pacientes.name as paciente_name, pacientes.age as paciente_age, pacientes.sex as paciente_sex").
		Joins("JOIN consultas ON graves.consulta_id = consultas.id").
		Joins("JOIN doctors ON consultas.doctor_id = doctors.id").
		Joins("JOIN pacientes ON consultas.paciente_id = pacientes.id").
		Where("graves.consulta_id = ?", consulta_id).
		First(&result).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return domain.GraveToDomain{}, nil
		}
		return domain.GraveToDomain{}, err
	}
	return result, nil
}
