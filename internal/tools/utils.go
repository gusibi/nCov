package tools

import (
	"os"
	"time"
)

func EnvGet(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func ParseBJTime(layout string, value string) (time.Time, error) {
	BeijingTimeZone, _ := time.LoadLocation("Asia/Shanghai")
	return time.ParseInLocation(layout, value, BeijingTimeZone)
}

func TimeStampToString(ts int64, layout string) string {
	date := time.Unix(ts, 0)
	return date.Format(layout)
}

func BjTimeStampToBJString(ts int64, layout string) string {
	// ts 为北京 时区时间戳，需要判断当前时区
	date := time.Unix(ts, 0)
	_, offset := time.Now().Zone()
	offset = 3600*8 - offset
	date = date.Add(time.Second * time.Duration(offset))
	dateStr := date.Format(layout)
	date, err := ParseBJTime(layout, dateStr)
	if err != nil {
		return ""
	}
	return date.Format(layout)
}

func DateToVersion(ts int64) string {
	return TimeStampToString(ts, "2006-01-02T15:04:05")
}
