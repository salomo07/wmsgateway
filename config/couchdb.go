package config

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

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

func FindDoc(db string, strquery string) string {
	return RequestByRoot("POST", db+"/_find", strquery)
}
func FindDocByCompany(basiccred string, db string, strquery string) string {
	return RequestByCompany(basiccred, "POST", db+"/_find", strquery)
}
func InsertDoc(path string, strquery string) string {
	return RequestByRoot("POST", path, strquery)
}
func InsertDocByCompany(basiccred string, db string, strquery string) string {
	return RequestByCompany(basiccred, "POST", db, strquery)
}
func UpdateDoc(path string, strquery string) string {
	return RequestByRoot("PUT", path, strquery)
}
func DeleteDoc(path string) string {
	return RequestByRoot("DELETE", path, "")
}
func CreateDatabase(db string) string {
	return RequestByRoot("PUT", db, "")
}
func CreateDatabase2(basiccred string, db string) string {
	return RequestByCompany(basiccred, "PUT", db, "")
}
func DeleteDatabase(db string) string {
	return RequestByRoot("DELETE", db, "")
}
func CreateUserDB(username string, strquery string) string {
	return RequestByRoot("PUT", "_users/org.couchdb.user:"+username, strquery)
}
func InsertAuthorDB(db string, strquery string) string {
	return RequestByRoot("PUT", db+"/_security", strquery)
}
func GetAuthorDB(db string, strquery string) string {
	return RequestByRoot("GET", db+"/_security", strquery)
}
func CreateReplication(strquery string) string {
	return RequestByRoot("POST", "_replicate", strquery)
}
func GetDataByView(db string, dsgname string, viewname string, str string) string {
	return RequestByRoot("POST", db+"/_design/"+dsgname+"/_view/"+viewname, str)
}
func BulkDocs(db string, strquery string) string {
	return RequestByRoot("POST", db+"/_bulk_docs", strquery)
}
func RequestByRoot(method string, pathURL string, strquery string) string {

	payload := strings.NewReader(strquery)
	client := &http.Client{}
	req, errCon := http.NewRequest(method, DB_STR_CON_DEFAULT+pathURL, payload)

	if errCon != nil {
		log.Println(errCon, " xxx")
		return string(`{"error":` + errCon.Error() + `}`)
	}
	req.Header.Add("Content-Type", "application/json")

	res, errDo := client.Do(req)
	if errDo != nil {
		log.Println(errDo, " yyy")
		return string(errDo.Error())
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err, " zzz")
		return string(`{"error":` + err.Error() + `}`)
	}

	return string(body)
}
func RequestByCompany(basiccred string, method string, pathURL string, strquery string) string {
	DB_URL := "https://" + basiccred + "@" + DB_BASE_URL + "/"
	payload := strings.NewReader(strquery)
	client := &http.Client{}
	req, errCon := http.NewRequest(method, DB_URL+pathURL, payload)

	if errCon != nil {
		log.Println(errCon, " xxx")
		return string(`{"error":` + errCon.Error() + `}`)
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Basic "+basiccred)
	res, errDo := client.Do(req)
	if errDo != nil {
		log.Println(errDo, " yyy")
		return string(`{"error":` + errDo.Error() + `}`)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err, " zzz")
		return string(`{"error":` + err.Error() + `}`)
	}
	return string(body)
}
