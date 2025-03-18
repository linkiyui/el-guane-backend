package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// func loadEnvVariables() (string, string, string, string) {
// 	user := os.Getenv("DB_USER")
// 	password := os.Getenv("DB_PASSWORD")
// 	host := os.Getenv("DB_HOST")
// 	port := os.Getenv("DB_PORT")

// 	return user, password, host, port
// }

var Database = func() *gorm.DB {
	dsn := "host=127.0.0.1 user=postgres password=sasa dbname=hospital_backend port=5432 sslmode=disable TimeZone=America/New_York"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
	}
	return db
}()
