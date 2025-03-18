package di_container

import (
	"gitlab.com/hospital-web/hospital-backend/database"
	leve_service "gitlab.com/hospital-web/hospital-backend/leves/application"
	leve_infra "gitlab.com/hospital-web/hospital-backend/leves/infra/postgres_repository"
)

func LeveService() *leve_service.LeveService {
	repo := leve_infra.NewLevePostgresRepository(database.Database)
	return leve_service.NewLeveService(repo)
}
