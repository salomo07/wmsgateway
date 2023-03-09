package controllers

import (
	"context"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

var rdb *redis.Client
var ctx = context.Background()
var REDIS_USER string
var REDIS_PASS string
var REDIS_PORT string
var REDIS_HOST_CLOUD string
var ERROR_LOAD_ENV string

func init() {
	er := godotenv.Load()
	if er != nil {
		panic("Fail to load .env file")
	}
	REDIS_HOST_CLOUD = os.Getenv("REDIS_HOST_CLOUD")
	REDIS_USER = os.Getenv("REDIS_USER")
	REDIS_PASS = os.Getenv("REDIS_PASS")
	REDIS_PORT = os.Getenv("REDIS_PORT")
	opt, _ := redis.ParseURL("redis://" + REDIS_USER + ":" + REDIS_PASS + "@" + REDIS_HOST_CLOUD + ":" + REDIS_PORT)
	rdb = redis.NewClient(opt)
}
func SaveRedis(key string, data string) (string, string) {
	err := rdb.Set(ctx, key, data, 0).Err()
	if err != nil {
		return `{"ok":true}`, `{"error":"` + err.Error() + `"}`
	}
	return `{"ok":true}`, ""
}
func GetRedis(key string) string {
	if ERROR_LOAD_ENV != "" {
		return `{"error":"` + ERROR_LOAD_ENV + `"}`
	}
	val, err := rdb.Get(ctx, key).Result()
	if err != nil {
		return `{"error":"` + err.Error() + `"}`
	}
	return val
}
func GetBestHost() string {
	return "http://localhost:7890"
}
func ForwardRequest(service string, c *gin.Context) {

	if service == "" {
		panic("Service tidak dikenali")
	}
	bestHost := GetBestHost()
	remote, err := url.Parse(bestHost + c.Request.URL.String())

	if err != nil { //Gagal parse URL
		c.JSON(502, map[string]interface{}{"error": err})
	}
	start := time.Now()
	proxy := httputil.NewSingleHostReverseProxy(remote)
	proxy.Director = func(req *http.Request) {
		req.Header = c.Request.Header
		req.URL.Scheme = remote.Scheme
		req.URL.Host = remote.Host
		req.URL.Path = c.Request.URL.String()
		// Disini tempat untuk save LOG Request
	}
	proxy.ErrorHandler = func(res http.ResponseWriter, req *http.Request, err error) {
		log.Println("ErrorHandler", err)
	}
	proxy.ModifyResponse = func(resp *http.Response) error {
		log.Println("ModifyResponse : ", time.Since(start), "StatusCode : ", resp.Body)
		// Disini tempat untuk save LOG Response
		return nil
	}

	proxy.ServeHTTP(c.Writer, c.Request)
}
