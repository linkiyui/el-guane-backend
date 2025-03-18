package application

import (
	"context"

	"gitlab.com/hospital-web/hospital-backend/user/domain"
)

func (u *UserService) CreateUser(ctx context.Context, user domain.User) error {
	return u.userRepo.CreateUser(ctx, user)
}
