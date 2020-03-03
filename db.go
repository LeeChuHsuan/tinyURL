package main
import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"fmt"
	"github.com/spf13/viper"
	"log"
)


var connectionString string 

func init(){
	viper.SetConfigName("dbconfig")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}

	settings := []string{"host", "port", "user", "dbname", "password", "sslmode"}
	for _, s := range settings {
		connectionString = fmt.Sprintf("%s%s=%v ", connectionString, s, viper.Get(s))
	}
	fmt.Println(connectionString)
}

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
