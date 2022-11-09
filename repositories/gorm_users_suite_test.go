package repositories

import (
	"database/sql/driver"
	"errors"
	"pocket-message/models"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type GormUsersSuite struct {
	suite.Suite
	mock sqlmock.Sqlmock
	repo Database
}

func TestSuiteGormUsers(t *testing.T) {
	suite.Run(t, new(GormUsersSuite))
}

func (s *GormUsersSuite) SetupSuite() {
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

func (s *GormUsersSuite) TearDownSuite() {}

type AnyTime struct{}

// Match satisfies sqlmock.Argument interface
func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}

// SaveNewUser
func (s *GormUsersSuite) TestSaveNewUser() {
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
func (s *GormUsersSuite) TestSaveNewUserError() {
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

// Login
func (s *GormUsersSuite) TestLogin() {
	testCase := []struct {
		name        string
		body        models.User
		expectBody  models.User
		expectError error
	}{
		{
			name: "login-normal",
			body: models.User{
				Username: "userTest",
				Password: "passwordTest",
			},
			expectBody: models.User{
				UUID:     uuid.Nil,
				Username: "userTest",
				Password: "passwordTest",
			},
			expectError: nil,
		},
	}
	for _, v := range testCase {
		s.T().Run(v.name, func(t *testing.T) {
			expectRow := s.mock.NewRows([]string{"uuid", "username", "password"}).
				AddRow("00000000-0000-0000-0000-000000000000", "userTest", "passwordTest")

			s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users` WHERE (username = ? AND password = ?) AND `users`.`deleted_at` IS NULL ORDER BY `users`.`id` LIMIT 1")).
				WithArgs("userTest", "passwordTest").
				WillReturnRows(expectRow)

			result, err := s.repo.Login(v.body)
			s.Equal(v.expectError, err)
			s.Equal(v.expectBody, result)
		})
	}
}
func (s *GormUsersSuite) TestLoginError() {
	testCase := []struct {
		name        string
		body        models.User
		expectBody  models.User
		expectError error
	}{
		{
			name: "login-error",
			body: models.User{
				Username: "userTest",
				Password: "passwordTest",
			},
			expectBody:  models.User{},
			expectError: errors.New("record not found"),
		},
	}
	for _, v := range testCase {
		s.T().Run(v.name, func(t *testing.T) {
			s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users` WHERE (username = ? AND password = ?) AND `users`.`deleted_at` IS NULL ORDER BY `users`.`id` LIMIT 1")).
				WithArgs("userTest", "passwordTest").
				WillReturnError(errors.New("record not found"))

			result, err := s.repo.Login(v.body)
			s.Equal(v.expectError, err)
			s.Equal(v.expectBody, result)
		})
	}
}

// Update Username
func (s *GormUsersSuite) TestUpdateUsername() {
	testCase := []struct {
		name        string
		body        models.User
		expectError error
	}{
		{
			name: "update_username-normal",
			body: models.User{
				UUID:     uuid.Nil,
				Username: "userTest69",
				Password: "passwordTest",
			},
			expectError: nil,
		},
	}
	for _, v := range testCase {
		s.T().Run(v.name, func(t *testing.T) {
			s.mock.ExpectBegin()
			s.mock.ExpectExec(regexp.QuoteMeta("UPDATE `users` SET `username`=?,`updated_at`=? WHERE uuid = ? AND `users`.`deleted_at` IS NULL")).
				WithArgs("userTest69", AnyTime{}, uuid.Nil).
				WillReturnResult(sqlmock.NewResult(1, 1))
			s.mock.ExpectCommit()

			err := s.repo.UpdateUsername(v.body)
			s.Equal(v.expectError, err)
		})
	}
}
func (s *GormUsersSuite) TestUpdateUsernameError() {
	testCase := []struct {
		name        string
		body        models.User
		expectError error
	}{
		{
			name: "update_username-error",
			body: models.User{
				UUID:     uuid.Nil,
				Username: "userTest69",
				Password: "passwordTest",
			},
			expectError: errors.New("record not found"),
		},
	}
	for _, v := range testCase {
		s.T().Run(v.name, func(t *testing.T) {
			s.mock.ExpectBegin()
			s.mock.ExpectExec(regexp.QuoteMeta("UPDATE `users` SET `username`=?,`updated_at`=? WHERE uuid = ? AND `users`.`deleted_at` IS NULL")).
				WithArgs("userTest69", AnyTime{}, uuid.Nil).
				WillReturnError(errors.New("record not found"))
			s.mock.ExpectRollback()

			err := s.repo.UpdateUsername(v.body)
			s.Equal(v.expectError, err)
		})
	}
}

// Update Password
func (s *GormUsersSuite) TestUpdatePassword() {
	testCase := []struct {
		name        string
		body        models.User
		expectError error
	}{
		{
			name: "update_password-normal",
			body: models.User{
				Username: "userTest69",
				Password: "passwordTest79",
			},
			expectError: nil,
		},
	}
	for _, v := range testCase {
		s.T().Run(v.name, func(t *testing.T) {
			s.mock.ExpectBegin()
			s.mock.ExpectExec(regexp.QuoteMeta("UPDATE `users` SET `password`=?,`updated_at`=? WHERE username = ? AND `users`.`deleted_at` IS NULL")).
				WithArgs("passwordTest79", AnyTime{}, "userTest69").
				WillReturnResult(sqlmock.NewResult(1, 1))
			s.mock.ExpectCommit()

			err := s.repo.UpdatePassword(v.body)
			s.Equal(v.expectError, err)
		})
	}
}
func (s *GormUsersSuite) TestUpdatePasswordError() {
	testCase := []struct {
		name        string
		body        models.User
		expectError error
	}{
		{
			name: "update_password-error",
			body: models.User{
				Username: "userTest69",
				Password: "passwordTest79",
			},
			expectError: errors.New("record not found"),
		},
	}
	for _, v := range testCase {
		s.T().Run(v.name, func(t *testing.T) {
			s.mock.ExpectBegin()
			s.mock.ExpectExec(regexp.QuoteMeta("UPDATE `users` SET `password`=?,`updated_at`=? WHERE username = ? AND `users`.`deleted_at` IS NULL")).
				WithArgs("passwordTest79", AnyTime{}, "userTest69").
				WillReturnError(errors.New("record not found"))
			s.mock.ExpectRollback()

			err := s.repo.UpdatePassword(v.body)
			s.Equal(v.expectError, err)
		})
	}
}
