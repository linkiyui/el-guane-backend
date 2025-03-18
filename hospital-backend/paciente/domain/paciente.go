package domain

type Paciente struct {
	ID   string `gorm:"primaryKey"`
	Name string
	Age  int32
	Sex  string
}
