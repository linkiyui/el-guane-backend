package application

import (
	"context"
)

func (u *UserService) DeleteUser(ctx context.Context, user_id string) error {
	return u.userRepo.DeleteUser(ctx, user_id)
}
