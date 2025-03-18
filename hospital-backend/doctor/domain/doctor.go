package domain

type Doctor struct {
	ID   string `gorm:"primaryKey"`
	Name string
}
