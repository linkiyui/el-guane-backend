package application

import (
	"context"

	"gitlab.com/hospital-web/hospital-backend/consulta/domain"
)

func (u *ConsultaService) GetConsultaInfo(ctx context.Context) (domain.Consulta, error) {
	return u.consulta_repo.GetConsultaInfo(ctx)
}
