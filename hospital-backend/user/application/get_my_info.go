package application

import (
	"context"

	"gitlab.com/hospital-web/hospital-backend/user/domain"
)

func (u *UserService) GetMyInfo(ctx context.Context, user_id string) (domain.User, error) {
	return u.userRepo.GetMyInfo(ctx, user_id)
}
