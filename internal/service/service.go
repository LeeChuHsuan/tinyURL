package service

import (
	"github.com/gin-gonic/gin"
)

const domainName = "localhost:8000"

type Service interface {
	Post(*gin.Context) (string, error)
	Get(*gin.Context) (string, error)
}
