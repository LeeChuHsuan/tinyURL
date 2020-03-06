package repository

import (
	"errors"
	"fmt"

	"github.com/jinzhu/gorm"
)

type URLMappingRepo struct {
	dbConn *gorm.DB
}

func NewURLMappingRepo(dbConn *gorm.DB) *URLMappingRepo {
	return &URLMappingRepo{dbConn}
}

func (r *URLMappingRepo) InsertDB(record interface{}) error {
	record, ok := record.(*URLMapping)
	if !ok {
		return errors.New("type conversion error")
	}

	if err := r.dbConn.Create(record).Error; err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (r *URLMappingRepo) GetByPrimaryKey(key interface{}) (interface{}, error) {

	hashval, ok := key.(string)
	if !ok {
		return nil, errors.New("type conversion error")
	}

	var record URLMapping
	if err := r.dbConn.Where("hashval=?", hashval).First(&record).Error; err != nil {
		fmt.Println(err)
		return "", err
	}
	return record, nil
}

type URLMapping struct {
	URL     string
	Hashval string
}

//NewURLMapping create and return a new URLMapping
func NewURLMapping(url, hashval string) *URLMapping {
	return &URLMapping{url, hashval}
}
