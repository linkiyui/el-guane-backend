package domain

type User struct {
	ID       string `gorm:"primaryKey"`
	CI       string `gorm:"unique"`
	Name     string
	Username string `gorm:"unique"`
	Password string
	Role     string
}
