package postgresrepository

import (
	"context"
	"fmt"

	domain_errors "gitlab.com/hospital-web/errors"
	"gitlab.com/hospital-web/hospital-backend/user/domain"
	"gorm.io/gorm"
)

type UserPostgresRepository struct {
	db *gorm.DB
}

func NewUserPostgresRepository(repo *gorm.DB) *UserPostgresRepository {
	return &UserPostgresRepository{
		db: repo,
	}
}

func (u *UserPostgresRepository) ExistsWithUserName(ctx context.Context, username string) (bool, error) {

	var user domain.User
	err := u.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		fmt.Println(err)
		return false, err
	}
	if user.ID == "" {
		return false, nil
	}
	return true, nil

}

func (u *UserPostgresRepository) CreateUser(ctx context.Context, user domain.User) error {
	err := u.db.Create(&user).Error
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (u *UserPostgresRepository) FindByUsername(ctx context.Context, username string) (domain.User, error) {

	var user domain.User
	err := u.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return domain.User{}, nil
		}
		fmt.Println(err)
		return domain.User{}, err
	}
	if user.ID == "" {
		return domain.User{}, nil
	}
	return user, nil
}

func (u *UserPostgresRepository) DeleteUser(ctx context.Context, user_id string) error {

	err := u.db.Where("id = ?", user_id).Delete(&domain.User{}).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return domain_errors.ErrNotFound
		}
		fmt.Println(err)
		return err
	}
	return nil
}

func (u *UserPostgresRepository) GetUsers(ctx context.Context) ([]domain.User, error) {

	var users []domain.User
	err := u.db.Where("role != ?", string(domain.RolAdmin)).Find(&users).Error
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return users, nil
}

func (u *UserPostgresRepository) Update(ctx context.Context, user domain.User) error {
	err := u.db.Save(&user).Error
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
