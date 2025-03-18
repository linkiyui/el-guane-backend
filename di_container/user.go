package di_container

import (
	"gitlab.com/hospital-web/hospital-backend/database"
	user_service "gitlab.com/hospital-web/hospital-backend/user/application"
	user_infra "gitlab.com/hospital-web/hospital-backend/user/infra/postgres_repository"
)

func UserService() *user_service.UserService {
	repo := user_infra.NewUserPostgresRepository(database.Database)
	return user_service.NewUserService(repo)
}
