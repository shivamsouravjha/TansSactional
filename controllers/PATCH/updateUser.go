package patch

import (
	"net/http"
	"transactional/constants"
	db "transactional/helpers/db/user"
	"transactional/struct/request"
	"transactional/struct/response"
	"transactional/utils"

	"github.com/gin-gonic/gin"
)

func UpdateUser(c *gin.Context) {
	updateUserEequest := request.UpdateUser{}
	userId := c.Request.Header.Get("userId")
	userID := request.UserID{
		UserID: userId,
	}
	if err := c.ShouldBind(&updateUserEequest); err != nil {
		c.JSON(422, utils.SendErrorResponse(err))
		return
	}
	resperr := response.Response{}
	err := db.UpdateUserDAO(c.Request.Context(), &updateUserEequest, &userID)
	if err != "" {
		resperr.Status = constants.API_FAILED_STATUS
		resperr.Message = err
		c.JSON(http.StatusInternalServerError, resperr)
		return
	}
	token, tokenerror := utils.GenerateToken()
	if tokenerror != nil {
		resperr.Status = constants.API_FAILED_STATUS
		resperr.Message = "User data updated,Please login"
		c.JSON(http.StatusInternalServerError, resperr)
		return
	}
	resp := response.CreatedResponse{}
	resp.Status.Status = "Success"
	resp.Status.Message = "User data updated successfully"
	resp.Token = token
	c.JSON(http.StatusOK, resp)
}
