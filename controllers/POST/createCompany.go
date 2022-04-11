package post

import (
	"net/http"
	"transactional/constants"
	"transactional/helpers/db"
	"transactional/struct/request"
	"transactional/struct/response"
	"transactional/utils"

	"github.com/gin-gonic/gin"
)

func CreateCompany(c *gin.Context) {
	createUserStruct := request.CreateCompany{}
	if err := c.ShouldBind(&createUserStruct); err != nil {
		c.JSON(422, utils.SendErrorResponse(err))
		return
	}
	resp := response.CreatedResponse{}
	err := db.CreateCompanyDAO(c.Request.Context(), &createUserStruct)
	if err != "" {
		resp.Status.Message = constants.API_FAILED_STATUS
		resp.Status.Status = err
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	token, tokenerror := utils.GenerateToken("company")
	if tokenerror != nil {
		resp.Status.Status = constants.API_FAILED_STATUS
		resp.Status.Message = "Company Created,Please login"
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	resp.Status.Status = "Success"
	resp.Status.Message = "Company Created successfully"
	resp.Token = token
	c.JSON(http.StatusOK, resp)
}
