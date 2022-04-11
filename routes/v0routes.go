package routes

import (
	get "transactional/controllers/GET"
	"transactional/middlewares"

	"github.com/gin-gonic/gin"
)

func v0Routes(route *gin.RouterGroup) {
	v1Routes := route.Group("/v0")
	{
		company := v1Routes.Group("/company")
		{
			//			company.POST("/createcompany")
			//	company.GET("/getcompany")
			company.Use(middlewares.JWT("companytoken"))
			//			company.PATCH("/updatecompany")
		}
		user := v1Routes.Group("/user")
		{
			user.GET("/getuser", get.GetUser)
			//			user.POST("/createuser")
			user.Use(middlewares.JWT("usertoken"))
			//			user.PATCH("/updateuser")
		}
		invoices := v1Routes.Group("/invoices")
		{
			invoices.Use(middlewares.JWT("usertoken"))
			//	invoices.GET("/getinovice")
			//	invoices.POST("/createinvoice")
			//	invoices.PATCH("/updateinvoices")
		}
	}
}
