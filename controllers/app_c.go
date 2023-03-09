package controllers

import (
	"context"
	"encoding/json"
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

type Server struct {
	Host    string  `json:"host"`
	Service string  `json:"service"`
	Time    float32 `json:"time"`
}

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
		return `{"ok":true}`, err.Error()
	}
	return `{"ok":true}`, ""
}
func GetRedis(key string) (string, string) {
	val, err := rdb.Get(ctx, key).Result()
	if err != nil {
		return "", err.Error()
	}
	return val, ""
}
func GetBestHost(c *gin.Context) string {
	var srv []Server
	// var b Server
	data, errGet := GetRedis("serverList")
	if errGet != "" {
		c.JSON(500, map[string]interface{}{"error": errGet})
	}
	err := json.Unmarshal([]byte(string(data)), &srv)
	if err != nil {
		c.JSON(500, map[string]interface{}{"error": err.Error()})
	}
	// var t float32
	// for i, s := range srv {
	// 	value, err := strconv.ParseFloat(myString, 32)
	// 	if i ==0{
	// 		b=s
	// 	}
	// 	log.Println(i, s)
	// }
	return "http://localhost:7890"
}
func ForwardRequest(service string, c *gin.Context) {

	if service == "" {
		panic("Service tidak dikenali")
	}
	bestHost := GetBestHost(c)
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
