package data_collection

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/gusibi/nCov/data_sync/models"
)

type DataCollection interface {
	LoadHtml() (*goquery.Document, error)
	GetCountryData(doc *goquery.Document) []models.CountryData
	GetProvinceData(doc *goquery.Document) []models.ProvinceData
	GetNewsData(doc *goquery.Document) []models.NewsData
	parseHtml(doc *goquery.Document)
}
