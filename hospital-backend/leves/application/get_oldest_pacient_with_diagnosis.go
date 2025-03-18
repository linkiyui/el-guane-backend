package application

import (
	"context"

	paciente_domain "gitlab.com/hospital-web/hospital-backend/paciente/domain"
)

func (u *LeveService) GetOldestLevePatientWithDiagnosis(ctx context.Context, diagnosis string) (paciente_domain.Paciente, error) {
	return u.leve_repo.GetOldestLevePatientWithDiagnosis(ctx, diagnosis)
}
