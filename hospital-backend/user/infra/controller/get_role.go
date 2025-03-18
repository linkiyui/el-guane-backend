package controller

import (
	"github.com/gin-gonic/gin"
	di_container "gitlab.com/hospital-web/di_container"
)

func GetRole(ctx *gin.Context) {

	user_service := di_container.UserService()
	rol := user_service.GetRole()

	rol_resp_arr := make([]RolResp, 0, len(rol))

	for _, r := range rol {
		rol_resp := RolResp{
			Rol: string(r),
		}
		rol_resp_arr = append(rol_resp_arr, rol_resp)
	}

	ctx.AbortWithStatusJSON(200, gin.H{
		"roles": rol_resp_arr,
	})

}

type RolResp struct {
	Rol string
}
