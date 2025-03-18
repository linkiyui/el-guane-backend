package application

import (
	"context"

	user_domain "gitlab.com/hospital-web/hospital-backend/user/domain"
)

type IUserService interface {
	ExistsWithUserName(ctx context.Context, username string) (bool, error)
	CreateUser(ctx context.Context, user user_domain.User) error
	FindByUsername(ctx context.Context, username string) (user_domain.User, error)
	DeleteUser(ctx context.Context, user_id string) error
	GetUsers(ctx context.Context) ([]user_domain.User, error)
	ChangePassword(ctx context.Context, user_id string, password string) error
	GetRole() []user_domain.Rol
	GetUser(ctx context.Context, user_id string) (user_domain.User, error)
	GetMyInfo(ctx context.Context, user_id string) (user_domain.User, error)
}

type UserService struct {
	userRepo user_domain.IUserRepository
}

func NewUserService(userRepo user_domain.IUserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (s *UserService) GetRole() []user_domain.Rol {
	return []user_domain.Rol{user_domain.RolAdmin, user_domain.Roluser}
}

func (s *UserService) GetUser(ctx context.Context, user_id string) (user_domain.User, error) {
	return s.userRepo.GetUser(ctx, user_id)
}
