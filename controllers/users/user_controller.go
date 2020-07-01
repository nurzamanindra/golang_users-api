package users

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nurzamanindra/bookstore_users-api/domain/users"
)

func CreateUser(c *gin.Context) {
	var user users.User

	bytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		//TODO: error handler
		return
	}
	if err := json.Unmarshal(bytes, &user); err != nil {
		//TODO: Handle Json Error
		fmt.Println(err)
		return
	}
	fmt.Println(user)

	// fmt.Println(string(bytes))

	c.String(http.StatusNotImplemented, "Create user function is not implemented! Implement me!")
}

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Create user function is not implemented! Implement me!")
}
