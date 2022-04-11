package get

import (
	"net/http"
	"transactional/constants"
	"transactional/helpers/db"
	"transactional/struct/request"
	"transactional/struct/response"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	userId := c.Request.Header.Get("userId")
	getContentRequest := request.UserID{
		UserID: userId,
	}
	resp := response.GetUserReponse{}
	UserDetails, err := db.GetUserDAO(c.Request.Context(), &getContentRequest)
	if err != nil {
		resp.Message.Status = constants.API_FAILED_STATUS
		resp.Message.Message = "Unlocking Content Failed"
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	resp.Message.Message = "Success"
	resp.Message.Status = "Contents unlocked successfully"
	resp.UserDetails = UserDetails
	c.JSON(http.StatusOK, resp)
}