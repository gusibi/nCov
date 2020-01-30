package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/tencentyun/scf-go-lib/cloudfunction"

	"github.com/jinzhu/gorm"

	"github.com/gusibi/nCov/data_sync/data_store"

	"github.com/PuerkitoBio/goquery"
	dc "github.com/gusibi/nCov/data_sync/data_collection"
	model "github.com/gusibi/nCov/data_sync/models"
)

func EnvGet(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func getCountryData(c dc.DataCollection, doc *goquery.Document) []model.CountryData {
	return c.GetCountryData(doc)
}

func getProvinceData(c dc.DataCollection, doc *goquery.Document) []model.ProvinceData {
	return c.GetProvinceData(doc)
}

func getNewsData(c dc.DataCollection, doc *goquery.Document) []model.NewsData {
	return c.GetNewsData(doc)
}

func dbConfig() data_store.DBConfig {
	dbUser := EnvGet("DBUSER", "ncov")
	dbPwd := EnvGet("DBPWD", "ncov")
	dbHost := EnvGet("DBHOST", "127.0.0.1")
	dbPort := EnvGet("DBPORT", "3306")
	dbName := EnvGet("DBNAME", "ncov")
	return data_store.DBConfig{
		DSN:          fmt.Sprintf("%s:%s@(%s:%s)/%s?timeout=1000ms&readTimeout=10000ms&charset=utf8mb4", dbUser, dbPwd, dbHost, dbPort, dbName),
		ConnMaxAge:   300,
		MaxIdleConns: 10,
		MaxOpenConns: 10,
	}
}

var DBConn *gorm.DB

func init() {
	dbConf := dbConfig()
	dbConn, err := data_store.GetDBConnect(dbConf)
	if err != nil {
		log.Fatal("conn db error")
	}
	DBConn = dbConn
}

func handle() {
	dci := dc.NewDxyDataCollection()
	doc, _ := dci.LoadHtml()
	dao := data_store.NewDataDao(DBConn)

	targetList := strings.Split(EnvGet("SyncTarget", "data,news"), ",")
	for _, target := range targetList {
		if target == "data" {
			countryDataList := getCountryData(dci, doc)
			fmt.Println("CountryData: ", len(countryDataList))
			dao.StoreCountryData(countryDataList)

			provinceDataList := getProvinceData(dci, doc)
			fmt.Println("ProvinceData: ", len(provinceDataList))
			dao.StoreProvinceData(provinceDataList)
		} else if target == "news" {
			newsDataList := getNewsData(dci, doc)
			fmt.Println("newsData: ", len(newsDataList))
			dao.StoreNewsData(newsDataList)
		}
	}
}

func main() {
	//handle()
	cloudfunction.Start(handle)
}
