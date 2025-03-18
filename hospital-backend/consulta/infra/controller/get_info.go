package controller

import (
	"time"

	"github.com/gin-gonic/gin"
	"gitlab.com/hospital-web/di_container"
	"gitlab.com/hospital-web/hospital-backend/consulta/domain"
)

func GetConsultaInfo(ctx *gin.Context) {
	consulta_service := di_container.ConsultaService()

	consulta, err := consulta_service.GetConsultaInfo(ctx)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	consulta_resp := ConsultaResponse{}

	if consulta.Grade == domain.GradeHigh {
		grave_service := di_container.GraveService()
		grave, err := grave_service.GetGraveInfo(ctx, consulta.ID)
		if err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}
		consulta_resp = ConsultaResponse{
			Consulta_Id:   grave.Consulta_Id,
			Doctor_id:     grave.Doctor_Id,
			Date:          grave.Date,
			Paciente_id:   grave.Paciente_Id,
			Sintomas:      grave.Sintomas,
			Ingreso:       grave.Ingreso,
			Temp:          grave.Temp,
			Pulse:         grave.Pulse,
			Press_Min:     grave.Press_Min,
			Press_Max:     grave.Press_Max,
			Doctor_Name:   grave.Doctor_Name,
			Paciente_Name: grave.Paciente_Name,
			Paciente_Age:  grave.Paciente_Age,
			Paciente_Sex:  grave.Paciente_Sex,
		}

	} else {
		leve_service := di_container.LeveService()
		leve, err := leve_service.GetLeveInfo(ctx, consulta.ID)
		if err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}
		consulta_resp = ConsultaResponse{
			Consulta_Id:   leve.Consulta_Id,
			Doctor_id:     leve.Doctor_Id,
			Date:          leve.Date,
			Paciente_id:   leve.Paciente_Id,
			Diagnosis:     leve.Diagnosis,
			Analisis:      leve.Analisis,
			Doctor_Name:   leve.Doctor_Name,
			Paciente_Name: leve.Paciente_Name,
			Paciente_Age:  leve.Paciente_Age,
			Paciente_Sex:  leve.Paciente_Sex,
		}

	}

	ctx.JSON(200, gin.H{"consulta": consulta_resp})
}

type ConsultaResponse struct {
	Consulta_Id   string    `json:"consulta_id"`
	Doctor_id     string    `json:"doctor_id"`
	Doctor_Name   string    `json:"doctor_name"`
	Date          time.Time `json:"date"`
	Paciente_id   string    `json:"paciente_id"`
	Paciente_Name string    `json:"paciente_id"`
	Paciente_Age  int32     `json:"paciente_age"`
	Paciente_Sex  string    `json:"paciente_sex"`
	Grade         string    `json:"grade"`
	Diagnosis     string    `json:"diagnosis,omitempty"`
	Analisis      bool      `json:"analisis,omitempty"`
	Sintomas      string    `json:"sintomas,omitempty"`
	Ingreso       bool      `json:"ingreso,omitempty"`
	Temp          float32   `json:"temp,omitempty"`
	Pulse         float32   `json:"pulse,omitempty"`
	Press_Min     int32     `json:"press_min,omitempty"`
	Press_Max     int32     `json:"press_max,omitempty"`
}
