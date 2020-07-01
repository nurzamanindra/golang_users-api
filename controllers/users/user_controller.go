package users

import (
	"net/http"

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
		//TODO : Handle user creator error
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Create user function is not implemented! Implement me!")
}
