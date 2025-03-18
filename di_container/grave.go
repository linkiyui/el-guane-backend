package di_container

import (
	"gitlab.com/hospital-web/hospital-backend/database"
	grave_service "gitlab.com/hospital-web/hospital-backend/graves/application"
	grave_infra "gitlab.com/hospital-web/hospital-backend/graves/infra/postgres_repository"
)

func GraveService() *grave_service.GraveService {
	repo := grave_infra.NewGravePostgresRepository(database.Database)
	return grave_service.NewGraveService(repo)
}
