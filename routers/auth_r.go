package routers

import (
	"wmsgateway/controllers"
	"github.com/gin-gonic/gin"
)

func GatewayRouter(r *gin.Engine) {
	master := r.Group("/test")
	{
		master.POST("/write", func(c *gin.Context) {
			jsonData, _ := c.GetRawData()
			c.Header("Content-Type", "application/json; charset=utf-8")
			succ,_:=controllers.SaveRedis("testx",string(jsonData))
			c.String(200, succ)
		})
		master.GET("/read", func(c *gin.Context) {
			q,_:=c.GetQuery("key")
			succ:=controllers.GetRedis(q)
			c.String(200, succ)
		})
	}
}
