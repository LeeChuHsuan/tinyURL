package main

import (
	"tinyURL/internal/router"

	"github.com/gin-gonic/gin"
)

func setup() *gin.Engine {

	router := router.SetupRouter()
	return router
}

func main() {
	/*
		db, err := repository.OpenDB()
		defer db.Close()
		if err != nil {
			log.Fatal(err)
		}
	*/
	router := setup()
	router.Run("localhost:8000")
}
