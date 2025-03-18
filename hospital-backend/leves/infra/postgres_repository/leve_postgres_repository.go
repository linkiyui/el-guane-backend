package postgresrepository

import (
	"context"

	domain_errors "gitlab.com/hospital-web/errors"
	"gitlab.com/hospital-web/hospital-backend/leves/domain"
	paciente_domain "gitlab.com/hospital-web/hospital-backend/paciente/domain"
	"gorm.io/gorm"
)

type LevePostgresRepository struct {
	db *gorm.DB
}

func NewLevePostgresRepository(repo *gorm.DB) *LevePostgresRepository {
	return &LevePostgresRepository{
		db: repo,
	}
}

func (u *LevePostgresRepository) CreateLeve(ctx context.Context, leve domain.Leve) error {
	err := u.db.Create(&leve).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *LevePostgresRepository) GetLeves(ctx context.Context) ([]domain.Leve, error) {

	var leves []domain.Leve
	err := u.db.Find(&leves).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, domain_errors.ErrNotFound
		}
		return nil, err
	}
	return leves, nil
}

func (u *LevePostgresRepository) GetTotalLeves(ctx context.Context) (int64, error) {

	var totalLeves int64
	err := u.db.Model(&domain.Leve{}).Count(&totalLeves).Error
	if err != nil {
		return 0, err
	}
	return totalLeves, nil
}

func (u *LevePostgresRepository) GetAnalysisLeves(ctx context.Context) (int64, error) {

	var analysisLeves int64
	err := u.db.Model(&domain.Leve{}).Where("analisis = ?", true).Count(&analysisLeves).Error
	if err != nil {
		return 0, err
	}
	return analysisLeves, nil
}

func (u *LevePostgresRepository) GetOldestLevePatientWithDiagnosis(ctx context.Context, diagnosis string) (paciente_domain.Paciente, error) {
	var paciente paciente_domain.Paciente
	err := u.db.Table("leves").
		Select("pacientes.*").
		Joins("JOIN pacientes ON leves.consulta_id = pacientes.id").
		Where("leves.diagnosis = ?", diagnosis).
		Order("pacientes.age DESC").
		First(&paciente).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return paciente_domain.Paciente{}, nil
		}
		return paciente_domain.Paciente{}, err
	}
	return paciente, nil
}

func (u *LevePostgresRepository) GetLevesByDiagnosis(ctx context.Context, diagnosis string) ([]domain.Leve_diagnosis, error) {

	var leves []domain.Leve_diagnosis
	err := u.db.Table("leves").
		Select("consultas.id as consulta_id, consultas.doctor_id as doctor_id, consultas.date as date, pacientes.name as paciente_name, pacientes.age as paciente_age, pacientes.sex as paciente_sex, leves.diagnosis as diagnosis, leves.analisis as analisis").
		Joins("JOIN consultas ON leves.consulta_id = consultas.id").
		Joins("JOIN pacientes ON consultas.paciente_id = pacientes.id").
		Where("leves.diagnosis = ?", diagnosis).
		Order("consultas.date ASC").
		Find(&leves).Error
	if err != nil {
		return nil, err
	}
	return leves, nil
}
