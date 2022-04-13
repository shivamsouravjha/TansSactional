package get

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"transactional/constants"
	db "transactional/helpers/db/invoices"
	"transactional/struct/request"
	"transactional/struct/response"

	"github.com/gin-gonic/gin"
)

func GetInvoice(c *gin.Context) {
	userID := c.Request.Header.Get("userId")
	filter := c.Request.Header.Get("filter")
	sort := c.Request.Header.Get("sort")
	page := c.Request.Header.Get("page")
	Page, _ := strconv.Atoi(page)
	getContentRequest := request.GetInvoice{
		UserID: userID,
		Sort:   sort,
		Page:   Page,
	}
	var filterUnmarshal map[string]string
	error := json.Unmarshal([]byte(filter), &filterUnmarshal)
	if error == nil {

		getContentRequest.Filter = request.Filter{
			Comperator:   filterUnmarshal["comperator"],
			Value:        filterUnmarshal["value"],
			Position:     filterUnmarshal["position"],
			Settled:      filterUnmarshal["settled"],
			Acknowledged: filterUnmarshal["acknowledged"],
		}
	}
	resp := response.InvoiceResponse{}
	if len(userID) == 0 {
		resp.Message.Message = "No UserID sent"
		resp.Message.Status = constants.API_FAILED_STATUS
		c.JSON(http.StatusOK, resp)
	} else {
		InvoiceDetails, err := db.GetInvoicesDAO(c.Request.Context(), &getContentRequest)
		fmt.Println(err)
		if err != "" {
			resp.Message.Status = constants.API_FAILED_STATUS
			resp.Message.Message = "User has no invoice"
			c.JSON(http.StatusInternalServerError, resp)
			return
		}
		resp.Message.Status = "Success"
		resp.Message.Message = "Invoice Fetched successfully"
		resp.InvoiceResponse = InvoiceDetails
		resp.Page = Page + 1
		c.JSON(http.StatusOK, resp)
	}
}
