package models

type CountryData struct {
	CountryName    string `json:"provinceName"`
	ConfirmedCount int    `json:"confirmedCount"`
	SuspectedCount int    `json:"suspectedCount"`
	CuredCount     int    `json:"curedCount"`
	DeadCount      int    `json:"deadCount"`
}

type CityData struct {
	CityName       string `json:"cityName"`
	ConfirmedCount int    `json:"confirmedCount"`
	SuspectedCount int    `json:"suspectedCount"`
	CuredCount     int    `json:"curedCount"`
	DeadCount      int    `json:"deadCount"`
}

type ProvinceData struct {
	ProvinceName      string     `json:"provinceName"`
	ProvinceShortName string     `json:"provinceShortName"`
	ConfirmedCount    int        `json:"confirmedCount"`
	SuspectedCount    int        `json:"suspectedCount"`
	CuredCount        int        `json:"curedCount"`
	DeadCount         int        `json:"deadCount"`
	Cities            []CityData `json:"cities"`
}

type NewsData struct {
	Id            int    `json:"id"`
	ProvinceName  string `json:"provinceName"`
	Title         string `json:"title"`
	Summary       string `json:"summary"`
	Source        string `json:"infoSource"`
	SourceUrl     string `json:"sourceUrl"`
	PublishedDate int64  `json:"pubDate"`
	CreatedTime   int64  `json:"createTime"`
}
