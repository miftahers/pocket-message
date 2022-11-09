package repositories

import (
	"errors"
	"pocket-message/dto"
	"pocket-message/models"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type GormPocketMessageSuite struct {
	suite.Suite
	mock sqlmock.Sqlmock
	repo Database
}

func TestSuiteGormPocketMessage(t *testing.T) {
	suite.Run(t, new(GormPocketMessageSuite))
}

func (s *GormPocketMessageSuite) SetupSuite() {
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

func (s *GormPocketMessageSuite) TearDownSuite() {}

// NewPocketMessage
func (s *GormPocketMessageSuite) TestNewPocketMessage() {
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
func (s *GormPocketMessageSuite) TestNewPocketMessageError() {
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
func (s *GormPocketMessageSuite) TestSaveNewRandomID() {
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
func (s *GormPocketMessageSuite) TestSaveNewRandomIDError() {
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
func (s *GormPocketMessageSuite) TestGetPocketMessageByRandomID() {
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
func (s *GormPocketMessageSuite) TestGetPocketMessageByRandomIDError() {
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
func (s *GormPocketMessageSuite) TestUpdateVisitCount() {
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
func (s *GormPocketMessageSuite) TestUpdateVisitCountError() {
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
func (s *GormPocketMessageSuite) TestUpdatePocketMessage() {
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
func (s *GormPocketMessageSuite) TestUpdatePocketMessageError() {
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
func (s *GormPocketMessageSuite) TestDeletePocketMessage() {
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
func (s *GormPocketMessageSuite) TestDeletePocketMessageError() {
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
func (s *GormPocketMessageSuite) TestGetPocketMessageByUserUUID() {
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
func (s *GormPocketMessageSuite) TestGetPocketMessageByUserUUIDError() {
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
