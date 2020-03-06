package controller

import (
	"github.com/gin-gonic/gin"
)

type Controller interface {
	Get(*gin.Context)
	Post(*gin.Context)
	GetIndexPage(*gin.Context)
}
