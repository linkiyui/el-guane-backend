package domain

type Leve struct {
	Consulta_Id string `gorm:"primaryKey"`
	Diagnosis   string
	Analisis    bool
}
