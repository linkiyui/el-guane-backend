package application

import (
	"context"

	domain_errors "gitlab.com/hospital-web/errors"
	"gitlab.com/hospital-web/utils"
)

func (u *UserService) ChangePassword(ctx context.Context, user_id string, password string) error {
	user, err := u.userRepo.FindByUsername(ctx, user_id)
	if err != nil {
		return err
	}
	if user.ID == "" {
		return domain_errors.ErrNotFound
	}
	hashedPassword := utils.HashingPasswordFunc(password)
	user.Password = hashedPassword
	err = u.userRepo.Update(ctx, user)
	if err != nil {
		return err
	}
	return nil
}
