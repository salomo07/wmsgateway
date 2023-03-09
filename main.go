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

	// port := os.Getenv("PORT")
	r.Run(":7777")
}
