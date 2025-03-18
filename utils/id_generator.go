package utils

import "github.com/google/uuid"

func GenerateUUIDv7() (string, error) {
	res, err := uuid.NewV7()
	if err != nil {
		return "", err
	}
	return res.String(), nil
}
