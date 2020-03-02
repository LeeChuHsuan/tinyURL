package main
import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"fmt"
)


const connectionString = "host=localhost port=6666 user=postgres dbname=postgres password=root sslmode=disable"


type urlMapping struct{
	Url string 
	Hashval string 
}


func InserturlMapping(url, hashval string){
	db,err:= gorm.Open("postgres",connectionString)
	defer db.Close()
	if err!=nil{
		fmt.Println(err)
		return 
	}
	record := urlMapping{url,hashval}
	db.Create(&record)

}

func GeturlMapping(hashval string) string{
	db,err:= gorm.Open("postgres",connectionString)
	defer db.Close()
	if err!=nil{
		fmt.Println(err)
		return ""
	}
	var record urlMapping
	if err =db.Where("hashval=?",hashval).First(&record).Error;err!=nil{
		fmt.Println(err)
		return ""
	}
	return record.Url
}
