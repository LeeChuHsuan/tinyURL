package service

import (
	"fmt"
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

type fileService struct {
	Service
}

type file struct{}

const maxFileSize = 1 << 20

func NewfileService() *fileService {
	return &fileService{&file{}}
}

func (f *file) Post(g *gin.Context) (string, error) {
	var fileStoreRoot = os.Getenv("uploadfileRoot")
	err := g.Request.ParseMultipartForm(maxFileSize)
	if err != nil {
		fmt.Print(err)
		return "", err
	}

	uploadFile, header, err := g.Request.FormFile("uploadFile")
	if err != nil {
		fmt.Print(err)
		return "", err
	}

	newFile, err := os.OpenFile(fileStoreRoot+header.Filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Print(err)
		return "", err
	}
	defer newFile.Close()

	_, err = io.Copy(newFile, uploadFile)
	if err != nil {
		fmt.Print(err)
		return "", err
	}
	return "", nil
}

func (f *file) Get(g *gin.Context) (string, error) {
	return "", nil
}
