package controller_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"tinyURL/internal/controller"
	"tinyURL/internal/objectmock"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestTinyURLControllerPost(t *testing.T) {

	testCase := []struct {
		Err        error
		StatusCode int
	}{
		{
			Err:        nil,
			StatusCode: http.StatusCreated,
		}, {
			Err:        errors.New("Insert Error"),
			StatusCode: http.StatusBadRequest,
		},
	}

	URLService := objectmock.NewServiceMock()
	tinyURLController := controller.NewtinyURLController(URLService)

	gin.SetMode(gin.TestMode)

	for _, test := range testCase {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = httptest.NewRequest("POST", "/url", nil)
		URLService.On("Post", c).Return("", test.Err)
		tinyURLController.Post(c)

		assert.Equal(t, test.StatusCode, w.Code)
	}
}

func TestTinyURLControllerGet(t *testing.T) {

	testCase := []struct {
		Err        error
		StatusCode int
	}{
		{
			nil,
			http.StatusFound,
		}, {
			errors.New("Test Error"),
			http.StatusNotFound,
		},
	}

	URLService := objectmock.NewServiceMock()
	tinyURLController := controller.NewtinyURLController(URLService)

	gin.SetMode(gin.TestMode)

	for _, test := range testCase {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		req := httptest.NewRequest("GET", "/url/", nil)
		c.Request = req

		URLService.On("Get", c).Return("", test.Err)
		tinyURLController.Get(c)

		assert.Equal(t, w.Code, test.StatusCode)
	}

}

func TestFileServiceControllerPost(t *testing.T) {

	testCase := []struct {
		StatusCode int
		Err        error
	}{
		{
			http.StatusBadRequest,
			errors.New("file upload error"),
		}, {
			http.StatusCreated,
			nil,
		},
	}

	uploadFileService := objectmock.NewServiceMock()
	uploadFileController := controller.NewfileServiceController(uploadFileService)

	for _, test := range testCase {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = httptest.NewRequest("POST", "/file/", nil)

		uploadFileService.On("Post", c).Return("", test.Err)
		uploadFileController.Post(c)
		assert.Equal(t, w.Code, test.StatusCode)
	}

}
