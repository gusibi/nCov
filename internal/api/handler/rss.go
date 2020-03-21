package handler

import (
	"encoding/xml"
	"fmt"
	"strings"
	"time"

	"github.com/gorilla/feeds"
	"github.com/gusibi/nCov/internal/models"
	"github.com/gusibi/nCov/internal/tools"
)

const (
	BaseUrl = "http://service-5lpzkj7x-1254035985.bj.apigw.tencentcs.com/release/ncov/news"
	RssUrl  = "http://service-5lpzkj7x-1254035985.bj.apigw.tencentcs.com/release/ncov/news.xml"
)

func getLink(new *models.NewsModel) string {
	if new.SourceUrl != "" {
		return new.SourceUrl
	}
	defaultLinks := map[string]string{
		"BBC":          "https://twitter.com/bbc",
		"CNN":          "https://twitter.com/cnn",
		"RETURNS":      "https://www.reuters.com",
		"FT":           "https://www.ft.com",
		"DXDOCTOR_":    "https://twitter.com/DXDoctor_",
		"THE GUARDIAN": "https://twitter.com/guardian",
		"DAILYMAIL":    "https://twitter.com/mailonline",
	}
	if link, ok := defaultLinks[strings.ToUpper(new.SourceUrl)]; ok {
		return link
	}
	return "https://3g.dxy.cn/newh5/view/pneumonia"
}

func NewsRssListByFeed(news []*models.NewsModel) ([]byte, error) {
	feed := &feeds.Feed{
		Title:       "新型肺炎疫情实时动态播报",
		Link:        &feeds.Link{Href: RssUrl, Rel: "self"},
		Description: "全国新型肺炎疫情实时动态播报-数据来自丁香园-https://3g.dxy.cn/newh5/view/pneumonias",
		Author:      &feeds.Author{Name: "gusibi", Email: "hi@gusibi.mobi"},
		Created:     time.Now(),
	}
	items := []*feeds.Item{}
	var pubAt int64
	for _, new := range news {
		if new.PublishedAt == pubAt {
			pubAt = new.PublishedAt - 1
		} else {
			pubAt = new.PublishedAt
		}
		item := &feeds.Item{
			Id:     fmt.Sprintf("%s/%d", BaseUrl, new.Id),
			Title:  new.Title,
			Author: &feeds.Author{Name: new.Source},
			Link:   &feeds.Link{Href: getLink(new), Rel: "self"},
			//Source:      &feeds.Link{Href: getLink(new), Rel: "self"},
			Created:     tools.TimeStampToDate(new.CreatedAt),
			Description: new.Summary,
			Updated:     tools.BjTimeStampToDate(new.PublishedAt),
			Content:     new.Summary,
		}
		items = append(items, item)
	}
	feed.Items = items
	rssStr, err := feed.ToAtom()
	if err != nil {
		return nil, err
	}
	return []byte(rssStr), err
}

func NewsRssList(news []*models.NewsModel) ([]byte, error) {
	items := []rssItem{}
	for _, new := range news {
		item := rssItem{
			Title:       new.Title,
			Link:        new.SourceUrl,
			Description: new.Summary,
			PubDate:     tools.BjTimeStampToBJString(new.PublishedAt, time.RFC822),
			//Author:      new.Source,
		}
		items = append(items, item)
	}
	feed := rss{
		Version:     "2.0",
		Description: "全国新型肺炎疫情实时动态播报-数据来自丁香园-https://3g.dxy.cn/newh5/view/pneumonia",
		Link:        "https://3g.dxy.cn/newh5/view/pneumonia",
		Title:       "全国新型肺炎疫情实时动态播报",
		Item:        items,
	}

	x, err := xml.MarshalIndent(feed, "", "  ")
	if err != nil {
		return nil, err
	}
	return x, nil
}
