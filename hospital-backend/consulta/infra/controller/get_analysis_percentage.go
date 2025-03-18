package controller

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/hospital-web/di_container"
)

func GetAnalysisPercentage(ctx *gin.Context) {

	var (
		totalLeves    int64
		analysisLeves int64
		percentage    float64
		err           error
	)

	leve_service := di_container.LeveService()

	totalLeves, err = leve_service.GetTotalLeves(ctx)
	if err != nil {
		ctx.AbortWithStatus(500)
		return
	}

	analysisLeves, err = leve_service.GetAnalysisLeves(ctx)
	if err != nil {
		ctx.AbortWithStatus(500)
		return
	}

	if totalLeves == 0 {
		ctx.JSON(200, gin.H{"percentage": 0})
		return
	}

	percentage = (float64(analysisLeves) / float64(totalLeves)) * 100
	ctx.JSON(200, gin.H{"percentage": percentage})

}
