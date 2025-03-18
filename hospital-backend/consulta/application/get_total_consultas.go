package application

import (
	"context"
)

func (u *ConsultaService) GetTotalConsultas(ctx context.Context) (int64, error) {
	return u.consulta_repo.GetTotalConsultas(ctx)
}
