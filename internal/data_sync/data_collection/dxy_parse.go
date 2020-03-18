package data_collection

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	models2 "github.com/gusibi/nCov/internal/models"

	"github.com/PuerkitoBio/goquery"
)

const (
	dxyUrl               = "https://3g.dxy.cn/newh5/view/pneumonia"
	getAreaStat          = "getAreaStat"
	getStatisticsService = "getStatisticsService"
	indexRecommendList   = "getIndexRecommendList"
	getOtherCountryList  = "getListByCountryTypeService2true"
	newsList             = "getTimelineService2"
	newsListCn             = "getTimelineServiceundefined"
)

type dxyDataCollection struct {
	Url string
}

func NewDxyDataCollection() *dxyDataCollection {
	return &dxyDataCollection{Url: dxyUrl}
}

func (ddc *dxyDataCollection) LoadHtml() (*goquery.Document, error) {
	res, err := http.Get(ddc.Url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	return doc, err
}

func (ddc *dxyDataCollection) GetProvinceData(doc *goquery.Document) []models2.ProvinceData {
	dID := fmt.Sprintf("#%s", getAreaStat)
	areaStatNode := doc.Find(dID)
	areaStat := areaStatNode.Text()
	startIndex := strings.Index(areaStat, "[")
	endIndex := strings.LastIndex(areaStat, "]")
	areaStat = areaStat[startIndex : endIndex+1]
	provinceDataList := make([]models2.ProvinceData, 0)
	if err := json.Unmarshal([]byte(areaStat), &provinceDataList); err != nil {
		log.Fatal("json decode province data error", err)
	}
	return provinceDataList
}

func (ddc *dxyDataCollection) GetCountryData(doc *goquery.Document) []models2.CountryData {
	dID := fmt.Sprintf("#%s", getOtherCountryList)
	areaStatNode := doc.Find(dID)
	areaStat := areaStatNode.Text()
	startIndex := strings.Index(areaStat, "[")
	endIndex := strings.LastIndex(areaStat, "]")
	fmt.Println(startIndex, endIndex)
	areaStat = areaStat[startIndex : endIndex+1]
	countryData := make([]models2.CountryData, 0)
	if err := json.Unmarshal([]byte(areaStat), &countryData); err != nil {
		log.Fatal("json decode country data error", err)
	}
	return countryData
}

func (ddc *dxyDataCollection) GetNewsDataByKey(doc *goquery.Document, key string) []models2.NewsData {
	nID := fmt.Sprintf("#%s", key)
	newsNode := doc.Find(nID)
	newsStat := newsNode.Text()
	startIndex := strings.Index(newsStat, "[")
	endIndex := strings.LastIndex(newsStat, "]")
	newsStat = newsStat[startIndex : endIndex+1]
	newsData := make([]models2.NewsData, 0)
	if err := json.Unmarshal([]byte(newsStat), &newsData); err != nil {
		log.Fatal("json decode news data error", err)
	}
	return newsData
}

func (ddc *dxyDataCollection) GetNewsData(doc *goquery.Document) []models2.NewsData {
	keys := []string{newsList, newsListCn}
	newsData := make([]models2.NewsData, 0)
	for _, key := range keys{
		news := ddc.GetNewsDataByKey(doc, key)
		for _, n := range news{
			newsData = append(newsData, n)
		}
	}
	return newsData
}

func (ddc *dxyDataCollection) parseHtml(doc *goquery.Document) {
	provinceData := ddc.GetProvinceData(doc)
	fmt.Println(provinceData)
	countryData := ddc.GetCountryData(doc)
	fmt.Println(countryData)
}
