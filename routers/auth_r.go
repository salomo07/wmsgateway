package routers

import (
	_"wmsgateway/controllers"
	"github.com/gin-gonic/gin"
	"log"
	"strings"
	"net/http"
	"net/http/httputil"
	"net/url"
)
var routerGroupV1 string
type RequestLog struct {
	Ip  string `json:"ip"`
	UserAgent string `json:"devicename"`
	Url string `json:"url"`
	Authorization string `json:"authorization"`
	Payload string `json:"payload"`
}


func GatewayRouter(r *gin.Engine) {
	r.LoadHTMLGlob("templates/*")
	// test := r.Group("/test")
	// {
	// 	test.POST("/write", func(c *gin.Context) {
	// 		// jsonData, _ := c.GetRawData()
	// 		xxx:=string(c.Request.URL.Path)
	// 		req:=RequestLog{Ip:c.ClientIP(),UserAgent:c.Request.Header.Get("User-Agent"),Url:xxx}
	// 		// log.Println(c.ClientIP(),c.Request.URL,c.Request.Header.Get("User-Agent"))
			
	// 		c.Header("Content-Type", "application/json; charset=utf-8")
	// 		// succ,_:=controllers.SaveRedis("testx",string(jsonData))
	// 		// c.String(200, succ)
	// 		log.Println(req)
	// 	})
	// 	test.GET("/read", func(c *gin.Context) {
	// 		q,_:=c.GetQuery("key")
	// 		succ:=controllers.GetRedis(q)
	// 		c.String(200, succ)
	// 	})
	// 	test.GET("/socket", func(c *gin.Context) {
	// 		log.Println("Ini socket")
	// 		c.HTML(200, "index.html", gin.H{})
	// 		log.Println("Ini socketxxx")
	// 	})
	// }
	// r.GET("/api/v1/:service/*path", func(c *gin.Context) {
	// 	log.Println(c.Param("service"))
	// 	log.Println(c.Param("path"))
	// })
	routerGroupV1="/api/v1"
	api1 := r.Group(routerGroupV1)
	{
		api1.GET("/:service/*path", func(c *gin.Context) {
			log.Println(c.Param("service"))
			remote, err := url.Parse("localhost:7890")
			if err != nil {
				panic(err)
			}
			log.Println(c.Param("service"))
			urlPath:=strings.Replace(c.Request.URL.String(), routerGroupV1, "",1)
			log.Println(urlPath)
			proxy := httputil.NewSingleHostReverseProxy(remote)
			proxy.Director = func(req *http.Request) {
				req.Header = c.Request.Header
				req.Host = remote.Host
				req.URL.Scheme = remote.Scheme
				req.URL.Host = remote.Host
				req.URL.Path = c.Param("xxx")
			}
		})
	}
}
