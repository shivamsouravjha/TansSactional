package routes

import (
	"github.com/gin-gonic/gin"
)

func v0Routes(route *gin.RouterGroup) {
	v1Routes := route.Group("/v0")
	{
		v1Routes.GET("/getUser/:penName")
	}
}
