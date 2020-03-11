package controller

import (
	"fmt"
	"net/http"
	"tinyURL/internal/service"

	"github.com/gin-gonic/gin"
)

type fileServiceController struct {
	s service.Service
}

func NewfileServiceController(s service.Service) *fileServiceController {
	return &fileServiceController{s}
}

func (ctrl *fileServiceController) GetIndexPage(c *gin.Context) {
	c.Writer.WriteHeader(http.StatusBadRequest)
}

func (ctrl *fileServiceController) Get(c *gin.Context) {
	c.Writer.WriteHeader(http.StatusBadRequest)
}

func (ctrl *fileServiceController) Post(c *gin.Context) {
	response, err := ctrl.s.Post(c)
	if err != nil {
		c.Writer.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(c.Writer, "%s", err.Error())
		return
	}
	c.Writer.WriteHeader(http.StatusCreated)
	fmt.Fprintf(c.Writer, "%s", response)
	return
}
