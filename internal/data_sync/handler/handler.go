package handler

import (
	"log"

	d "github.com/gusibi/nCov/internal/dao"
	errors2 "github.com/gusibi/nCov/internal/dao/errors"
	"github.com/gusibi/nCov/internal/models"
)

/*
db conn save
*/

type DataSyncHandler struct {
	dao d.DataStore
}

func NewDataSyncHandler(dao d.DataStore) *DataSyncHandler {
	return &DataSyncHandler{dao: dao}
}

func (dh *DataSyncHandler) StoreNewsData(data []models.NewsData) error {
	log.Println("start store news data")
	for _, n := range data {
		_, err := dh.dao.GetNewsData(n.Title)
		if err != nil {
			if e, ok := err.(*errors2.DBError); ok && e.Code == "NotFound" {
				dh.dao.CreateNewsData(n)
			}
			continue
		}
	}
	return nil
}

func (dh *DataSyncHandler) StoreCountryData(data []models.CountryData) error {
	log.Println("start store country data")
	for _, country := range data {
		countryData, err := dh.dao.GetCountryData(country.CountryName)
		if err != nil {
			if e, ok := err.(*errors2.DBError); ok && e.Code == "NotFound" {
				dh.dao.CreateCountryData(country)
			}
			continue
		}
		// 如果数据有更新
		if countryData.ConfirmedCount != country.ConfirmedCount || countryData.SuspectedCount != country.SuspectedCount || countryData.CuredCount != country.CuredCount || countryData.DeadCount != country.DeadCount {
			dh.dao.CreateCountryData(country)
			dh.dao.UpdateCountryDataVersion(*countryData)
		}
	}
	return nil
}

func (dh *DataSyncHandler) StoreProvinceData(data []models.ProvinceData) error {
	log.Println("start store province data")
	for _, p := range data {
		//time.Sleep(10 * time.Second)
		//fmt.Println(p)
		pd, err := dh.dao.GetProvinceData(p.ProvinceShortName)
		if err != nil {
			if e, ok := err.(*errors2.DBError); ok && e.Code == "NotFound" {
				dh.dao.CreateProvinceData(p)
				dh.dao.BatchCreateCityData(p.ProvinceShortName, p.Cities)
			}
			continue
		}
		// update province data
		if pd.ConfirmedCount != p.ConfirmedCount || pd.SuspectedCount != p.SuspectedCount || pd.CuredCount != p.CuredCount || pd.DeadCount != p.DeadCount {
			dh.dao.CreateProvinceData(p)
			dh.dao.UpdateProvinceDataVersion(*pd)
		}
		dh.StoreCityData(p.ProvinceShortName, p.Cities)
	}
	return nil
}

func (dh *DataSyncHandler) StoreCityData(provinceName string, data []models.CityData) error {
	log.Printf("start store %s city data", provinceName)
	for _, c := range data {
		cd, err := dh.dao.GetCityData(provinceName, c.CityName)
		if err != nil {
			if e, ok := err.(*errors2.DBError); ok && e.Code == "NotFound" {
				dh.dao.CreateCityData(provinceName, c)
			}
			continue
		}
		// update province data
		if cd.ConfirmedCount != c.ConfirmedCount || cd.SuspectedCount != c.SuspectedCount || cd.CuredCount != c.CuredCount || cd.DeadCount != c.DeadCount {
			dh.dao.CreateCityData(provinceName, c)
			dh.dao.UpdateCityDataVersion(*cd)
		}
	}
	return nil
}
