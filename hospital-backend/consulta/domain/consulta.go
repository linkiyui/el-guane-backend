package domain

import "time"

type Consulta struct {
	ID          string `gorm:"primaryKey"`
	Doctor_id   string
	Date        time.Time
	Paciente_id string
	Grade       Grade
}

type Grade string

var (
	GradeLow  Grade = "L"
	GradeHigh Grade = "H"
)
