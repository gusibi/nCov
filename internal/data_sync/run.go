package data_sync

import (
	"fmt"
	"strings"

	"github.com/gusibi/nCov/internal/tools"

	"github.com/PuerkitoBio/goquery"
	"github.com/gusibi/nCov/internal/dao"
	"github.com/gusibi/nCov/internal/data_sync/data_collection"
	"github.com/gusibi/nCov/internal/data_sync/handler"
	"github.com/gusibi/nCov/internal/models"
)

func getCountryData(c data_collection.DataCollection, doc *goquery.Document) []models.CountryData {
	return c.GetCountryData(doc)
}

func getProvinceData(c data_collection.DataCollection, doc *goquery.Document) []models.ProvinceData {
	return c.GetProvinceData(doc)
}

func getNewsData(c data_collection.DataCollection, doc *goquery.Document) []models.NewsData {
	return c.GetNewsData(doc)
}

func Run() {
	dci := data_collection.NewDxyDataCollection()
	doc, _ := dci.LoadHtml()

	d := dao.NewDataDao(tools.DBConn)
	handler := handler.NewDataSyncHandler(d)

	targetList := strings.Split(tools.EnvGet("SyncTarget", "data,news"), ",")
	for _, target := range targetList {
		if target == "data" {
			countryDataList := getCountryData(dci, doc)
			fmt.Println("CountryData: ", len(countryDataList))
			handler.StoreCountryData(countryDataList)

			provinceDataList := getProvinceData(dci, doc)
			fmt.Println("ProvinceData: ", len(provinceDataList))
			handler.StoreProvinceData(provinceDataList)
		} else if target == "news" {
			newsDataList := getNewsData(dci, doc)
			fmt.Println("newsData: ", len(newsDataList))
			handler.StoreNewsData(newsDataList)
		}
	}
}
