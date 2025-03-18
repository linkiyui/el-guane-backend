package domain

import (
	"context"

	paciente_domain "gitlab.com/hospital-web/hospital-backend/paciente/domain"
)

type ILeveRepository interface {
	CreateLeve(ctx context.Context, leve Leve) error
	GetLeves(ctx context.Context) ([]Leve, error)
	GetTotalLeves(ctx context.Context) (int64, error)
	GetAnalysisLeves(ctx context.Context) (int64, error)
	GetOldestLevePatientWithDiagnosis(ctx context.Context, diagnosis string) (paciente_domain.Paciente, error)
	GetLevesByDiagnosis(ctx context.Context, diagnosis string) ([]Leve_diagnosis, error)
	GetLeveInfo(ctx context.Context, leve_id string) (LeveToDomain, error)
}
