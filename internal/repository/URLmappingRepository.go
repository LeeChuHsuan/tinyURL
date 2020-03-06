package repository

import (
	"errors"
	"fmt"
)

type URLMapping struct {
	URL     string
	Hashval string
}

//NewURLMapping create and return a new URLMapping
func NewURLMapping(url, hashval string) *URLMapping {
	return &URLMapping{url, hashval}
}

func (u *URLMapping) New(m map[string]string) Repository {
	return NewURLMapping(m["URL"], m["Hashval"])
}

func (u *URLMapping) InsertDB() error {
	if u == nil {
		return errors.New("nil URLMapping pointer")
	}

	db, err := OpenDB()
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

func (u *URLMapping) GetByPrimaryKey(hashval string) (string, error) {
	db, err := OpenDB()
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	var record URLMapping
	if err = db.Where("hashval=?", hashval).First(&record).Error; err != nil {
		fmt.Println(err)
		return "", err
	}
	return record.URL, nil
}
