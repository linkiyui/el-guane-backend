package di_container

import (
	consulta_service "gitlab.com/hospital-web/hospital-backend/consulta/application"
	consulta_infra "gitlab.com/hospital-web/hospital-backend/consulta/infra/postgres_repository"
	"gitlab.com/hospital-web/hospital-backend/database"
)

func ConsultaService() *consulta_service.ConsultaService {
	repo := consulta_infra.NewConsultaPostgresRepository(database.Database)
	return consulta_service.NewConsultaService(repo)
}
