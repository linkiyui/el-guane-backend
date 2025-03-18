package di_container

import (
	auth_service "gitlab.com/hospital-web/hospital-backend/auth/application"
)

func AuthService() *auth_service.AuthService {
	userSrv := UserService()
	return auth_service.NewAuthService(userSrv)

}
