package api

import (
	"github.com/gusibi/nCov/internal/dao"
	"github.com/gusibi/nCov/internal/tools"
	"github.com/julienschmidt/httprouter"
)

func GetRouters() *httprouter.Router {
	router := httprouter.New()
	d := dao.NewDataDao(tools.DBConn)

	jsonHandler := NewJsonHandler(d)
	router.GET("/api/ncov/news", jsonHandler.NewsListHandler)
	router.GET("/api/ncov/news/:province", jsonHandler.NewsListByProvinceHandler)

	xmlHandler := NewXmlHandler(d)
	router.GET("/ncov/news.xml", xmlHandler.NewsListHandler)
	//router.GET("/ncov/:province/news.xml", xmlHandler.NewsListHandler)
	return router
}
