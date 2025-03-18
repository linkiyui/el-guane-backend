package application

import (
	user_service "gitlab.com/hospital-web/hospital-backend/user/application"
)

type AuthService struct {
	user_serv user_service.IUserService
}

func NewAuthService(user_serv user_service.IUserService) *AuthService {
	return &AuthService{
		user_serv: user_serv,
	}
}
