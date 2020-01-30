package data_store

import (
	"fmt"
	"log"
	"time"

	dberr "github.com/gusibi/nCov/data_sync/data_store/errors"
	"github.com/gusibi/nCov/data_sync/models"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

/*
db conn save
*/

type DBConfig struct {
	DSN          string
	ConnMaxAge   int
	MaxIdleConns int
	MaxOpenConns int
}

func GetDBConnect(dbConfig DBConfig) (*gorm.DB, error) {
	return Connect(
		dbConfig.DSN,
		dbConfig.MaxIdleConns,
		dbConfig.MaxOpenConns,
		time.Duration(dbConfig.ConnMaxAge)*time.Second,
	)
}

func Connect(dsn string, maxIdleConns int, maxOpenConns int, connMaxLifetime time.Duration) (*gorm.DB, error) {

	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		msg := fmt.Sprintf("gorm.Open database fail|DSN: %s", dsn)
		fmt.Println(msg, err)
		return nil, errors.WithMessage(err, msg)
	}

	db.DB().SetMaxIdleConns(maxIdleConns)
	db.DB().SetConnMaxLifetime(connMaxLifetime)
	db.DB().SetMaxOpenConns(maxOpenConns)
	db.LogMode(false)
	err = db.DB().Ping()
	if err != nil {
		log.Printf("Ping mysql failed, DSN: %s, maxIdleConns: %d, maxOpenConns: %d, connMaxLifetime: %d", dsn, maxIdleConns, maxOpenConns, maxOpenConns)
		return nil, errors.WithMessage(err, "Ping")
	}

	return db, nil
}

func DateToVersion(ts int64) string {
	date := time.Unix(ts, 0)
	return date.Format("2006-01-02T15:04:05")
}

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
			return nil, dberr.NewDBError("NotFound", "record not found")
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
	version := DateToVersion(countryData.CreatedAt)
	db := dao.db.Model(countryData).Update("version", version)
	if db.Error != nil {
		log.Printf("UpdateCountryDataVersion | update country version error | err=%v", db.Error)
		return db.Error
	}
	return nil
}

func (dao *DataDao) StoreCountryData(data []models.CountryData) error {
	log.Println("start store country data")
	for _, country := range data {
		countryData, err := dao.GetCountryData(country.CountryName)
		if err != nil {
			if e, ok := err.(*dberr.DBError); ok && e.Code == "NotFound" {
				dao.CreateCountryData(country)
			}
			continue
		}
		// 如果数据有更新
		if countryData.ConfirmedCount != country.ConfirmedCount || countryData.SuspectedCount != country.SuspectedCount || countryData.CuredCount != country.CuredCount || countryData.DeadCount != country.DeadCount {
			dao.CreateCountryData(country)
			dao.UpdateCountryDataVersion(*countryData)
		}
	}
	return nil
}

func (dao *DataDao) GetProvinceData(provinceName string) (*models.ProvinceDataModel, error) {
	pd := new(models.ProvinceDataModel)
	db := dao.db.Where("version = 0 and province = ?", provinceName).Find(&pd)
	if db.Error != nil {
		log.Printf("GetProvinceData | get province data error | province=%s, err=%v", provinceName, db.Error)
		if db.Error.Error() == "record not found" {
			return nil, dberr.NewDBError("NotFound", "record not found")
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
			return nil, dberr.NewDBError("NotFound", "record not found")
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
	version := DateToVersion(pd.CreatedAt)
	db := dao.db.Model(pd).Update("version", version)
	if db.Error != nil {
		log.Printf("UpdateProvinceDataVersion | update province version error | err=%v \n", db.Error)
		return db.Error
	}
	return nil
}

func (dao *DataDao) UpdateCityDataVersion(cd models.CityDataModel) error {
	version := DateToVersion(cd.CreatedAt)
	db := dao.db.Model(cd).Update("version", version)
	if db.Error != nil {
		log.Printf("UpdateCityDataVersion | update cd version error | err=%v \n", db.Error)
		return db.Error
	}
	return nil
}

func (dao *DataDao) StoreProvinceData(data []models.ProvinceData) error {
	log.Println("start store province data")
	for _, p := range data {
		//time.Sleep(10 * time.Second)
		//fmt.Println(p)
		pd, err := dao.GetProvinceData(p.ProvinceShortName)
		if err != nil {
			if e, ok := err.(*dberr.DBError); ok && e.Code == "NotFound" {
				dao.CreateProvinceData(p)
				dao.BatchCreateCityData(p.ProvinceShortName, p.Cities)
			}
			continue
		}
		// update province data
		if pd.ConfirmedCount != p.ConfirmedCount || pd.SuspectedCount != p.SuspectedCount || pd.CuredCount != p.CuredCount || pd.DeadCount != p.DeadCount {
			dao.CreateProvinceData(p)
			dao.UpdateProvinceDataVersion(*pd)
		}
		dao.StoreCityData(p.ProvinceShortName, p.Cities)
	}
	return nil
}

func (dao *DataDao) StoreCityData(provinceName string, data []models.CityData) error {
	log.Printf("start store %s city data", provinceName)
	for _, c := range data {
		cd, err := dao.GetCityData(provinceName, c.CityName)
		if err != nil {
			if e, ok := err.(*dberr.DBError); ok && e.Code == "NotFound" {
				dao.CreateCityData(provinceName, c)
			}
			continue
		}
		// update province data
		if cd.ConfirmedCount != c.ConfirmedCount || cd.SuspectedCount != c.SuspectedCount || cd.CuredCount != c.CuredCount || cd.DeadCount != c.DeadCount {
			dao.CreateCityData(provinceName, c)
			dao.UpdateCityDataVersion(*cd)
		}
	}
	return nil
}

func (dao *DataDao) GetNewsData(title string) (*models.NewsModel, error) {
	nd := new(models.NewsModel)
	db := dao.db.Where("title = ?", title).Find(&nd)
	if db.Error != nil {
		log.Printf("GetNewsData | get news data error | title=%s, err=%v", title, db.Error)
		if db.Error.Error() == "record not found" {
			return nil, dberr.NewDBError("NotFound", "record not found")
		}
		return nil, db.Error
	}
	return nd, nil
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

func (dao *DataDao) StoreNewsData(data []models.NewsData) error {
	log.Println("start store news data")
	for _, n := range data {
		_, err := dao.GetNewsData(n.Title)
		if err != nil {
			if e, ok := err.(*dberr.DBError); ok && e.Code == "NotFound" {
				dao.CreateNewsData(n)
			}
			continue
		}
	}
	return nil
}
