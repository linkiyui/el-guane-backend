package controller

import (
	"github.com/gin-gonic/gin"
	di_container "gitlab.com/hospital-web/di_container"
	"gitlab.com/hospital-web/errors"
	auth_serv "gitlab.com/hospital-web/hospital-backend/auth/application"
	auth_token "gitlab.com/hospital-web/hospital-backend/auth/infra/auth_token"
)

type SignupRequest struct {
	CI       string `json:"ci"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

func Signup(ctx *gin.Context) {

	var (
		err error
	)

	var req SignupRequest
	if err = ctx.ShouldBindBodyWithJSON(&req); err != nil {
		ctx.AbortWithStatus(400)
		return
	}

	auth_service := di_container.AuthService()

	resp, err := auth_service.Signup(auth_serv.SignupRequest{
		CI:       req.CI,
		Name:     req.Name,
		Username: req.Username,
		Password: req.Password,
		Role:     req.Role,
	})
	if err != nil {
		ctx.AbortWithStatus(500)
		return
	}

	token, err := auth_token.GenerateLoginToken(resp.ID)
	if err != nil {
		de := errors.IsDomainError(err)
		if de != nil {
			ctx.AbortWithStatusJSON(de.Code, gin.H{"error": de.Message})
			return
		}
		ctx.AbortWithStatus(500)
		return

	}

	refresh_token, err := auth_token.GenerateRefreshToken(resp.ID)
	if err != nil {
		de := errors.IsDomainError(err)
		if de != nil {
			ctx.AbortWithStatusJSON(de.Code, gin.H{"error": de.Message})
			return
		}
		ctx.AbortWithStatus(500)
		return
	}

	ctx.AbortWithStatusJSON(200, gin.H{
		"token":         token,
		"refresh_token": refresh_token,
		"user":          resp,
	})

}
