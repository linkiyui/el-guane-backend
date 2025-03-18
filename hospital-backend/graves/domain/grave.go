package domain

type Grave struct {
	Consulta_Id string `gorm:"primaryKey"`
	Sintomas    string
	Ingreso     bool
	Temp        float32
	Pulse       float32
	Press_Min   int32
	Press_Max   int32
}
