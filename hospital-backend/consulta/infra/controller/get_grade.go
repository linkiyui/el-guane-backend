package controller

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/hospital-web/di_container"
)

func GetConsultaGrade(ctx *gin.Context) {
	consulta_service := di_container.ConsultaService()

	grade, err := consulta_service.Getgrade(ctx)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	grades := make([]string, 0, len(grade))

	for _, g := range grade {
		grades = append(grades, string(g))
	}

	ctx.JSON(200, gin.H{"grades": grades})
}
