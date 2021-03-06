package post

import (
	"net/http"
	"transactional/constants"
	db "transactional/helpers/db/user"
	"transactional/struct/request"
	"transactional/struct/response"
	"transactional/utils"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	createUserStruct := request.CreateUser{}
	if err := c.ShouldBind(&createUserStruct); err != nil {
		c.JSON(422, utils.SendErrorResponse(err))
		return
	}
	resp := response.CreatedResponse{}
	err := db.CreateUserDAO(c.Request.Context(), &createUserStruct)
	if err != "" {
		resp.Status.Status = constants.API_FAILED_STATUS
		resp.Status.Message = err
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

func CreateProduct(c *gin.Context) {
	userId := c.Request.Header.Get("userId")
	createUserStruct := request.CreateProduct{}
	if err := c.ShouldBind(&createUserStruct); err != nil {
		c.JSON(422, utils.SendErrorResponse(err))
		return
	}
	resp := response.Response{}
	err := db.CreateProductDAO(c.Request.Context(), &createUserStruct, userId)
	if err != "" {
		resp.Message = constants.API_FAILED_STATUS
		resp.Status = err
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	c.JSON(http.StatusOK, resp)
}
