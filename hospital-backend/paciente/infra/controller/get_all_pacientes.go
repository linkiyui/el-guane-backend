package controller

import (
	"github.com/gin-gonic/gin"
	di_container "gitlab.com/hospital-web/di_container"
)

func GetAllPacientes(ctx *gin.Context) {

	paciente_service := di_container.PacienteService()
	pacientes, err := paciente_service.GetAllPacientes(ctx)
	if err != nil {
		ctx.AbortWithStatus(500)
		return
	}

	pacientes_resp := make([]PacienteResponse, 0, len(pacientes))

	for _, p := range pacientes {
		paciente_resp := PacienteResponse{
			ID:   p.ID,
			Name: p.Name,
			Age:  p.Age,
			Sex:  p.Sex,
		}
		pacientes_resp = append(pacientes_resp, paciente_resp)
	}

	ctx.AbortWithStatusJSON(200, gin.H{
		"pacientes": pacientes_resp,
	})

}

type PacienteResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  int32  `json:"age"`
	Sex  string `json:"sex"`
}
