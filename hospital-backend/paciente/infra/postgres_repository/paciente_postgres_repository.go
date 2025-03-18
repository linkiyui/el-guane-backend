package postgresrepository

import (
	"context"

	"gitlab.com/hospital-web/hospital-backend/paciente/domain"
	"gorm.io/gorm"
)

type PacientePostgresRepository struct {
	db *gorm.DB
}

func NewPacientePostgresRepository(repo *gorm.DB) *PacientePostgresRepository {
	return &PacientePostgresRepository{
		db: repo,
	}
}

func (u *PacientePostgresRepository) CreatePaciente(ctx context.Context, paciente domain.Paciente) error {
	err := u.db.Create(&paciente).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *PacientePostgresRepository) GetAllPacientes(ctx context.Context) ([]domain.Paciente, error) {

	var pacientes []domain.Paciente
	err := u.db.Find(&pacientes).Error
	if err != nil {
		return nil, err
	}
	return pacientes, nil
}
