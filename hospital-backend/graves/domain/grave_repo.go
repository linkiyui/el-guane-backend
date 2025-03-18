package domain

import "context"

type IGraveRepository interface {
	CreateGrave(ctx context.Context, grave Grave) error
	GetGraves(ctx context.Context) ([]Grave, error)
	GetGravesHighPressure(ctx context.Context) (int64, error)
	GetGraveInfo(ctx context.Context, consulta_id string) (GraveToDomain, error)
}
