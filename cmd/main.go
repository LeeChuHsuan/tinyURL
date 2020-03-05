package main

import (
	"tinyURL/internal/router"
)

func main() {
	router := router.SetupRouter()
	router.Run("localhost:8000")
}
