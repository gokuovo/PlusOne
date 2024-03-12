package controller

import (
	"PlusOne/response"
	"PlusOne/service"
	"github.com/gin-gonic/gin"
)

func Test(c *gin.Context) {
	serv := service.TestService{}
	ret := serv.Test()
	c.JSON(200, response.Success(ret))
}
