package route

import (
	"PlusOne/controller"
	"github.com/gin-gonic/gin"
)

func TestRoutes(api *gin.RouterGroup) {
	api.GET("/test", controller.Test)

}
