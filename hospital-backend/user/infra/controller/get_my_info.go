package controller

import (
	"fmt"

	gin "github.com/gin-gonic/gin"
	"gitlab.com/hospital-web/di_container"
)

func GetMyInfo(ctx *gin.Context) {

	uid, exists := ctx.Get("user_id")

	if !exists {
		ctx.JSON(401, gin.H{"error": "No user found"})
		return
	}

	user_id := fmt.Sprint(uid)

	userService := di_container.UserService()

	user, err := userService.GetMyInfo(ctx, user_id)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"user": user})
}
