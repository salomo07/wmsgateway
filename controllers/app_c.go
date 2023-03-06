package controllers

import (
	"context"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
	"time"
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
func ForwardRequest(service string, c *gin.Context) {
	if service == "" {
		panic("Service tidak dikenali")
	}
	remote, err := url.Parse("http://localhost:7890/")
	if err != nil {
		log.Println("Error : ",err)
		// panic(err)
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

	proxy.ModifyResponse = func(resp *http.Response) error {
		log.Println("ModifyResponse : ",time.Since(start))
		// Disini tempat untuk save LOG Response
		if resp.StatusCode != 200{
			return errors.New(http.StatusText(resp.StatusCode))	 
		}		
		return nil
	}
	proxy.ServeHTTP(c.Writer, c.Request)
}
