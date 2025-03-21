package domain

import "context"

type IUserRepository interface {
	ExistsWithUserName(ctx context.Context, username string) (bool, error)
	CreateUser(ctx context.Context, user User) error
	DeleteUser(ctx context.Context, user_id string) error
	FindByUsername(ctx context.Context, username string) (User, error)
	GetUsers(ctx context.Context) ([]User, error)
	Update(ctx context.Context, user User) error
	ChangePassword(ctx context.Context, user_id string, password string) error
	GetUser(ctx context.Context, user_id string) (User, error)
	GetMyInfo(ctx context.Context, user_id string) (User, error)
}
