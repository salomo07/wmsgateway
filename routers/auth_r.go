package routers

import (
	"wmsgateway/controllers"
	"github.com/gin-gonic/gin"
)

func GatewayRouter(r *gin.Engine) {
	// r.Static("assets/", "./assets")

	master := r.Group("/auth")
	{
		master.POST("/login", func(c *gin.Context) {
			c.Header("Content-Type", "application/json; charset=utf-8")
			controllers.SaveRedis()
		})
	}
}
