package domain

import "time"

type LeveToDomain struct {
	Consulta_Id   string `gorm:"primaryKey"`
	Paciente_Id   string
	Doctor_Id     string
	Doctor_Name   string
	Date          time.Time
	Paciente_Name string
	Paciente_Age  int32
	Paciente_Sex  string
	Diagnosis     string
	Analisis      bool
}
