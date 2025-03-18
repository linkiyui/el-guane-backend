package application

import (
	"context"

	"gitlab.com/hospital-web/hospital-backend/consulta/domain"
)

func (u *ConsultaService) FindByConsultaId(ctx context.Context, consulta_id string) (domain.Consulta, error) {
	return u.consulta_repo.FindByConsultaId(ctx, consulta_id)
}
