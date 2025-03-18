package application

import (
	"context"

	domain_errors "gitlab.com/hospital-web/errors"
	"gitlab.com/hospital-web/hospital-backend/user/domain"
	utils "gitlab.com/hospital-web/utils"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (s *AuthService) Login(req LoginRequest) (*domain.User, error) {

	ctx := context.Background()

	user, err := s.user_serv.FindByUsername(ctx, req.Username)
	if err != nil {
		return nil, err
	}
	if user.ID == "" {
		return nil, domain_errors.ErrNotFound
	}

	if !utils.CheckPasswordHashFunc(req.Password, user.Password) {
		return nil, domain_errors.ErrInvalidPassword
	}

	resp := &domain.User{
		CI:       user.CI,
		Name:     user.Name,
		Username: user.Username,

		Role: user.Role,
	}

	return resp, err

}
