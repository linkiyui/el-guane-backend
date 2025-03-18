package application

import (
	"context"
)

func (u *LeveService) GetTotalLeves(ctx context.Context) (int64, error) {
	return u.leve_repo.GetTotalLeves(ctx)
}

func (u *LeveService) GetAnalysisLeves(ctx context.Context) (int64, error) {
	return u.leve_repo.GetAnalysisLeves(ctx)
}
