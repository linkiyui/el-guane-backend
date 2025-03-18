package controller

import (
	"github.com/gin-gonic/gin"
	di_container "gitlab.com/hospital-web/di_container"
)

func GetAllDoctors(ctx *gin.Context) {

	doctor_service := di_container.DoctorService()
	doctors, err := doctor_service.GetAllDoctors(ctx)
	if err != nil {
		ctx.AbortWithStatus(500)
		return
	}

	doctors_resp := make([]DoctorResponse, 0, len(doctors))

	for _, d := range doctors {
		doctor_resp := DoctorResponse{
			ID:   d.ID,
			Name: d.Name,
		}
		doctors_resp = append(doctors_resp, doctor_resp)
	}

	ctx.AbortWithStatusJSON(200, gin.H{
		"doctors": doctors_resp,
	})

}

type DoctorResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
