package controller

import (
	"time"

	"github.com/gin-gonic/gin"
	"gitlab.com/hospital-web/di_container"
)

func GetLeveConsultasByDiagnosis(ctx *gin.Context) {

	var (
		err error
	)

	diagnosis := ctx.Query("diagnosis")

	if diagnosis == "" {
		ctx.JSON(400, gin.H{"error": "Diagnosis query parameter is required"})
		return
	}

	leve_service := di_container.LeveService()

	leves, err := leve_service.GetLevesByDiagnosis(ctx, diagnosis)
	if err != nil {
		ctx.AbortWithStatus(500)
		return
	}

	leves_resp := make([]leveResponse, 0, len(leves))

	for _, l := range leves {
		leve_resp := leveResponse{
			Consulta_Id:   l.Consulta_Id,
			Diagnosis:     l.Diagnosis,
			Analisis:      l.Analisis,
			Doctor_name:   l.Doctor_name,
			Date:          l.Date,
			Paciente_name: l.Paciente_name,
			Paciente_age:  l.Paciente_age,
			Paciente_sex:  l.Paciente_sex,
		}
		leves_resp = append(leves_resp, leve_resp)
	}

	ctx.AbortWithStatusJSON(200, gin.H{
		"leves": leves_resp,
	})

}

type leveResponse struct {
	Consulta_Id   string    `json:"consulta_id"`
	Doctor_name   string    `json:"doctor_name"`
	Date          time.Time `json:"date"`
	Paciente_name string    `json:"paciente_name"`
	Paciente_age  int32     `json:"paciente_age"`
	Paciente_sex  string    `json:"paciente_sex"`
	Diagnosis     string    `json:"diagnosis"`
	Analisis      bool      `json:"analisis"`
}
