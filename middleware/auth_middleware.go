package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/thoriqaufar/gin-jwt-impl/errorhandler"
	"github.com/thoriqaufar/gin-jwt-impl/helper"
)

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			errorhandler.HandleError(c, &errorhandler.UnauthorizedError{
				Message: "Unauthorized",
			})
			c.Abort()
			return
		}

		userId, err := helper.ValidateToken(tokenString)
		if err != nil {
			errorhandler.HandleError(c, &errorhandler.UnauthorizedError{
				Message: err.Error(),
			})
			c.Abort()
			return
		}

		c.Set("userId", *userId)
		c.Next()
	}
}
