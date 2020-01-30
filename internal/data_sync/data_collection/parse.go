package data_collection

import (
	"github.com/PuerkitoBio/goquery"
	models2 "github.com/gusibi/nCov/internal/models"
)

type DataCollection interface {
	LoadHtml() (*goquery.Document, error)
	GetCountryData(doc *goquery.Document) []models2.CountryData
	GetProvinceData(doc *goquery.Document) []models2.ProvinceData
	GetNewsData(doc *goquery.Document) []models2.NewsData
	parseHtml(doc *goquery.Document)
}
