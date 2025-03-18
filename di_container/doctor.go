package di_container

import (
	"gitlab.com/hospital-web/hospital-backend/database"
	doctor_service "gitlab.com/hospital-web/hospital-backend/doctor/application"
	doctor_infra "gitlab.com/hospital-web/hospital-backend/doctor/infra/postgres_repository"
)

func DoctorService() *doctor_service.DoctorService {
	repo := doctor_infra.NewDoctorPostgresRepository(database.Database)
	return doctor_service.NewDoctorService(repo)
}
