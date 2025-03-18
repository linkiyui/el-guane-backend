package di_container

import (
	"gitlab.com/hospital-web/hospital-backend/database"
	paciente_service "gitlab.com/hospital-web/hospital-backend/paciente/application"
	paciente_infra "gitlab.com/hospital-web/hospital-backend/paciente/infra/postgres_repository"
)

func PacienteService() *paciente_service.PacienteService {
	repo := paciente_infra.NewPacientePostgresRepository(database.Database)
	return paciente_service.NewPacienteService(repo)
}
