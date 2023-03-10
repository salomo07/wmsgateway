package routers

import (
	"log"
	"wmsgateway/controllers"
	_ "wmsgateway/controllers"

	"github.com/gin-gonic/gin"
)

var routerGroupV1 string

type RequestLog struct {
	Ip            string `json:"ip"`
	UserAgent     string `json:"devicename"`
	Url           string `json:"url"`
	Authorization string `json:"authorization"`
	Payload       string `json:"payload"`
}

func GatewayRouter(r *gin.Engine) {
	api1 := r.Group("gateway/api/v1")
	{
		log.Println("gateway/api/v1")
		api1.POST("/setredis?:key", func(c *gin.Context) {
			c.Header("Content-Type", "application/json; charset=utf-8")
			jsonData, err := c.GetRawData()
			if err != nil {
				c.JSON(400, map[string]interface{}{"error": "Bad request body"})
			} else if c.Query("key") == "" {
				c.JSON(400, map[string]interface{}{"error": "Key is not found"})
			} else {
				xxx, errSave := controllers.SaveRedis(c.Query("key"), string(jsonData))
				if errSave != "" {
					c.JSON(500, map[string]interface{}{"error": "Error occured when saving to redis server"})
				} else {
					c.String(200, xxx)
				}
			}

		})
		api1.GET("/getredis?:key", func(c *gin.Context) {
			c.Header("Content-Type", "application/json; charset=utf-8")
			if c.Query("key") == "" {
				c.JSON(400, map[string]interface{}{"error": "Key is not found"})
			} else {
				xxx, errGet := controllers.GetRedis(c.Query("key"))
				if errGet != "" {
					c.JSON(500, map[string]interface{}{"error": errGet})
				} else {
					c.String(200, xxx)
				}
			}
		})
	}
	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "404", "message": "Gateway endpoint isn't found"})
	})
	// api2 := r.Group("gateway/api/v2")
	// {
	// 	api2.GET("/authenticate", func(c *gin.Context) {
	// 		log.Println("authenticate")
	// 	})
	// }

}
