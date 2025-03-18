package middlewares

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"gitlab.com/hospital-web/di_container"
	"gitlab.com/hospital-web/errors"
	"gitlab.com/hospital-web/hospital-backend/user/domain"
)

func VerifyAdmin(c *gin.Context) {

	uid, exists := c.Get("user_id")
	if !exists {
		c.AbortWithStatus(401)
		return
	}

	user_id := fmt.Sprint(uid)

	user_service := di_container.UserService()

	user, err := user_service.FindByUsername(c, user_id)
	if err != nil {
		de := errors.IsDomainError(err)
		if de != nil {
			c.AbortWithStatusJSON(de.Code, gin.H{"error": de.Message})
			return
		}
		c.AbortWithStatus(500)
		return
	}

	role := user.Role

	if role != string(domain.RolAdmin) {
		c.AbortWithStatus(401)
		return
	}
	c.Next()
}
