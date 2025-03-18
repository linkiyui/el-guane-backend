package application

import (
	"context"

	domain "gitlab.com/hospital-web/hospital-backend/paciente/domain"
	"gitlab.com/hospital-web/utils"
)

type IPacienteService interface {
	CreatePaciente(ctx context.Context, paciente domain.Paciente) error
	GetAllPacientes(ctx context.Context) ([]domain.Paciente, error)
}

type PacienteService struct {
	paciente_repo domain.IPacienteRepository
}

func NewPacienteService(paciente_repo domain.IPacienteRepository) *PacienteService {
	return &PacienteService{
		paciente_repo: paciente_repo,
	}
}

func (s *PacienteService) CreatePaciente(ctx context.Context, paciente domain.Paciente) error {
	paciente_id, err := utils.GenerateUUIDv7()
	if err != nil {
		return err
	}

	paciente.ID = paciente_id

	return s.paciente_repo.CreatePaciente(ctx, paciente)
}

func (s *PacienteService) GetAllPacientes(ctx context.Context) ([]domain.Paciente, error) {
	return s.paciente_repo.GetAllPacientes(ctx)
}
