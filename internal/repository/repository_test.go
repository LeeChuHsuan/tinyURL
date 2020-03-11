package repository_test

import (
	"errors"
	"testing"
	"tinyURL/internal/repository"

	mocket "github.com/Selvatico/go-mocket"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type RepoSuite struct {
	suite.Suite
	DB   *gorm.DB
	repo repository.Repository
}

func (s *RepoSuite) SetupSuite() {
	mocket.Catcher.Register()
	db, err := gorm.Open(mocket.DriverName, "any_string")
	require.NoError(s.T(), err)
	s.DB = db
	db.LogMode(true)
	s.repo = repository.NewURLMappingRepo(s.DB)
}

func TestInit(t *testing.T) {
	suite.Run(t, new(RepoSuite))
}

func (s *RepoSuite) TestRepositoryInsertDB() {
	testcase := []struct {
		Data interface{}
		Err  error
	}{
		{
			repository.NewURLMapping("url1", "hashval1"),
			nil,
		}, {
			"",
			errors.New("type conversion error"),
		},
	}

	for _, test := range testcase {

		if test.Err == nil {
			mocket.Catcher.Reset().NewMock().WithQuery("^INSERT INTO \"url_mappings\"")
			err := s.repo.InsertDB(test.Data)
			require.NoError(s.T(), err)
		} else {
			err := s.repo.InsertDB(test.Data)
			require.EqualError(s.T(), err, test.Err.Error())
		}
	}
}

func (s *RepoSuite) TestRepositoryGetByPrimaryKey() {
	testcase := []struct {
		Key interface{}
		Val interface{}
		Err error
	}{
		{
			"hash1",
			"url",
			nil,
		}, {
			1,
			"",
			errors.New("type conversion error"),
		},
	}

	for _, test := range testcase {
		if test.Err == nil {
			commonReply := []map[string]interface{}{{"hashval": test.Key, "url": test.Val}}
			mocket.Catcher.Reset().NewMock().WithReply(commonReply)
			_, err := s.repo.GetByPrimaryKey(test.Key)
			require.NoError(s.T(), err)
		} else {
			_, err := s.repo.GetByPrimaryKey(test.Key)
			require.EqualError(s.T(), err, test.Err.Error())
		}
	}
}
