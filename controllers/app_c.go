package controllers
import ("github.com/joho/godotenv")

var DB_STR_CON_DEFAULT string
var DB_BASE_URL string

func init() {
	er := godotenv.Load()
	if er != nil {
		panic("Fail to load .env file")
	}
	DB_STR_CON_DEFAULT = os.Getenv("DB_STR_CON")
	DB_BASE_URL = os.Getenv("HOST_DB")
}
func SaveRedis() {

}
