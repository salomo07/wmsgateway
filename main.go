package main

import (
	"log"
	"wmsgateway/controllers"
	"wmsgateway/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "404", "message": "Endpoint isn't found"})
	})
	routers.GatewayRouter(r)
	r.Any("/:service/*path", func(c *gin.Context) {
		if c.Param("service") != "gateway" {
			// To Proxy
			log.Println("To Proxy")
			controllers.ForwardRequest(c.Param("service"), c)
		} else {
			log.Println("/:service/*path -> else")
			c.JSON(404, gin.H{"code": "404", "message": "Endpoint isn't found"})
		}
	})
	// r.NoRoute(func(c *gin.Context) {
	// 	c.JSON(404, gin.H{"code": "404", "message": "Endpoint isn't found"})
	// })
	// port := os.Getenv("PORT")
	r.Run(":7777")
}
