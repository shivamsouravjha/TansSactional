package post

import (
	"net/http"
	"transactional/constants"
	db "transactional/helpers/db/invoices"
	"transactional/struct/request"
	"transactional/struct/response"

	"github.com/gin-gonic/gin"
)

func CreateInvoice(c *gin.Context) {
	createInvoiceStruct := request.CreateInvoice{}
	if err := c.ShouldBind(&createInvoiceStruct); err != nil {
		resp := response.Response{}
		resp.Status = "failed"
		resp.Message = "Validation failed"
		c.JSON(422, resp)
		return
	}
	resp := response.Response{}
	err := db.CreateInvoiceDAO(c.Request.Context(), &createInvoiceStruct)
	if err != "" {
		resp.Message = constants.API_FAILED_STATUS
		resp.Status = err
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	resp.Status = "Success"
	resp.Message = "Invoice Created successfully"
	c.JSON(http.StatusOK, resp)
}
