package route

import (
	"github.com/gin-gonic/gin"
)

func Route(route *gin.Engine) {
	// 使用路由组
	v1 := route.Group("/v1")
	{
		api := v1.Group("/api")
		TestRoutes(api)
	}
	//route.SetupAPIRoutes(api)

	//admin := route.Group("/admin")
	//routes.SetupAdminRoutes(admin)
}
