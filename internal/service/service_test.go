package service_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http/httptest"
	"os"
	"strconv"
	"strings"
	"testing"
	"tinyURL/internal/objectmock"
	"tinyURL/internal/repository"
	"tinyURL/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestTinyURLServicePost(t *testing.T) {

	testCase := []struct {
		URL    string
		Result string
		Err    error
	}{
		{
			URL:    "test.com",
			Result: "",
			Err:    nil,
		}, {
			URL:    "test2.com",
			Result: "",
			Err:    errors.New("Insert Error"),
		},
	}

	repo := objectmock.NewRepoMock()
	URLService := service.NewTinyURLService(
		service.NewtinyURL("", repo, nil),
	)

	gin.SetMode(gin.TestMode)

	for _, test := range testCase {

		URLhash := service.HashURL(test.URL)
		record := repository.NewURLMapping(test.URL, URLhash)
		if test.Err == nil {
			res, _ := json.Marshal(*(service.NewtinyURL(URLhash, nil, nil)))
			test.Result = string(res)
		}

		repo.On("InsertDB", record).Return(test.Err)

		postData := strings.NewReader("url=" + test.URL)
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		req := httptest.NewRequest("POST", "/url", postData)
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		c.Request = req

		result, err := URLService.Post(c)
		if test.Err != nil {
			assert.EqualError(t, err, test.Err.Error())
		} else {
			assert.Nil(t, err)
		}
		assert.Equal(t, result, test.Result)
	}

}

func TestTinyURLServiceGet(t *testing.T) {

	repo := objectmock.NewRepoMock()
	URLService := service.NewTinyURLService(
		service.NewtinyURL("", repo, nil),
	)

	gin.SetMode(gin.TestMode)

	urlmapping := repository.NewURLMapping("url", "hashval")
	testCase := []struct {
		Record interface{}
		Result string
		Err    error
	}{
		{
			*urlmapping,
			"url",
			nil,
		}, {
			*urlmapping,
			"",
			errors.New("Test Error"),
		},
	}

	for i, test := range testCase {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		req := httptest.NewRequest("GET", "/url/", nil)

		c.Request = req
		hashval := strconv.Itoa(i)
		c.Params = []gin.Param{{"hashval", hashval}}

		repo.On("GetByPrimaryKey", hashval).Return(test.Record, test.Err)
		result, err := URLService.Get(c)

		if test.Err != nil {
			assert.EqualError(t, err, test.Err.Error())
		} else {
			assert.Nil(t, err)
		}
		assert.Equal(t, result, test.Result)
	}
}

func TestFileServicePost(t *testing.T) {

	body_buf := bytes.NewBufferString("")
	body_writer := multipart.NewWriter(body_buf)

	_, err := body_writer.CreateFormFile("uploadFile", "./mytest.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fh, err := os.Open("./mytest.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fi, err := fh.Stat()
	if err != nil {
		fmt.Println(err)
		return
	}

	boundary := body_writer.Boundary()
	close_buf := bytes.NewBufferString(fmt.Sprintf("\r\n--%s--\r\n", boundary))
	request_reader := io.MultiReader(body_buf, fh, close_buf)

	req := httptest.NewRequest("POST", "/file", request_reader)
	req.Header.Add("Content-Type", "multipart/form-data; boundary="+boundary)
	req.ContentLength = fi.Size() + int64(body_buf.Len()) + int64(close_buf.Len())

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	os.Setenv("uploadfileRoot", "../../uploadFiles/")
	c.Request = req
	uploadFileService := service.NewfileService()

	_, err = uploadFileService.Post(c)
	assert.Nil(t, err)
}
