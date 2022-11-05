package controllers

import (
	"bytes"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"pocket-message/middleware"
	"pocket-message/models"
	"pocket-message/repositories"
	"pocket-message/services"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type PocketMessageSuite struct {
	suite.Suite
	mock    sqlmock.Sqlmock
	handler pocketMessageHandler
}

type AnyTime struct{}

// Match satisfies sqlmock.Argument interface
func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}

func TestSuiteBlog(t *testing.T) {
	suite.Run(t, new(PocketMessageSuite))
}

func (s *PocketMessageSuite) SetupSuite() {
	// Create mock db
	db, mock, err := sqlmock.New()
	s.NoError(err)

	var GormDB *gorm.DB
	GormDB, err = gorm.Open(mysql.New(mysql.Config{
		SkipInitializeWithVersion: true,
		Conn:                      db,
	}), &gorm.Config{})
	s.NoError(err)

	repo := repositories.NewGorm(GormDB)
	_, pmServices := services.NewServices(repo)
	handler := pocketMessageHandler{
		PocketMessageServices: pmServices,
	}
	s.handler = handler
	s.mock = mock
}

func (s *PocketMessageSuite) TearDownSuite() {
	s.mock = nil
}

func (s *PocketMessageSuite) TestNewPocketMessage() {

	testCase := []struct {
		name          string
		path          string
		method        string
		expectCode    int
		body          models.PocketMessage
		expectMessage string
	}{
		{
			name:       "new pocket message - normal",
			path:       "/pocket-messages",
			method:     http.MethodPost,
			expectCode: http.StatusCreated,
			body: models.PocketMessage{
				UUID:     uuid.Nil,
				Title:    "Untuk kamu yang disana~",
				Content:  "I Love U",
				UserUUID: uuid.Nil,
			},
			expectMessage: "created",
		},
	}

	for _, v := range testCase {
		s.T().Run(v.name, func(t *testing.T) {

			s.mock.ExpectBegin()
			s.mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `pocket_messages` (`created_at`,`updated_at`,`deleted_at`,`uuid`,`title`,`content`,`user_uuid`) VALUES (?,?,?,?,?,?,?)")).
				WithArgs(AnyTime{}, AnyTime{}, nil, uuid.Nil, "Untuk kamu yang disana~", "I Love U", uuid.Nil).
				WillReturnResult(sqlmock.NewResult(1, 1))
			s.mock.ExpectCommit()

			s.mock.ExpectBegin()
			s.mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `pocket_message_random_id` (`created_at`,`updated_at`,`deleted_at`,`uuid`,`title`,`content`,`user_uuid`) VALUES (?,?,?,?,?,?,?)")).
				WithArgs(AnyTime{}, AnyTime{}, nil, uuid.Nil, "Untuk kamu yang disana~", "I Love U", uuid.Nil).
				WillReturnResult(sqlmock.NewResult(1, 1))
			s.mock.ExpectCommit()

			res, _ := json.Marshal(v.body)
			r := httptest.NewRequest(v.method, "/", bytes.NewBuffer(res))
			w := httptest.NewRecorder()
			c := echo.New().NewContext(r, w)
			c.SetPath(v.path)

			token, err := middleware.GetToken(uuid.Nil, "")
			if err != nil {
				s.Error(err)
			}

			auth := fmt.Sprintf("Bearer %s", token)
			c.Request().Header.Add("Authorization", auth)
			c.Request().Header.Set("test", "true")
			c.Request().Header.Set("Content-Type", "application/json")

			if s.NoError(s.handler.NewPocketMessage(c)) {
				body := w.Body.Bytes()

				type response struct {
					Message string `json:"message"`
				}

				var resp response
				err := json.Unmarshal(body, &resp)
				if err != nil {
					s.Error(err, "error unmarshalling")
				}

				s.Equal(v.expectCode, w.Result().StatusCode)
				s.Equal(v.expectMessage, resp.Message)
			}
		})
	}
}
