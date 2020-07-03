package users

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nurzamanindra/bookstore_users-api/domain/users"
	"github.com/nurzamanindra/bookstore_users-api/services"
	"github.com/nurzamanindra/bookstore_users-api/utils/errors"
)

func CreateUser(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		//TODO : Handle error bad request
		restErr := errors.NewBadRequestError("Invalid Json Body")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestError("Invalid user id or user id should be a number")
		c.JSON(err.Status, err)
		return
	}
	user, getUserErr := services.GetUser(userId)
	if getUserErr != nil {
		c.JSON(getUserErr.Status, getUserErr)
		return
	}
	c.JSON(http.StatusOK, user)
}

func UpdateUser(c *gin.Context) {
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestError("Invalid user id or user id should be a number")
		c.JSON(err.Status, err)
		return
	}

	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid Json request")
		c.JSON(restErr.Status, restErr)
		return
	}
	user.Id = userId
	result, err := services.UpdateUser(user)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, result)
}
