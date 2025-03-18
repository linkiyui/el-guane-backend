package controller

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/hospital-web/di_container"
)

func GetAllConsultas(ctx *gin.Context) {
	consulta_service := di_container.ConsultaService()

	consultas, err := consulta_service.GetAllConsultas(ctx)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, consultas)
}
