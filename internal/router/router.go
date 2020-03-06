package router

import (
	"tinyURL/internal/controller"
	"tinyURL/internal/repository"
	"tinyURL/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func SetupRouter(dbConn *gorm.DB) *gin.Engine {
	router := gin.Default()
	repo := repository.NewURLMappingRepo(dbConn)
	URLService := service.NewTinyURLService(
		service.NewtinyURL("", repo, nil),
	)

	tinyURLController := controller.NewtinyURLController(URLService)

	router.GET("/", tinyURLController.GetIndexPage)
	router.GET("/:hashval", tinyURLController.Get)
	router.POST("/", tinyURLController.Post)
	return router
}
