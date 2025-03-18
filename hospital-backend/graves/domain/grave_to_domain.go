package domain

import "time"

type GraveToDomain struct {
	Consulta_Id   string `gorm:"primaryKey"`
	Paciente_Id   string
	Doctor_Id     string
	Doctor_Name   string
	Date          time.Time
	Grade         string
	Paciente_Name string
	Paciente_Age  int32
	Paciente_Sex  string
	Sintomas      string
	Ingreso       bool
	Temp          float32
	Pulse         float32
	Press_Min     int32
	Press_Max     int32
}
