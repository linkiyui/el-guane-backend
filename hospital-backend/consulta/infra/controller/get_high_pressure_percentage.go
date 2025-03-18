package controller

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/hospital-web/di_container"
)

func GetHighPressurePercentage(ctx *gin.Context) {

	var (
		totalConsultas        int64
		highPressureConsultas int64
		percentage            float64
		err                   error
	)

	consulta_service := di_container.ConsultaService()
	grave_service := di_container.GraveService()

	totalConsultas, err = consulta_service.GetTotalConsultas(ctx)
	if err != nil {
		ctx.AbortWithStatus(500)
		return
	}

	highPressureConsultas, err = grave_service.GetGravesHighPressure(ctx)
	if err != nil {
		ctx.AbortWithStatus(500)
		return
	}

	if totalConsultas == 0 {
		ctx.JSON(200, gin.H{"percentage": 0})
		return
	}

	percentage = (float64(highPressureConsultas) / float64(totalConsultas)) * 100
	ctx.JSON(200, gin.H{"percentage": percentage})

}
