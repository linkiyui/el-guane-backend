package application

import (
	"context"

	domain "gitlab.com/hospital-web/hospital-backend/doctor/domain"
	"gitlab.com/hospital-web/utils"
)

type IDoctorService interface {
	CreateDoctor(ctx context.Context, doctor domain.Doctor) error
	GetAllDoctors(ctx context.Context) ([]domain.Doctor, error)
}

type DoctorService struct {
	doctor_repo domain.IDoctorRepository
}

func NewDoctorService(doctor_repo domain.IDoctorRepository) *DoctorService {
	return &DoctorService{
		doctor_repo: doctor_repo,
	}
}

func (s *DoctorService) CreateDoctor(ctx context.Context, doctor domain.Doctor) error {
	doctor_id, err := utils.GenerateUUIDv7()
	if err != nil {
		return err
	}

	doctor.ID = doctor_id
	return s.doctor_repo.CreateDoctor(ctx, doctor)
}

func (s *DoctorService) GetAllDoctors(ctx context.Context) ([]domain.Doctor, error) {
	return s.doctor_repo.GetAllDoctors(ctx)
}
