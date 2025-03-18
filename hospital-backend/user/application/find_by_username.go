package application

import (
	"context"

	"gitlab.com/hospital-web/hospital-backend/user/domain"
)

func (u *UserService) FindByUsername(ctx context.Context, username string) (domain.User, error) {
	return u.userRepo.FindByUsername(ctx, username)
}
