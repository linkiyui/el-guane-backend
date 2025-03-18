package controller

import (
	"time"

	"github.com/gin-gonic/gin"
	di_container "gitlab.com/hospital-web/di_container"
)

func GetGrave(ctx *gin.Context) {
	grave_service := di_container.GraveService()
	graves, err := grave_service.GetGraves(ctx)
	if err != nil {
		ctx.AbortWithStatus(500)
		return
	}

	consulta_service := di_container.ConsultaService()

	graves_resp := make([]GraveResponse, 0, len(graves))

	for _, g := range graves {
		grave_resp := GraveResponse{
			Consulta_Id: g.Consulta_Id,
			Sintomas:    g.Sintomas,
			Ingreso:     g.Ingreso,
			Temp:        g.Temp,
			Pulse:       g.Pulse,
			Press_Min:   g.Press_Min,
			Press_Max:   g.Press_Max,
		}
		consulta, err := consulta_service.FindByConsultaId(ctx, g.Consulta_Id)
		if err != nil {
			ctx.AbortWithStatus(500)
			return
		}
		grave_resp.Doctor_id = consulta.Doctor_id
		grave_resp.Date = consulta.Date
		grave_resp.Paciente_id = consulta.Paciente_id
		graves_resp = append(graves_resp, grave_resp)
	}

	ctx.AbortWithStatusJSON(200, gin.H{
		"graves": graves_resp,
	})
}

type GraveResponse struct {
	Consulta_Id string    `json:"consulta_id"`
	Doctor_id   string    `json:"doctor_id"`
	Date        time.Time `json:"date"`
	Paciente_id string    `json:"paciente_id"`
	Sintomas    string    `json:"sintomas"`
	Ingreso     bool      `json:"ingreso"`
	Temp        float32   `json:"temp"`
	Pulse       float32   `json:"pulse"`
	Press_Min   int32     `json:"press_min"`
	Press_Max   int32     `json:"press_max"`
}
