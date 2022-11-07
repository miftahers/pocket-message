package repositories

import (
	"database/sql/driver"
	"errors"
	"pocket-message/models"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type GormSuite struct {
	suite.Suite
	mock sqlmock.Sqlmock
	repo Database
}

func TestSuiteGorm(t *testing.T) {
	suite.Run(t, new(GormSuite))
}

func (s *GormSuite) SetupSuite() {
	db, mock, err := sqlmock.New()
	if err != nil {
		s.Error(err)
	}

	gDB, err := gorm.Open(mysql.New(mysql.Config{
		SkipInitializeWithVersion: true,
		Conn:                      db,
	}), &gorm.Config{})
	if err != nil {
		s.Error(err)
	}

	repo := NewGorm(gDB)
	s.repo = repo
	s.mock = mock
}

func (s *GormSuite) TearDownSuite() {}

type AnyTime struct{}

func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}

// SaveNewUser
func (s *GormSuite) TestSaveNewUser() {
	testCase := []struct {
		name        string
		body        models.User
		expectError error
	}{
		{
			name: "save_new_user-normal",
			body: models.User{
				Username: "aku",
				Password: "akuGantenk",
			},
			expectError: nil,
		},
	}
	for _, v := range testCase {
		s.T().Run(v.name, func(t *testing.T) {
			s.mock.ExpectBegin()
			s.mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `users` (`created_at`,`updated_at`,`deleted_at`,`uuid`,`username`,`password`,`pocket_message`) VALUES (?,?,?,?,?,?,(NULL))")).
				WithArgs(AnyTime{}, AnyTime{}, nil, "00000000-0000-0000-0000-000000000000", "aku", "akuGantenk").
				WillReturnResult(sqlmock.NewResult(1, 1))
			s.mock.ExpectCommit()

			err := s.repo.SaveNewUser(v.body)
			s.Equal(v.expectError, err)
		})
	}
}
func (s *GormSuite) TestSaveNewUserError() {
	testCase := []struct {
		name        string
		body        models.User
		expectError error
	}{
		{
			name: "save_new_user-error",
			body: models.User{
				Username: "aku",
				Password: "akuGantenk",
			},
			expectError: errors.New("database error"),
		},
	}
	for _, v := range testCase {
		s.T().Run(v.name, func(t *testing.T) {
			s.mock.ExpectBegin()
			s.mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `users` (`created_at`,`updated_at`,`deleted_at`,`uuid`,`username`,`password`,`pocket_message`) VALUES (?,?,?,?,?,?,(NULL))")).
				WithArgs(AnyTime{}, AnyTime{}, nil, "00000000-0000-0000-0000-000000000000", "aku", "akuGantenk").
				WillReturnError(errors.New("database error"))
			s.mock.ExpectRollback()

			err := s.repo.SaveNewUser(v.body)
			s.Equal(v.expectError, err)
		})
	}
}
