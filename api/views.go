package api

import (
	"encoding/json"
	"github.com/gusibi/nCov/internal/api/handler"
	d "github.com/gusibi/nCov/internal/dao"
	"net/http"

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
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	rss, err := handler.NewsRssListByFeed(news)
	if err!= nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/xml")
	w.Write(rss)
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
