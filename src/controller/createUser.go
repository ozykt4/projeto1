package controller

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/ozykt4/projeto1/src/controller/model/request"
	"github.com/ozykt4/projeto1/src/validation"
)

func CreateUser(c *gin.Context) {
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		log.Printf("Error trying to marshal object, error=%s\n", err.Error())
		errRest := validation.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)
		return
	}
	fmt.Println(userRequest)
}
