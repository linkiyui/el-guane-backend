package application

import (
	"context"

	domain_errors "gitlab.com/hospital-web/errors"
	"gitlab.com/hospital-web/hospital-backend/user/domain"
	utils "gitlab.com/hospital-web/utils"
)

type SignupRequest struct {
	CI       string `json:"ci"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

func (s *AuthService) Signup(req SignupRequest) (*domain.User, error) {

	ctx := context.Background()

	exists, err := s.user_serv.ExistsWithUserName(ctx, req.Username)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, domain_errors.ErrExistsUsername
	}

	hashedPassword := utils.HashingPasswordFunc(req.Password)
	user_id, err := utils.GenerateUUIDv7()
	if err != nil {
		return nil, err
	}
	user := domain.User{
		ID:       user_id,
		CI:       req.CI,
		Name:     req.Name,
		Username: req.Username,
		Password: hashedPassword,
		Role:     req.Role,
	}

	err = s.user_serv.CreateUser(ctx, user)

	return &user, err

}
