package domain

import "time"

type Leve_diagnosis struct {
	Consulta_Id   string    `gorm:"primaryKey"`
	Doctor_name   string    `gorm:"primaryKey"`
	Date          time.Time `gorm:"primaryKey"`
	Paciente_name string    `gorm:"primaryKey"`
	Paciente_age  int32     `gorm:"primaryKey"`
	Paciente_sex  string    `gorm:"primaryKey"`
	Diagnosis     string    `gorm:"primaryKey"`
	Analisis      bool      `gorm:"primaryKey"`
}
