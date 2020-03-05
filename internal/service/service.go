package service

import (
	"github.com/gin-gonic/gin"
)

const domainName = "localhost:8000"

type Service interface {
	post(*gin.Context) (string, error)
	get(*gin.Context) (string, error)
}
