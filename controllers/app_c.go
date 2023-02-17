package controllers
import (
	"github.com/joho/godotenv"
	"context"
	"os"
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
	opt, _ := redis.ParseURL("redis://"+REDIS_USER+":"+REDIS_PASS+"@"+REDIS_HOST_CLOUD+":"+REDIS_PORT)
	rdb = redis.NewClient(opt)
}
func SaveRedis(key string, data string) (string,string){
	err := rdb.Set(ctx, key, data, 0).Err()
	if err != nil {
		return `{"ok":true}`, `{"error":"` + err.Error() + `"}`
	}
	return `{"ok":true}`, ""
}
func GetRedis(key string) string{
	if ERROR_LOAD_ENV != "" {
		return `{"error":"` + ERROR_LOAD_ENV + `"}`
	}
	val, err := rdb.Get(ctx, key).Result()
	if err != nil {
		return `{"error":"` + err.Error() + `"}`
	}
	return val
}