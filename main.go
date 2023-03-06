package main

import (
	_ "log"
	"wmsgateway/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Any("/:service/*path", func(c *gin.Context) {
		controllers.ForwardRequest(c.Param("service"), c)
	})
	// r.Any("/*proxyPath", controllers.ForwardRequest)

	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "404", "message": "Page not found"})
	})
	// port := os.Getenv("PORT")
	r.Run(":7777")
}
