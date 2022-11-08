package repositories

import (
	"database/sql/driver"
	"errors"
	"pocket-message/dto"
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

// Login
func (s *GormSuite) TestLogin() {
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
func (s *GormSuite) TestLoginError() {
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
func (s *GormSuite) TestUpdateUsername() {
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
func (s *GormSuite) TestUpdateUsernameError() {
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
func (s *GormSuite) TestUpdatePassword() {
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
func (s *GormSuite) TestUpdatePasswordError() {
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

// NewPocketMessage
func (s *GormSuite) TestNewPocketMessage() {
	testCase := []struct {
		name        string
		body        models.PocketMessage
		expectError error
	}{
		{
			name: "new_pocket_message-normal",
			body: models.PocketMessage{
				UUID:     uuid.Nil,
				Title:    "testJudul",
				Content:  "testContent",
				UserUUID: uuid.Nil,
			},
			expectError: nil,
		},
	}
	for _, v := range testCase {
		s.T().Run(v.name, func(t *testing.T) {
			s.mock.ExpectBegin()
			s.mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `pocket_messages` (`created_at`,`updated_at`,`deleted_at`,`uuid`,`title`,`content`,`user_uuid`) VALUES (?,?,?,?,?,?,?)")).
				WithArgs(AnyTime{}, AnyTime{}, nil, "00000000-0000-0000-0000-000000000000", "testJudul", "testContent", "00000000-0000-0000-0000-000000000000").
				WillReturnResult(sqlmock.NewResult(1, 1))
			s.mock.ExpectCommit()

			err := s.repo.SaveNewPocketMessage(v.body)
			s.Equal(v.expectError, err)
		})
	}
}
func (s *GormSuite) TestNewPocketMessageError() {
	testCase := []struct {
		name        string
		body        models.PocketMessage
		expectError error
	}{
		{
			name: "new_pocket_message-error",
			body: models.PocketMessage{
				UUID:     uuid.Nil,
				Title:    "testJudul",
				Content:  "testContent",
				UserUUID: uuid.Nil,
			},
			expectError: errors.New("database error"),
		},
	}
	for _, v := range testCase {
		s.T().Run(v.name, func(t *testing.T) {
			s.mock.ExpectBegin()
			s.mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `pocket_messages` (`created_at`,`updated_at`,`deleted_at`,`uuid`,`title`,`content`,`user_uuid`) VALUES (?,?,?,?,?,?,?)")).
				WithArgs(AnyTime{}, AnyTime{}, nil, "00000000-0000-0000-0000-000000000000", "testJudul", "testContent", "00000000-0000-0000-0000-000000000000").
				WillReturnError(errors.New("database error"))
			s.mock.ExpectRollback()

			err := s.repo.SaveNewPocketMessage(v.body)
			s.Equal(v.expectError, err)
		})
	}
}

// Save NewRandomID
func (s *GormSuite) TestSaveNewRandomID() {
	testCase := []struct {
		name        string
		body        models.PocketMessageRandomID
		expectError error
	}{
		{
			name: "new_pocket_message_random_id-normal",
			body: models.PocketMessageRandomID{
				RandomID:          "asdfghjkl",
				Visit:             0,
				PocketMessageUUID: uuid.Nil,
			},
			expectError: nil,
		},
	}
	for _, v := range testCase {
		s.T().Run(v.name, func(t *testing.T) {
			s.mock.ExpectBegin()
			s.mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `pocket_message_random_id` (`created_at`,`updated_at`,`deleted_at`,`random_id`,`visit`,`pocket_message_uuid`) VALUES (?,?,?,?,?,?)")).
				WithArgs(AnyTime{}, AnyTime{}, nil, "asdfghjkl", 0, "00000000-0000-0000-0000-000000000000").
				WillReturnResult(sqlmock.NewResult(1, 1))
			s.mock.ExpectCommit()

			err := s.repo.SaveNewRandomID(v.body)
			s.Equal(v.expectError, err)
		})
	}
}
func (s *GormSuite) TestSaveNewRandomIDError() {
	testCase := []struct {
		name        string
		body        models.PocketMessageRandomID
		expectError error
	}{
		{
			name: "new_pocket_message_random_id-error",
			body: models.PocketMessageRandomID{
				RandomID:          "asdfghjkl",
				Visit:             0,
				PocketMessageUUID: uuid.Nil,
			},
			expectError: errors.New("database error"),
		},
	}
	for _, v := range testCase {
		s.T().Run(v.name, func(t *testing.T) {
			s.mock.ExpectBegin()
			s.mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `pocket_message_random_id` (`created_at`,`updated_at`,`deleted_at`,`random_id`,`visit`,`pocket_message_uuid`) VALUES (?,?,?,?,?,?)")).
				WithArgs(AnyTime{}, AnyTime{}, nil, "asdfghjkl", 0, "00000000-0000-0000-0000-000000000000").
				WillReturnError(errors.New("database error"))
			s.mock.ExpectRollback()

			err := s.repo.SaveNewRandomID(v.body)
			s.Equal(v.expectError, err)
		})
	}
}

// GetPocketMessageByRandomID
func (s *GormSuite) TestGetPocketMessageByRandomID() {
	testCase := []struct {
		name        string
		randomID    string
		expectBody  dto.PocketMessageWithRandomID
		expectError error
	}{
		{
			name:     "get_pocket_message_by_random_id-normal",
			randomID: "asdfghjkl",
			expectBody: dto.PocketMessageWithRandomID{
				Title:    "superman mencari jodoh",
				Content:  "tapi boong",
				Visit:    0,
				RandomID: "asdfghjkl",
			},
			expectError: nil,
		},
	}
	for _, v := range testCase {
		s.T().Run(v.name, func(t *testing.T) {
			expectRow := s.mock.NewRows([]string{"title", "content", "visit", "random_id"}).
				AddRow("superman mencari jodoh", "tapi boong", 0, "asdfghjkl")

			s.mock.ExpectQuery(regexp.QuoteMeta("SELECT pocket_messages.UUID, pocket_messages.title, pocket_messages.content,pocket_message_random_id.visit, pocket_message_random_id.random_id FROM `pocket_messages` LEFT JOIN pocket_message_random_id ON pocket_messages.uuid = pocket_message_random_id.pocket_message_uuid WHERE pocket_message_random_id.random_id = ? AND `pocket_messages`.`deleted_at` IS NULL ORDER BY `pocket_messages`.`id` LIMIT 1")).
				WithArgs("asdfghjkl").
				WillReturnRows(expectRow)

			result, err := s.repo.GetPocketMessageByRandomID(v.randomID)
			s.Equal(v.expectError, err)
			s.Equal(v.expectBody, result)
		})
	}
}
func (s *GormSuite) TestGetPocketMessageByRandomIDError() {
	testCase := []struct {
		name        string
		randomID    string
		expectBody  dto.PocketMessageWithRandomID
		expectError error
	}{
		{
			name:        "get_pocket_message_by_random_id-error",
			randomID:    "asdfghjkl",
			expectBody:  dto.PocketMessageWithRandomID{},
			expectError: errors.New("record not found"),
		},
	}
	for _, v := range testCase {
		s.T().Run(v.name, func(t *testing.T) {
			s.mock.ExpectQuery(regexp.QuoteMeta("SELECT pocket_messages.UUID, pocket_messages.title, pocket_messages.content,pocket_message_random_id.visit, pocket_message_random_id.random_id FROM `pocket_messages` LEFT JOIN pocket_message_random_id ON pocket_messages.uuid = pocket_message_random_id.pocket_message_uuid WHERE pocket_message_random_id.random_id = ? AND `pocket_messages`.`deleted_at` IS NULL ORDER BY `pocket_messages`.`id` LIMIT 1")).
				WithArgs("asdfghjkl").
				WillReturnError(errors.New("record not found"))

			result, err := s.repo.GetPocketMessageByRandomID(v.randomID)
			s.Equal(v.expectError, err)
			s.Equal(v.expectBody, result)
		})
	}
}

// Update VisitCount
func (s *GormSuite) TestUpdateVisitCount() {
	testCase := []struct {
		name        string
		body        dto.PocketMessageWithRandomID
		expectError error
	}{
		{
			name: "update_visit_count-normal",
			body: dto.PocketMessageWithRandomID{
				UUID:  uuid.Nil,
				Visit: 0,
			},
			expectError: nil,
		},
	}
	for _, v := range testCase {
		s.T().Run(v.name, func(t *testing.T) {
			s.mock.ExpectBegin()
			s.mock.ExpectExec(regexp.QuoteMeta("UPDATE `pocket_message_random_id` SET `visit`=?,`updated_at`=? WHERE pocket_message_uuid = ? AND `pocket_message_random_id`.`deleted_at` IS NULL")).
				WithArgs(v.body.Visit+1, AnyTime{}, uuid.Nil).
				WillReturnResult(sqlmock.NewResult(1, 1))
			s.mock.ExpectCommit()

			err := s.repo.UpdateVisitCount(v.body)
			s.Equal(v.expectError, err)
		})
	}
}
func (s *GormSuite) TestUpdateVisitCountError() {
	testCase := []struct {
		name        string
		body        dto.PocketMessageWithRandomID
		expectError error
	}{
		{
			name: "update_username-error",
			body: dto.PocketMessageWithRandomID{
				UUID:  uuid.Nil,
				Visit: 0,
			},
			expectError: errors.New("record not found"),
		},
	}
	for _, v := range testCase {
		s.T().Run(v.name, func(t *testing.T) {
			s.mock.ExpectBegin()
			s.mock.ExpectExec(regexp.QuoteMeta("UPDATE `pocket_message_random_id` SET `visit`=?,`updated_at`=? WHERE pocket_message_uuid = ? AND `pocket_message_random_id`.`deleted_at` IS NULL")).
				WithArgs(v.body.Visit+1, AnyTime{}, uuid.Nil).
				WillReturnError(errors.New("record not found"))
			s.mock.ExpectRollback()

			err := s.repo.UpdateVisitCount(v.body)
			s.Equal(v.expectError, err)
		})
	}
}

// UpdatePocketMessage
func (s *GormSuite) TestUpdatePocketMessage() {
	testCase := []struct {
		name        string
		body        models.PocketMessage
		expectError error
	}{
		{
			name: "update_pocket_message-normal",
			body: models.PocketMessage{
				UUID:    uuid.Nil,
				Title:   "super idol",
				Content: "super",
			},
			expectError: nil,
		},
	}
	for _, v := range testCase {
		s.T().Run(v.name, func(t *testing.T) {
			s.mock.ExpectBegin()
			s.mock.ExpectExec(regexp.QuoteMeta("UPDATE `pocket_messages` SET `updated_at`=?,`title`=?,`content`=? WHERE uuid = ? AND `pocket_messages`.`deleted_at` IS NULL")).
				WithArgs(AnyTime{}, "super idol", "super", uuid.Nil).
				WillReturnResult(sqlmock.NewResult(1, 1))
			s.mock.ExpectCommit()

			err := s.repo.UpdatePocketMessage(v.body)
			s.Equal(v.expectError, err)
		})
	}
}
func (s *GormSuite) TestUpdatePocketMessageError() {
	testCase := []struct {
		name        string
		body        models.PocketMessage
		expectError error
	}{
		{
			name: "update_pocket_message-error",
			body: models.PocketMessage{
				UUID:    uuid.Nil,
				Title:   "super idol",
				Content: "super",
			},
			expectError: errors.New("record not found"),
		},
	}
	for _, v := range testCase {
		s.T().Run(v.name, func(t *testing.T) {
			s.mock.ExpectBegin()
			s.mock.ExpectExec(regexp.QuoteMeta("UPDATE `pocket_messages` SET `updated_at`=?,`title`=?,`content`=? WHERE uuid = ? AND `pocket_messages`.`deleted_at` IS NULL")).
				WithArgs(AnyTime{}, "super idol", "super", uuid.Nil).
				WillReturnError(errors.New("record not found"))
			s.mock.ExpectRollback()

			err := s.repo.UpdatePocketMessage(v.body)
			s.Equal(v.expectError, err)
		})
	}
}

// DeletePocketMessage
func (s *GormSuite) TestDeletePocketMessage() {
	testCase := []struct {
		name        string
		id          uuid.UUID
		expectError error
	}{
		{
			name:        "delete_pocket_message-normal",
			id:          uuid.Nil,
			expectError: nil,
		},
	}
	for _, v := range testCase {
		s.T().Run(v.name, func(t *testing.T) {
			s.mock.ExpectBegin()
			s.mock.ExpectExec(regexp.QuoteMeta("DELETE FROM `pocket_messages` WHERE uuid = ?")).
				WithArgs(uuid.Nil).
				WillReturnResult(sqlmock.NewResult(1, 1))
			s.mock.ExpectCommit()

			err := s.repo.DeletePocketMessage(v.id)
			s.Equal(v.expectError, err)
		})
	}
}
func (s *GormSuite) TestDeletePocketMessageError() {
	testCase := []struct {
		name        string
		id          uuid.UUID
		expectError error
	}{
		{
			name:        "delete_pocket_message-error",
			id:          uuid.Nil,
			expectError: errors.New("record not found"),
		},
	}
	for _, v := range testCase {
		s.T().Run(v.name, func(t *testing.T) {
			s.mock.ExpectBegin()
			s.mock.ExpectExec(regexp.QuoteMeta("DELETE FROM `pocket_messages` WHERE uuid = ?")).
				WithArgs(uuid.Nil).
				WillReturnError(errors.New("record not found"))
			s.mock.ExpectRollback()

			err := s.repo.DeletePocketMessage(v.id)
			s.Equal(v.expectError, err)
		})
	}
}

// GetPocketMessageByUserUUID
func (s *GormSuite) TestGetPocketMessageByUserUUID() {
	testCase := []struct {
		name        string
		id          uuid.UUID
		expectBody  []dto.OwnedMessage
		expectError error
	}{
		{
			name: "login-normal",
			id:   uuid.Nil,
			expectBody: []dto.OwnedMessage{
				{
					RandomID: "asdfghjkl",
					Title:    "waw",
					Content:  "super",
					Visit:    100,
				},
			},
			expectError: nil,
		},
	}
	for _, v := range testCase {
		s.T().Run(v.name, func(t *testing.T) {
			expectRow := s.mock.NewRows([]string{"random_id", "title", "content", "visit"}).
				AddRow("asdfghjkl", "waw", "super", 100)

			s.mock.ExpectQuery(regexp.QuoteMeta("SELECT pocket_message_random_id.random_id, pocket_messages.title, pocket_messages.content, pocket_message_random_id.visit FROM `pocket_messages` LEFT JOIN pocket_message_random_id ON pocket_messages.uuid = pocket_message_random_id.pocket_message_uuid WHERE pocket_messages.user_uuid = ? AND `pocket_messages`.`deleted_at` IS NULL")).
				WithArgs(v.id).
				WillReturnRows(expectRow)

			result, err := s.repo.GetPocketMessageByUserUUID(v.id)
			s.Equal(v.expectError, err)
			s.Equal(v.expectBody, result)
		})
	}
}
func (s *GormSuite) TestGetPocketMessageByUserUUIDError() {
	testCase := []struct {
		name        string
		id          uuid.UUID
		expectBody  []dto.OwnedMessage
		expectError error
	}{
		{
			name:        "login-error",
			id:          uuid.Nil,
			expectBody:  nil,
			expectError: errors.New("record not found"),
		},
	}
	for _, v := range testCase {
		s.T().Run(v.name, func(t *testing.T) {
			s.mock.ExpectQuery(regexp.QuoteMeta("SELECT pocket_message_random_id.random_id, pocket_messages.title, pocket_messages.content, pocket_message_random_id.visit FROM `pocket_messages` LEFT JOIN pocket_message_random_id ON pocket_messages.uuid = pocket_message_random_id.pocket_message_uuid WHERE pocket_messages.user_uuid = ? AND `pocket_messages`.`deleted_at` IS NULL")).
				WithArgs(v.id).
				WillReturnError(errors.New("record not found"))

			result, err := s.repo.GetPocketMessageByUserUUID(v.id)
			s.Equal(v.expectError, err)
			s.Equal(v.expectBody, result)
		})
	}
}
