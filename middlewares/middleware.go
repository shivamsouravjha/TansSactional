package middlewares

import (
	"transactional/constants"
	"transactional/utils"

	"github.com/gin-gonic/gin"
)

func JWT(header string) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader(header)
		if token != "" {
			jwtCheck, err := utils.TokenParse(token, "mySigningKey")
			if err != nil {
				c.AbortWithStatusJSON(401, constants.INVALID_TOKEN_RESPONSE)
				return
			} else if jwtCheck.Valid {
				c.Next()
			}
		} else {
			c.AbortWithStatusJSON(401, constants.INVALID_TOKEN_RESPONSE)
		}
	}
}
