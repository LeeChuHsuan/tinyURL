package controller

import (
	"fmt"
	"net/http"
	"tinyURL/internal/service"

	"github.com/gin-gonic/gin"
)

type tinyURLController struct {
	s service.Service
}

func NewtinyURLController(s service.Service) *tinyURLController {
	return &tinyURLController{s}
}

func (ctrl *tinyURLController) GetIndexPage(c *gin.Context) {
	c.Writer.WriteHeader(http.StatusOK)
	http.ServeFile(c.Writer, c.Request, "../web/index.html")
}

func (ctrl *tinyURLController) Get(c *gin.Context) {
	response, err := ctrl.s.Get(c)
	if err != nil {
		http.NotFound(c.Writer, c.Request)
		return
	}
	http.Redirect(c.Writer, c.Request, response, http.StatusFound)
}

func (ctrl *tinyURLController) Post(c *gin.Context) {
	response, err := ctrl.s.Post(c)
	if err != nil {
		c.Writer.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(c.Writer, "%s", err.Error())
		return
	}

	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusCreated)
	fmt.Fprintf(c.Writer, "%s", response)
}
