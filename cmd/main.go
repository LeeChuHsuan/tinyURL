package main

import (
	"log"
	"os"
	"tinyURL/internal/repository"
	"tinyURL/internal/router"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func setup(dbConn *gorm.DB) *gin.Engine {
	os.Setenv("uploadfileRoot", "../uploadFiles/")
	router := router.SetupRouter(dbConn)
	return router
}

func main() {

	dbConn, err := repository.OpenDB()
	defer dbConn.Close()
	if err != nil {
		log.Fatal(err)
	}

	router := setup(dbConn)
	router.Run("localhost:8000")
}
