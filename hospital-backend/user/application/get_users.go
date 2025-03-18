package application

import (
	"context"

	"gitlab.com/hospital-web/hospital-backend/user/domain"
)

func (u *UserService) GetUsers(ctx context.Context) ([]domain.User, error) {
	return u.userRepo.GetUsers(ctx)
}
