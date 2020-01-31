package dao

import (
	"github.com/gusibi/nCov/internal/models"
)

type DataStore interface {
	GetCountryData(countryName string) (*models.CountryDataModel, error)
	CreateCountryData(data models.CountryData) error
	UpdateCountryDataVersion(countryData models.CountryDataModel) error

	GetProvinceData(provinceName string) (*models.ProvinceDataModel, error)
	CreateProvinceData(data models.ProvinceData) error
	UpdateProvinceDataVersion(pd models.ProvinceDataModel) error

	GetCityData(provinceName, cityName string) (*models.CityDataModel, error)
	CreateCityData(provinceName string, data models.CityData) error
	BatchCreateCityData(provinceName string, data []models.CityData) error
	UpdateCityDataVersion(cd models.CityDataModel) error

	GetNewsDataByTitle(title string) (*models.NewsModel, error)
	GetNewsList() ([]*models.NewsModel, error)
	GetNewsListByProvince(provinceName string) ([]*models.NewsModel, error)
	CreateNewsData(data models.NewsData) error
}
