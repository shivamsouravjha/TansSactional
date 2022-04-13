package patch

import (
	"net/http"
	"transactional/constants"
	db "transactional/helpers/db/invoices"
	"transactional/struct/request"
	"transactional/struct/response"
	"transactional/utils"

	"github.com/gin-gonic/gin"
)

func UpdateInvoice(c *gin.Context) {
	invoicesID := c.Request.Header.Get("idinvoices")
	updateCompanyRequest := request.UpdateInvoce{
		InvoicesID: invoicesID,
	}
	if err := c.ShouldBind(&updateCompanyRequest); err != nil {
		c.JSON(422, utils.SendErrorResponse(err))
		return
	}
	resp := response.Response{}
	err := db.UpdateInvociesDAO(c.Request.Context(), &updateCompanyRequest)
	if err != "" {
		resp.Message = constants.API_FAILED_STATUS
		resp.Status = err
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	resp.Status = "Success"
	resp.Message = "Invoice updated successfully"
	c.JSON(http.StatusOK, resp)
}

func SendInvoice(c *gin.Context) {
	invoicesID := c.Request.Header.Get("idinvoices")
	updateCompanyRequest := request.UpdateInvoce{
		InvoicesID: invoicesID,
	}
	resp := response.Response{}
	err := db.SettleInvoiceDAO(c.Request.Context(), &updateCompanyRequest)
	if err != "" {
		resp.Message = constants.API_FAILED_STATUS
		resp.Status = err
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	resp.Status = "Success"
	resp.Message = "Invoice Sent successfully"
	c.JSON(http.StatusOK, resp)
}

func AcknowledgeInvoice(c *gin.Context) {
	invoicesID := c.Request.Header.Get("idinvoices")
	updateCompanyRequest := request.UpdateInvoce{
		InvoicesID: invoicesID,
	}
	resp := response.Response{}
	err := db.AcknowledgeInvoiceDAO(c.Request.Context(), &updateCompanyRequest)
	if err != "" {
		resp.Message = constants.API_FAILED_STATUS
		resp.Status = err
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	resp.Status = "Success"
	resp.Message = "Invoice Acknowledged successfully"
	c.JSON(http.StatusOK, resp)
}
