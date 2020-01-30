package data_collection

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

/*
抓取 https://news.qq.com/zt2020/page/feiyan.htm 数据，存入mysql
*/

const qqUrl = "https://news.qq.com/zt2020/page/feiyan.htm"

type qqDataCollection struct {
	Url string
}

func NewQQDataCollection(url string) *qqDataCollection{
	return &qqDataCollection{Url: url}
}

func (qdc *qqDataCollection) loadHtml() (*goquery.Document, error){
	res, err := http.Get(qdc.Url)
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

func (qdc *qqDataCollection) parseHtml(doc *goquery.Document) {
	doc.Find(".recentNumber .icbar").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the number and title
		number := s.Find(".number").Text()
		title := s.Find(".text").Text()
		fmt.Printf("Review %d: %s - %s\n", i, number, title)
	})
}



