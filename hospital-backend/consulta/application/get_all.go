package application

import (
	"context"

	"gitlab.com/hospital-web/hospital-backend/consulta/domain"
)

func (u *ConsultaService) GetAllConsultas(ctx context.Context) ([]domain.Consulta, error) {
	return u.consulta_repo.GetAllConsultas(ctx)
}
