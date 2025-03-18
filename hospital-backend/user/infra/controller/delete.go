package controller

import (
	"github.com/gin-gonic/gin"
	di_container "gitlab.com/hospital-web/di_container"
)

func DeleteUser(ctx *gin.Context) {

	var (
		err error
	)

	user_id := ctx.Query("user_id")

	user_service := di_container.UserService()

	err = user_service.DeleteUser(ctx, user_id)
	if err != nil {
		ctx.AbortWithStatus(500)
		return
	}

	ctx.AbortWithStatus(200)

}
