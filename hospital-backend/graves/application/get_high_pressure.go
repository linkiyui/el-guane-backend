package application

import (
	"context"
)

func (u *GraveService) GetGravesHighPressure(ctx context.Context) (int64, error) {
	return u.grave_repo.GetGravesHighPressure(ctx)
}
