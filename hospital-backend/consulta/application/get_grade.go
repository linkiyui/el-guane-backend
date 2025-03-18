package application

import (
	"context"

	"gitlab.com/hospital-web/hospital-backend/consulta/domain"
)

func (u *ConsultaService) Getgrade(ctx context.Context) ([]domain.Grade, error) {
	return []domain.Grade{domain.GradeLow, domain.GradeHigh}, nil
}
