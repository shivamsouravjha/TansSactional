package patch

import (
	"net/http"
	"transactional/constants"
	"transactional/helpers/db"
	"transactional/struct/request"
	"transactional/struct/response"
	"transactional/utils"

	"github.com/gin-gonic/gin"
)

func UpdateCompany(c *gin.Context) {

	companyID := c.Request.Header.Get("companyId")
	updateCompanyRequest := request.CompanyID{
		CompanyID: companyID,
	}
	updateUserEequest := request.CreateCompany{}

	if err := c.ShouldBind(&updateUserEequest); err != nil {
		c.JSON(422, utils.SendErrorResponse(err))
		return
	}
	resp := response.CreatedResponse{}
	err := db.UpdateCompanyDAO(c.Request.Context(), &updateUserEequest, &updateCompanyRequest)
	if err != "" {
		resp.Status.Message = constants.API_FAILED_STATUS
		resp.Status.Status = err
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	token, tokenerror := utils.GenerateToken()
	if tokenerror != nil {
		resp.Status.Status = constants.API_FAILED_STATUS
		resp.Status.Message = "User Created,Please login"
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	resp.Status.Status = "Success"
	resp.Status.Message = "User Created successfully"
	resp.Token = token
	c.JSON(http.StatusOK, resp)
}
