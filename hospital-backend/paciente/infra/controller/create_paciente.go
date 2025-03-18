package controller

import (
	"github.com/gin-gonic/gin"
	di_container "gitlab.com/hospital-web/di_container"
	"gitlab.com/hospital-web/hospital-backend/paciente/domain"
)

type CreatePacienteRequest struct {
	Name string `json:"name"`
	Age  int32  `json:"age"`
	Sex  string `json:"sex"`
}

func CreatePaciente(ctx *gin.Context) {

	var (
		err error
	)

	var req CreatePacienteRequest
	if err = ctx.ShouldBindJSON(&req); err != nil {
		ctx.AbortWithStatus(400)
		return
	}

	paciente_service := di_container.PacienteService()

	paciente := domain.Paciente{
		Name: req.Name,
		Age:  req.Age,
		Sex:  req.Sex,
	}

	err = paciente_service.CreatePaciente(ctx, paciente)
	if err != nil {
		ctx.AbortWithStatus(500)
		return
	}

	ctx.AbortWithStatus(200)

}
