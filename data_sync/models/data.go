package models

/*
CREATE TABLE `country_data` (
    `id` bigint UNSIGNED AUTO_INCREMENT NOT NULL PRIMARY KEY,
    `country` varchar(64) NOT NULL,
    `confirmed_count` INT NOT NULL,
    `suspected_count` INT NOT NULL,
    `cured_count` INT NOT NULL,
    `dead_count` INT NOT NULL,
    `version` varchar(64) NOT NULL,
    `created_at` bigint NOT NULL
);
*/
type CountryDataModel struct {
	Id             int    `gorm:"column:id" json:"id"`
	CountryName    string `gorm:"column:country" json:"country"`
	ConfirmedCount int    `gorm:"column:confirmed_count" json:"confirmed_count"`
	SuspectedCount int    `gorm:"column:suspected_count" json:"suspected_count"`
	CuredCount     int    `gorm:"column:cured_count" json:"cured_count"`
	DeadCount      int    `gorm:"column:dead_count" json:"dead_count"`
	Version        string `gorm:"column:version" json:"version"`
	CreatedAt      int64  `gorm:"column:created_at" json:"created_at"`
}

func (rs *CountryDataModel) TableName() string {
	return "country_data"
}

/*
CREATE TABLE `city_data` (
    `id` bigint UNSIGNED AUTO_INCREMENT NOT NULL PRIMARY KEY,
    `country` varchar(64) NOT NULL,
    `province` varchar(64) NOT NULL,
    `city` varchar(64) NOT NULL,
    `confirmed_count` INT NOT NULL,
    `suspected_count` INT NOT NULL,
    `cured_count` INT NOT NULL,
    `dead_count` INT NOT NULL,
    `version` varchar(64) NOT NULL,
    `created_at` bigint NOT NULL
);
*/
type CityDataModel struct {
	Id             int    `gorm:"column:id" json:"id"`
	Province       string `gorm:"column:province" json:"province"`
	City           string `gorm:"column:city" json:"city"`
	ConfirmedCount int    `gorm:"column:confirmed_count" json:"confirmed_count"`
	SuspectedCount int    `gorm:"column:suspected_count" json:"suspected_count"`
	CuredCount     int    `gorm:"column:cured_count" json:"cured_count"`
	DeadCount      int    `gorm:"column:dead_count" json:"dead_count"`
	Version        string `gorm:"column:version" json:"version"`
	CreatedAt      int64  `gorm:"column:created_at" json:"created_at"`
}

func (rs *CityDataModel) TableName() string {
	return "city_data"
}

/*
CREATE TABLE `province_data` (
    `id` bigint UNSIGNED AUTO_INCREMENT NOT NULL PRIMARY KEY,
    `country` varchar(64) NOT NULL,
    `province` varchar(64) NOT NULL,
    `confirmed_count` INT NOT NULL,
    `suspected_count` INT NOT NULL,
    `cured_count` INT NOT NULL,
    `dead_count` INT NOT NULL,
    `version` varchar(64) NOT NULL,
    `created_at` bigint NOT NULL
);
*/
type ProvinceDataModel struct {
	Id             int    `gorm:"column:id" json:"id"`
	Country        string `gorm:"column:country" json:"country"`
	Province       string `gorm:"column:province" json:"province"`
	ConfirmedCount int    `gorm:"column:confirmed_count" json:"confirmed_count"`
	SuspectedCount int    `gorm:"column:suspected_count" json:"suspected_count"`
	CuredCount     int    `gorm:"column:cured_count" json:"cured_count"`
	DeadCount      int    `gorm:"column:dead_count" json:"dead_count"`
	Version        string `gorm:"column:version" json:"version"`
	CreatedAt      int64  `gorm:"column:created_at" json:"created_at"`
}

func (rs *ProvinceDataModel) TableName() string {
	return "province_data"
}

/*
CREATE TABLE `news` (
    `id` bigint UNSIGNED AUTO_INCREMENT NOT NULL PRIMARY KEY,
    `province` varchar(64) NOT NULL,
    `title` varchar(256) NOT NULL,
    `summary` varchar(1024) NOT NULL,
    `source` varchar(64) NOT NULL,
    `source_url` varchar(256) NOT NULL,
    `published_at` bigint NOT NULL,
    `created_at` bigint NOT NULL
);
*/
type NewsModel struct {
	Id          int    `gorm:"column:id" json:"id"`
	Province    string `gorm:"column:province" json:"province"`
	Title       string `gorm:"column:title" json:"title"`
	Summary     string `gorm:"column:summary" json:"summary"`
	Source      string `gorm:"column:source" json:"source"`
	SourceUrl   string `gorm:"column:source_url" json:"source_url"`
	PublishedAt int64  `gorm:"column:published_at" json:"published_at"`
	CreatedAt   int64  `gorm:"column:created_at" json:"created_at"`
}

func (rs *NewsModel) TableName() string {
	return "news"
}
