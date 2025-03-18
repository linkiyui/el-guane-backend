package controller

import (
	"github.com/gin-gonic/gin"
	di_container "gitlab.com/hospital-web/di_container"
)

func GetUsers(ctx *gin.Context) {

	user_service := di_container.UserService()
	users, err := user_service.GetUsers(ctx)
	if err != nil {
		ctx.AbortWithStatus(500)
		return
	}

	users_resp := make([]UserResponse, 0, len(users))
	for _, u := range users {
		user_resp := UserResponse{
			ID:       u.ID,
			CI:       u.CI,
			Name:     u.Name,
			Username: u.Username,
			Role:     u.Role,
		}
		users_resp = append(users_resp, user_resp)
	}

	ctx.AbortWithStatusJSON(200, gin.H{
		"users": users_resp,
	})

}

type UserResponse struct {
	ID       string `json:"id"`
	CI       string `json:"ci"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Role     string `json:"role"`
}
