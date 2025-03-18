package controller

import (
	"time"

	"github.com/gin-gonic/gin"
	di_container "gitlab.com/hospital-web/di_container"
)

func GetLeve(ctx *gin.Context) {

	leve_service := di_container.LeveService()
	leves, err := leve_service.GetLeves(ctx)
	if err != nil {
		ctx.AbortWithStatus(500)
		return
	}

	consulta_service := di_container.ConsultaService()

	leves_resp := make([]LeveResponse, 0, len(leves))

	for _, l := range leves {
		leve_resp := LeveResponse{
			Consulta_Id: l.Consulta_Id,
			Diagnosis:   l.Diagnosis,
			Analisis:    l.Analisis,
		}
		consulta, err := consulta_service.FindByConsultaId(ctx, l.Consulta_Id)
		if err != nil {
			ctx.AbortWithStatus(500)
			return
		}
		leve_resp.Doctor_id = consulta.Doctor_id
		leve_resp.Date = consulta.Date
		leve_resp.Paciente_id = consulta.Paciente_id
		leves_resp = append(leves_resp, leve_resp)
	}

	ctx.AbortWithStatusJSON(200, gin.H{
		"leves": leves_resp,
	})

}

type LeveResponse struct {
	Consulta_Id string    `json:"consulta_id"`
	Doctor_id   string    `json:"doctor_id"`
	Date        time.Time `json:"date"`
	Paciente_id string    `json:"paciente_id"`
	Diagnosis   string    `json:"diagnosis"`
	Analisis    bool      `json:"analisis"`
}
