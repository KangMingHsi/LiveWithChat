package postgres

import (
	"authentication"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type Suite struct {
	suite.Suite
	DB   *gorm.DB
	mock sqlmock.Sqlmock
 
	repository authentication.UserRepository
	user     *User
}

// func (s *Suite) SetupSuite() {
// 	var (
// 	   db  *sql.DB
// 	   err error
// 	)
 
// 	db, s.mock, err = sqlmock.New()
// 	require.NoError(s.T(), err)

// 	s.DB, err = gorm.Open(db, )
// 	require.NoError(s.T(), err)
 
// 	s.repository = NewUserRepository(s.DB)
//  }

 
func TestStore(t *testing.T) {

}

func TestFind(t *testing.T) {

}

func TestFindAll(t *testing.T) {

}