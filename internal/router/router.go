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

	uploadFileService := service.NewfileService()
	uploadFileController := controller.NewfileServiceController(uploadFileService)

	r1 := router.Group("/url")
	r1.GET("/", tinyURLController.GetIndexPage)
	r1.GET("/:hashval", tinyURLController.Get)
	r1.POST("/", tinyURLController.Post)

	r2 := router.Group("/file")
	r2.POST("/", uploadFileController.Post)

	return router
}
