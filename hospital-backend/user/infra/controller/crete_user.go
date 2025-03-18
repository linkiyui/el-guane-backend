package controller

import (
	"github.com/gin-gonic/gin"
	di_container "gitlab.com/hospital-web/di_container"
	"gitlab.com/hospital-web/hospital-backend/user/domain"
)

type CreateUserRequest struct {
	CI       string `json:"ci"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

func CreateUser(ctx *gin.Context) {

	var (
		err error
	)

	var req CreateUserRequest
	if err = ctx.ShouldBindJSON(&req); err != nil {
		ctx.AbortWithStatus(400)
		return
	}

	user_service := di_container.UserService()

	user := domain.User{
		CI:       req.CI,
		Name:     req.Name,
		Username: req.Username,
		Password: req.Password,
		Role:     req.Role,
	}

	err = user_service.CreateUser(ctx, user)
	if err != nil {
		ctx.AbortWithStatus(500)
		return
	}

	ctx.AbortWithStatus(200)

}
