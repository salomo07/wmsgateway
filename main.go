package main

import (
	_ "log"
	"wmsgateway/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	routers.GatewayRouter(r)

	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "404", "message": "Page not found"})
	})
	// port := os.Getenv("PORT")
	r.Run(":7777")
}
