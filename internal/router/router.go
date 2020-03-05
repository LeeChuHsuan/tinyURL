package router

import (
	"tinyURL/internal/repository"
	"tinyURL/internal/service"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	repo := repository.NewURLMapping("", "")
	URLService := service.NewTinyURLService(
		service.NewtinyURL("", repo, nil),
	)
	router.GET("/", URLService.GetIndexPage)
	router.GET("/:hashval", URLService.GetHandler)
	router.POST("/", URLService.PostHandler)
	return router
}
