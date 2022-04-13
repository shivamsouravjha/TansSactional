package get

import (
	"net/http"
	"transactional/constants"
	db "transactional/helpers/db/company"
	"transactional/struct/request"
	"transactional/struct/response"

	"github.com/gin-gonic/gin"
)

func GetCompany(c *gin.Context) {
	companyID := c.Request.Header.Get("companyId")
	getContentRequest := request.CompanyID{
		CompanyID: companyID,
	}
	if len(companyID) > 0 {
		resp := response.GetCompanyResponse{}
		CompanyDetails, err := db.GetCompanyDAO(c.Request.Context(), &getContentRequest)
		if err != nil {
			resp.Message.Status = constants.API_FAILED_STATUS
			resp.Message.Message = "No such user exists"
			c.JSON(http.StatusInternalServerError, resp)
			return
		}
		resp.Message.Status = "Success"
		resp.Message.Message = "Company detail fetched successfully"
		resp.CompanyDetails = CompanyDetails
		c.JSON(http.StatusOK, resp)
	} else {
		resp := response.GetAllCompanyResponse{}
		CompanyDetails, err := db.GetAllCompanyDAO(c.Request.Context())
		if err != nil {
			resp.Message.Status = constants.API_FAILED_STATUS
			resp.Message.Message = "Fetching detail of all companies Failed"
			c.JSON(http.StatusInternalServerError, resp)
			return
		}
		resp.Message.Status = "Success"
		resp.Message.Message = "Detail of all companies fetched successfully"
		resp.CompanyDetails = CompanyDetails
		c.JSON(http.StatusOK, resp)
	}
}
