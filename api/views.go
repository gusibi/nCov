package api

import (
	"encoding/json"
	"encoding/xml"
	"net/http"

	"github.com/gusibi/nCov/internal/tools"

	d "github.com/gusibi/nCov/internal/dao"

	"github.com/julienschmidt/httprouter"
)

type XmlHandler struct {
	dao d.DataStore
}

func NewXmlHandler(dao d.DataStore) *XmlHandler {
	return &XmlHandler{dao: dao}
}

func (h *XmlHandler) NewsListHandler(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	news, err := h.dao.GetNewsList()
	items := []rssItem{}
	for _, new := range news {
		item := rssItem{
			Title:       new.Title,
			Link:        new.SourceUrl,
			Description: new.Summary,
			PubDate:     tools.BjTimeStampToBJString(new.PublishedAt, "2006-01-02 15:04:05"),
			Author:      new.Source,
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
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/xml")
	w.Write(x)
}

type JsonHandler struct {
	dao d.DataStore
}

func NewJsonHandler(dao d.DataStore) *JsonHandler {
	return &JsonHandler{dao: dao}
}

func JsonResponse(body []byte, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.Write(body)
}

func (h *JsonHandler) NewsListHandler(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	news, err := h.dao.GetNewsList()

	jsonString, err := json.Marshal(news)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	JsonResponse(jsonString, w)
}

func (h *JsonHandler) NewsListByProvinceHandler(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	province := ps.ByName("province")
	news, err := h.dao.GetNewsListByProvince(province)

	jsonString, err := json.Marshal(news)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	JsonResponse(jsonString, w)
}
