package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	di_container "gitlab.com/hospital-web/di_container"
)

type ChangePasswordRequest struct {
	Password string `json:"password"`
}

func ChangePassword(ctx *gin.Context) {

	var (
		err     error
		uid     any
		user_id string
		exists  bool
	)

	uid, exists = ctx.Get("user_id")
	if !exists {
		ctx.AbortWithStatus(401)
		return
	}

	user_id = fmt.Sprint(uid)

	var req ChangePasswordRequest
	if err = ctx.ShouldBindJSON(&req); err != nil {
		ctx.AbortWithStatus(400)
		return
	}

	user_service := di_container.UserService()

	err = user_service.ChangePassword(ctx, user_id, req.Password)
	if err != nil {
		ctx.AbortWithStatus(500)
		return
	}

	ctx.AbortWithStatus(200)

}
