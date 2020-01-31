package tools

import (
	"fmt"
	"log"
	"time"

	"github.com/pkg/errors"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func dbConfig() DBConfig {
	dbUser := EnvGet("DBUSER", "ncov")
	dbPwd := EnvGet("DBPWD", "ncov")
	dbHost := EnvGet("DBHOST", "127.0.0.1")
	dbPort := EnvGet("DBPORT", "3306")
	dbName := EnvGet("DBNAME", "ncov")
	return DBConfig{
		DSN:          fmt.Sprintf("%s:%s@(%s:%s)/%s?timeout=1000ms&readTimeout=10000ms&charset=utf8mb4", dbUser, dbPwd, dbHost, dbPort, dbName),
		ConnMaxAge:   300,
		MaxIdleConns: 10,
		MaxOpenConns: 10,
	}
}

var DBConn *gorm.DB

func init() {
	dbConf := dbConfig()
	dbConn, err := GetDBConnect(dbConf)
	if err != nil {
		log.Fatal("conn db error")
	}
	DBConn = dbConn
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

type DBConfig struct {
	DSN          string
	ConnMaxAge   int
	MaxIdleConns int
	MaxOpenConns int
}
