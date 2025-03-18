package application

import (
	"context"
)

func (u *UserService) ExistsWithUserName(ctx context.Context, username string) (bool, error) {
	return u.userRepo.ExistsWithUserName(ctx, username)
}
