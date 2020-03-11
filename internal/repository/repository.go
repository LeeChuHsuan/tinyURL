package repository

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/spf13/viper"
)

type Repository interface {
	InsertDB(interface{}) error
	GetByPrimaryKey(interface{}) (interface{}, error)
}

func init() {
	viper.AddConfigPath(".")
	viper.AddConfigPath("..")
	viper.SetConfigName("../config/dbconfig")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	var connectionString string
	settings := []string{"host", "port", "user", "dbname", "password", "sslmode"}
	for _, s := range settings {
		connectionString = fmt.Sprintf("%s%s=%v ", connectionString, s, viper.Get(s))
	}
	os.Setenv("dbConnectionString", connectionString)
}

func OpenDB() (*gorm.DB, error) {
	return gorm.Open("postgres", os.Getenv("dbConnectionString"))
}
