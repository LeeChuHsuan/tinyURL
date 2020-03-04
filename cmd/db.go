package main

import (
	"errors"
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/spf13/viper"
)

var connectionString string

func init() {
	viper.AddConfigPath(".")
	viper.SetConfigName("../config/dbconfig")
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}

	settings := []string{"host", "port", "user", "dbname", "password", "sslmode"}
	for _, s := range settings {
		connectionString = fmt.Sprintf("%s%s=%v ", connectionString, s, viper.Get(s))
	}
}

type urlMapping struct {
	URL     string
	Hashval string
}

type repository interface {
	InsertDB() error
	GetByPrimaryKey(string) (interface{}, error)
}

func (u *urlMapping) InsertDB() error {
	if u == nil {
		return errors.New("nil urlMapping pointer")
	}

	db, err := openDB()
	defer db.Close()
	if err != nil {
		fmt.Println(err)
		return err
	}

	if err = db.Create(u).Error; err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (u *urlMapping) GetByPrimaryKey(hashval string) (*urlMapping, error) {
	db, err := openDB()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	var record urlMapping
	if err = db.Where("hashval=?", hashval).First(&record).Error; err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &record, nil
}

func openDB() (*gorm.DB, error) {
	return gorm.Open("postgres", connectionString)
}

func newurlMapping(url, hashval string) urlMapping {
	return urlMapping{url, hashval}
}
