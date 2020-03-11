package objectmock

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
)

type serviceMock struct {
	mock.Mock
}

func (s *serviceMock) Get(c *gin.Context) (string, error) {
	args := s.Called(c)
	return args.String(0), args.Error(1)
}

func (s *serviceMock) Post(c *gin.Context) (string, error) {
	args := s.Called(c)
	return args.String(0), args.Error(1)
}

func NewServiceMock() *serviceMock {
	return &serviceMock{}
}
