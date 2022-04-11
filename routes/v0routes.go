package routes

import (
	"transactional/middlewares"

	"github.com/gin-gonic/gin"
)

func v0Routes(route *gin.RouterGroup) {
	v1Routes := route.Group("/v0")
	{
		company := v1Routes.Group("/company")
		{
			company.POST("/createcompany")
			company.GET("/getcompany")
			v1Routes.Use(middlewares.JWT("companytoken"))
			company.PATCH("/updatecompany")
		}
		user := v1Routes.Group("/user")
		{
			user.GET("/getuser")
			user.POST("/createuser")
			v1Routes.Use(middlewares.JWT("usertoken"))
			user.PATCH("/updateuser")
		}
		invoices := v1Routes.Group("/invoices")
		{
			v1Routes.Use(middlewares.JWT("usertoken"))
			invoices.GET("/getinovice")
			invoices.POST("/createinvoice")
			invoices.PATCH("/updateinvoices")
		}
	}
}
