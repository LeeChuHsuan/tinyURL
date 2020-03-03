package main
import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"fmt"
	"github.com/spf13/pflag"
)


const connectionString = "host=localhost port=6666 user=postgres dbname=postgres password=root sslmode=disable"


type urlMapping struct{
	URL string 
	Hashval string 
}


func InsertURLMapping(url, hashval string){
	db, err := gorm.Open("postgres", connectionString)
	defer db.Close()
	if err != nil{
		fmt.Println(err)
		return 
	}
	record := urlMapping{url, hashval}
	db.Create(&record)
}

func GetURLMapping(hashval string) (string, error){
	db, err := gorm.Open("postgres", connectionString)
	defer db.Close()
	if err!=nil{
		fmt.Println(err)
		return "", err
	}
	var record urlMapping
	if err = db.Where("hashval=?", hashval).First(&record).Error; err != nil {
		fmt.Println(err)
		return "", err 
	}
	return record.URL, nil
}
