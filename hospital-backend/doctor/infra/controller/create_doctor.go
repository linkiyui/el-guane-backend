package controller

import (
	"github.com/gin-gonic/gin"
	di_container "gitlab.com/hospital-web/di_container"
	"gitlab.com/hospital-web/hospital-backend/doctor/domain"
)

type CreateDoctorRequest struct {
	Name string `json:"name"`
}

func CreateDoctor(ctx *gin.Context) {

	var (
		err error
	)

	var req CreateDoctorRequest
	if err = ctx.ShouldBindJSON(&req); err != nil {
		ctx.AbortWithStatus(400)
		return
	}

	doctor_service := di_container.DoctorService()

	doctor := domain.Doctor{
		Name: req.Name,
	}

	err = doctor_service.CreateDoctor(ctx, doctor)
	if err != nil {
		ctx.AbortWithStatus(500)
		return
	}

	ctx.AbortWithStatus(200)

}
