package objectmock

import "github.com/stretchr/testify/mock"

type repoMock struct {
	mock.Mock
}

func (r *repoMock) InsertDB(record interface{}) error {
	args := r.Called(record)
	return args.Error(0)
}

func (r *repoMock) GetByPrimaryKey(key interface{}) (interface{}, error) {
	args := r.Called(key)
	return args.Get(0), args.Error(1)
}

func NewRepoMock() *repoMock {
	return &repoMock{}
}
