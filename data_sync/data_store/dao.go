package data_store

import "github.com/gusibi/nCov/data_sync/models"

type DataStore interface {
	StoreCountryData(data models.CountryData) error
	StoreProvinceData(data models.ProvinceData) error
}
