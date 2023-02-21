package routers

import (
	"wmsgateway/controllers"
	"github.com/gin-gonic/gin"
)
type RequestLog struct {
	Ip  string `json:"ip"`
	DeviceName string `json:"devicename"`
	Url string `json:"url"`
	Payload string `json:"payload"`
}

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
