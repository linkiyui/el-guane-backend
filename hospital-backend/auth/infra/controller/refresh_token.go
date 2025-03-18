package controller

import (
	"strings"

	"github.com/gin-gonic/gin"
	"gitlab.com/hospital-web/hospital-backend/auth/infra/auth_token"
)

func RefreshToken(ctx *gin.Context) {
	tkn := ctx.GetHeader("Authorization")
	if tkn == "" || !strings.HasPrefix(tkn, "Bearer ") {
		ctx.AbortWithStatus(401)
		return
	}
	claims, err := auth_token.ValidateRefreshToken(strings.TrimPrefix(tkn, "Bearer "))
	if err != nil {
		ctx.AbortWithStatus(401)
		return
	}

	var token, refreshToken string
	if token, err = auth_token.GenerateLoginToken(claims.UserID); err != nil {
		ctx.AbortWithStatus(500)
		return
	}
	if refreshToken, err = auth_token.GenerateRefreshToken(claims.UserID); err != nil {
		ctx.AbortWithStatus(500)
		return
	}
	ctx.AbortWithStatusJSON(200, gin.H{
		"token":         token,
		"refresh_token": refreshToken,
	})
}
