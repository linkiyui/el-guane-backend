package controller

import (
	"github.com/gin-gonic/gin"
	di_container "gitlab.com/hospital-web/di_container"
)

func GetOldestLevePatientWithDiagnosis(ctx *gin.Context) {

	var (
		err       error
		diagnosis string
	)

	leve_service := di_container.LeveService()

	diagnosis = ctx.Query("diagnosis")

	if diagnosis == "" {
		ctx.JSON(400, gin.H{"error": "Diagnosis query parameter is required"})
		return
	}

	paciente, err := leve_service.GetOldestLevePatientWithDiagnosis(ctx, diagnosis)
	if err != nil {
		ctx.AbortWithStatus(500)
		return
	}
	resp := resp{
		Name: paciente.Name,
		Age:  paciente.Age,
		Sex:  paciente.Sex,
	}

	ctx.JSON(200, gin.H{
		"paciente": resp,
	})

}

type resp struct {
	Name string `json:"name"`
	Age  int32  `json:"age"`
	Sex  string `json:"sex"`
}
