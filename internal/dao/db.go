package dao

import (
	"log"
	"time"

	"github.com/gusibi/nCov/internal/tools"

	dbErr "github.com/gusibi/nCov/internal/dao/errors"
	"github.com/gusibi/nCov/internal/models"

	"github.com/jinzhu/gorm"
)

type DataDao struct {
	db *gorm.DB
}

func NewDataDao(db *gorm.DB) *DataDao {
	return &DataDao{db: db}
}

func (dao *DataDao) GetCountryData(countryName string) (*models.CountryDataModel, error) {
	countryData := new(models.CountryDataModel)
	db := dao.db.Where("version = 0 and country = ?", countryName).Find(&countryData)
	if db.Error != nil {
		log.Printf("UpdateCountryData | get country data error | country=%s, err=%v", countryName, db.Error)
		if db.Error.Error() == "record not found" {
			return nil, dbErr.NewDBError("NotFound", "record not found")
		}
		return nil, db.Error
	}
	return countryData, nil
}

func (dao *DataDao) CreateCountryData(data models.CountryData) error {
	countryData := &models.CountryDataModel{
		CountryName:    data.CountryName,
		ConfirmedCount: data.ConfirmedCount,
		SuspectedCount: data.SuspectedCount,
		CuredCount:     data.CuredCount,
		DeadCount:      data.DeadCount,
		Version:        "0",
		CreatedAt:      time.Now().Unix(),
	}
	db := dao.db.Create(countryData)
	if db.Error != nil {
		log.Printf("CreateCountryData | insert country data error | err=%v", db.Error)
		return db.Error
	}
	return nil
}

func (dao *DataDao) UpdateCountryDataVersion(countryData models.CountryDataModel) error {
	version := tools.DateToVersion(countryData.CreatedAt)
	db := dao.db.Model(countryData).Update("version", version)
	if db.Error != nil {
		log.Printf("UpdateCountryDataVersion | update country version error | err=%v", db.Error)
		return db.Error
	}
	return nil
}

func (dao *DataDao) GetProvinceData(provinceName string) (*models.ProvinceDataModel, error) {
	pd := new(models.ProvinceDataModel)
	db := dao.db.Where("version = 0 and province = ?", provinceName).Find(&pd)
	if db.Error != nil {
		log.Printf("GetProvinceData | get province data error | province=%s, err=%v", provinceName, db.Error)
		if db.Error.Error() == "record not found" {
			return nil, dbErr.NewDBError("NotFound", "record not found")
		}
		return nil, db.Error
	}
	return pd, nil
}

func (dao *DataDao) GetCityData(provinceName, cityName string) (*models.CityDataModel, error) {
	cd := new(models.CityDataModel)
	db := dao.db.Where("version = 0 and province = ? and city = ?", provinceName, cityName).Find(&cd)
	if db.Error != nil {
		log.Printf("GetCityData | get city data error | city=%s, err=%v", cityName, db.Error)
		if db.Error.Error() == "record not found" {
			return nil, dbErr.NewDBError("NotFound", "record not found")
		}
		return nil, db.Error
	}
	return cd, nil
}

func (dao *DataDao) CreateProvinceData(data models.ProvinceData) error {
	countryData := &models.ProvinceDataModel{
		Country:        "中国",
		Province:       data.ProvinceShortName,
		ConfirmedCount: data.ConfirmedCount,
		SuspectedCount: data.SuspectedCount,
		CuredCount:     data.CuredCount,
		DeadCount:      data.DeadCount,
		Version:        "0",
		CreatedAt:      time.Now().Unix(),
	}
	db := dao.db.Create(countryData)
	if db.Error != nil {
		log.Printf("CreateProvinceData | insert province data error | err=%v", db.Error)
		return db.Error
	}
	return nil
}

func (dao *DataDao) CreateCityData(provinceName string, data models.CityData) error {
	countryData := &models.CityDataModel{
		Province:       provinceName,
		City:           data.CityName,
		ConfirmedCount: data.ConfirmedCount,
		SuspectedCount: data.SuspectedCount,
		CuredCount:     data.CuredCount,
		DeadCount:      data.DeadCount,
		Version:        "0",
		CreatedAt:      time.Now().Unix(),
	}
	db := dao.db.Create(countryData)
	if db.Error != nil {
		log.Printf("CreateProvinceData | insert province data error | err=%v \n", db.Error)
		return db.Error
	}
	return nil
}

func (dao *DataDao) BatchCreateCityData(provinceName string, data []models.CityData) error {
	var err error
	for _, city := range data {
		err = dao.CreateCityData(provinceName, city)
		if err != nil {
			log.Printf("BatchCreateCityData | insert city data error | data=%v, err=%v \n", city, err)
		}
	}
	return err
}

func (dao *DataDao) UpdateProvinceDataVersion(pd models.ProvinceDataModel) error {
	version := tools.DateToVersion(pd.CreatedAt)
	db := dao.db.Model(pd).Update("version", version)
	if db.Error != nil {
		log.Printf("UpdateProvinceDataVersion | update province version error | err=%v \n", db.Error)
		return db.Error
	}
	return nil
}

func (dao *DataDao) UpdateCityDataVersion(cd models.CityDataModel) error {
	version := tools.DateToVersion(cd.CreatedAt)
	db := dao.db.Model(cd).Update("version", version)
	if db.Error != nil {
		log.Printf("UpdateCityDataVersion | update cd version error | err=%v \n", db.Error)
		return db.Error
	}
	return nil
}

func (dao *DataDao) GetNewsDataByTitle(title string) (*models.NewsModel, error) {
	nd := new(models.NewsModel)
	db := dao.db.Where("title = ?", title).Find(&nd)
	if db.Error != nil {
		log.Printf("GetNewsDataByTitle | get news data error | title=%s, err=%v", title, db.Error)
		if db.Error.Error() == "record not found" {
			return nil, dbErr.NewDBError("NotFound", db.Error.Error())
		}
		return nil, db.Error
	}
	return nd, nil
}

func (dao *DataDao) GetNewsList() ([]*models.NewsModel, error) {
	nl := make([]*models.NewsModel, 0)
	db := dao.db.Order("created_at DESC").Limit(100).Find(&nl)
	if db.Error != nil {
		log.Printf("GetNewsList | get news list error | err=%v\n", db.Error)
		return nil, dbErr.NewDBError("QueryError", db.Error.Error())
	}
	return nl, nil
}

func (dao *DataDao) GetNewsListByProvince(provinceName string) ([]*models.NewsModel, error) {
	nl := make([]*models.NewsModel, 0)
	db := dao.db.Where("province = ?", provinceName).Order("created_at DESC").Limit(100).Find(&nl)
	if db.Error != nil {
		log.Printf("GetNewsList | get news list error | err=%v\n", db.Error)
		return nil, dbErr.NewDBError("QueryError", db.Error.Error())
	}
	return nl, nil
}

func (dao *DataDao) CreateNewsData(data models.NewsData) error {
	newsData := &models.NewsModel{
		//Id:          data.Id,
		Province:    data.ProvinceName,
		Title:       data.Title,
		Summary:     data.Summary,
		Source:      data.Source,
		SourceUrl:   data.SourceUrl,
		PublishedAt: data.PublishedDate / 1000,
		CreatedAt:   time.Now().Unix(),
	}
	db := dao.db.Create(newsData)
	if db.Error != nil {
		log.Printf("CreateNewsData | insert news data error | err=%v", db.Error)
		return db.Error
	}
	return nil
}
