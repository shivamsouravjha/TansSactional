package get

import (
	"net/http"
	"transactional/constants"
	"transactional/helpers/db"
	"transactional/struct/request"
	"transactional/struct/response"

	"github.com/gin-gonic/gin"
)

func GetCompany(c *gin.Context) {
	companyID := c.Request.Header.Get("companyId")
	getContentRequest := request.CompanyID{
		CompanyID: companyID,
	}
	resp := response.GetCompanyResponse{}
	CompanyDetails, err := db.GetCompanyDAO(c.Request.Context(), &getContentRequest)
	if err != nil {
		resp.Message.Status = constants.API_FAILED_STATUS
		resp.Message.Message = "Unlocking Content Failed"
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	resp.Message.Message = "Success"
	resp.Message.Status = "Contents unlocked successfully"
	resp.CompanyDetails = CompanyDetails
	c.JSON(http.StatusOK, resp)
}
