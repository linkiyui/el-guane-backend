package application

import (
	"context"

	domain "gitlab.com/hospital-web/hospital-backend/consulta/domain"
	"gitlab.com/hospital-web/utils"
)

type IConsultaService interface {
	CreateConsulta(ctx context.Context, consulta domain.Consulta) (string, error)
	GetConsultaInfo(ctx context.Context, consulta_id string) (domain.Consulta, error)
	GetTotalConsultas(ctx context.Context) (int64, error)
	GetAllConsultas(ctx context.Context) ([]domain.Consulta, error)
}

type ConsultaService struct {
	consulta_repo domain.IConsultaRepository
}

func NewConsultaService(consulta_repo domain.IConsultaRepository) *ConsultaService {
	return &ConsultaService{
		consulta_repo: consulta_repo,
	}
}

func (s *ConsultaService) CreateConsulta(ctx context.Context, consulta domain.Consulta) (string, error) {
	ID, err := utils.GenerateUUIDv7()
	if err != nil {
		return "", err
	}
	con := domain.Consulta{
		ID:          ID,
		Doctor_id:   consulta.Doctor_id,
		Date:        consulta.Date,
		Paciente_id: consulta.Paciente_id,
		Grade:       consulta.Grade,
	}
	return s.consulta_repo.CreateConsulta(ctx, con)
}
