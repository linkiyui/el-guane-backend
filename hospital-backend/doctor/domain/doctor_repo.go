package domain

import "context"

type IDoctorRepository interface {
	CreateDoctor(ctx context.Context, doctor Doctor) error
	GetAllDoctors(ctx context.Context) ([]Doctor, error)
}
