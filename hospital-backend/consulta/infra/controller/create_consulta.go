package controller

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	di_container "gitlab.com/hospital-web/di_container"
	"gitlab.com/hospital-web/hospital-backend/consulta/domain"
	grave_domain "gitlab.com/hospital-web/hospital-backend/graves/domain"
	leve_domain "gitlab.com/hospital-web/hospital-backend/leves/domain"
)

type CreateConsultaRequest struct {
	Doctor_id   string    `json:"doctor_id"`
	Date        time.Time `json:"date"`
	Paciente_id string    `json:"paciente_id"`
	Grade       string    `json:"grade"`
	Diagnosis   string    `json:"diagnosis,omitempty"`
	Analisis    bool      `json:"analisis,omitempty"`
	Sintomas    string    `json:"sintomas,omitempty"`
	Ingreso     bool      `json:"ingreso,omitempty"`
	Temp        float32   `json:"temp,omitempty"`
	Pulse       float32   `json:"pulse,omitempty"`
	Press_Min   int32     `json:"press_min,omitempty"`
	Press_Max   int32     `json:"press_max,omitempty"`
}

func CreateConsulta(ctx *gin.Context) {

	var (
		err error
	)

	var req CreateConsultaRequest
	if err = ctx.ShouldBindJSON(&req); err != nil {
		ctx.AbortWithStatus(400)
		return
	}

	consulta_service := di_container.ConsultaService()

	consulta := domain.Consulta{
		Doctor_id:   req.Doctor_id,
		Date:        req.Date,
		Paciente_id: req.Paciente_id,
		Grade:       domain.Grade(req.Grade),
	}
	// fmt.Println(consulta)

	id, err := consulta_service.CreateConsulta(ctx, consulta)
	if err != nil {
		ctx.AbortWithStatus(500)
		return
	}

	if req.Grade == string(domain.GradeHigh) {
		grave := grave_domain.Grave{
			Consulta_Id: id,
			Sintomas:    req.Sintomas,
			Ingreso:     req.Ingreso,
			Temp:        req.Temp,
			Pulse:       req.Pulse,
			Press_Min:   req.Press_Min,
			Press_Max:   req.Press_Max,
		}

		grave_service := di_container.GraveService()
		err = grave_service.CreateGrave(ctx, grave)
		if err != nil {
			ctx.AbortWithStatus(500)
			return
		}

	}

	if req.Grade == string(domain.GradeLow) {

		leve := leve_domain.Leve{
			Consulta_Id: id,
			Diagnosis:   req.Diagnosis,
			Analisis:    req.Analisis,
		}

		fmt.Print(leve.Consulta_Id)

		fmt.Println(leve)
		leve_service := di_container.LeveService()
		err = leve_service.CreateLeve(ctx, leve)
		if err != nil {
			ctx.AbortWithStatus(500)
			return

		}
	}

	ctx.AbortWithStatus(200)

}
