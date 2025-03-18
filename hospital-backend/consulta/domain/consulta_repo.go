package domain

import "context"

type IConsultaRepository interface {
	CreateConsulta(ctx context.Context, consulta Consulta) error
	FindByConsultaId(ctx context.Context, consulta_id string) (Consulta, error)
	GetTotalConsultas(ctx context.Context) (int64, error)
	GetAllConsultas(ctx context.Context) ([]Consulta, error)
	GetConsultaInfo(ctx context.Context) (Consulta, error)
}
