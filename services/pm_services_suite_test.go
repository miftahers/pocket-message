package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"pocket-message/dto"
	"pocket-message/middleware"
	"pocket-message/models"
	m "pocket-message/services/mock"
	"testing"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"
)

type PocketMessageSuite struct {
	suite.Suite
	service PocketMessageServices
}

func TestSuitePocketMessage(t *testing.T) {
	suite.Run(t, new(PocketMessageSuite))
}

func (s *PocketMessageSuite) SetupSuite() {
	service := NewPocketMessageServices(&m.MockGorm{})
	s.service = service
}

func (s *PocketMessageSuite) TearDownSuite() {}

// NewPocketMessage
func (s *PocketMessageSuite) TestNewPocketMessage() {
	testCase := []struct {
		name        string
		body        models.PocketMessage
		method      string
		expectError error
	}{
		{
			name: "new_pocket_message-normal",
			body: models.PocketMessage{
				Title:   "yes",
				Content: "no",
			},
			method:      http.MethodPost,
			expectError: nil,
		},
	}
	for _, v := range testCase {
		s.T().Run(v.name, func(t *testing.T) {
			res, _ := json.Marshal(v.body)
			r := httptest.NewRequest(v.method, "/", bytes.NewBuffer(res))
			w := httptest.NewRecorder()
			c := echo.New().NewContext(r, w)
			c.Request().Header.Set("Content-Type", "application/json")

			token, err := middleware.GetToken(uuid.Nil, "super")
			if err != nil {
				s.Error(err, "error get token")
			}
			bearer := fmt.Sprintf("Bearer %s", token)
			c.Request().Header.Set("Authorization", bearer)

			err = s.service.NewPocketMessage(c)
			s.Equal(v.expectError, err)
		})
	}
}
func (s *PocketMessageSuite) TestNewPocketMessageErrorBind() {
	testCase := []struct {
		name        string
		body        models.PocketMessage
		method      string
		expectError error
	}{
		{
			name: "new_pocket_message-error_bind",
			body: models.PocketMessage{
				Title:   "abd",
				Content: "no",
			},
			method:      http.MethodPost,
			expectError: echo.NewHTTPError(415, "Unsupported Media Type"),
		},
	}
	for _, v := range testCase {
		s.T().Run(v.name, func(t *testing.T) {
			res, _ := json.Marshal(v.body)
			r := httptest.NewRequest(v.method, "/", bytes.NewBuffer(res))
			w := httptest.NewRecorder()
			c := echo.New().NewContext(r, w)

			token, err := middleware.GetToken(uuid.Nil, "super")
			if err != nil {
				s.Error(err, "error get token")
			}
			bearer := fmt.Sprintf("Bearer %s", token)
			c.Request().Header.Set("Authorization", bearer)

			err = s.service.NewPocketMessage(c)
			s.Equal(v.expectError, err)
		})
	}
}
func (s *PocketMessageSuite) TestNewPocketMessageErrorTitleEmpty() {
	testCase := []struct {
		name        string
		body        models.PocketMessage
		method      string
		expectError error
	}{
		{
			name: "new_pocket_message-error_title_empty",
			body: models.PocketMessage{
				Title:   "",
				Content: "no",
			},
			method:      http.MethodPost,
			expectError: errors.New("error, title should not be empty"),
		},
	}
	for _, v := range testCase {
		s.T().Run(v.name, func(t *testing.T) {
			res, _ := json.Marshal(v.body)
			r := httptest.NewRequest(v.method, "/", bytes.NewBuffer(res))
			w := httptest.NewRecorder()
			c := echo.New().NewContext(r, w)
			c.Request().Header.Set("Content-Type", "application/json")
			token, err := middleware.GetToken(uuid.Nil, "super")
			if err != nil {
				s.Error(err, "error get token")
			}
			bearer := fmt.Sprintf("Bearer %s", token)
			c.Request().Header.Set("Authorization", bearer)

			err = s.service.NewPocketMessage(c)
			s.Equal(v.expectError, err)
		})
	}
}
func (s *PocketMessageSuite) TestNewPocketMessageErrorContentEmpty() {
	testCase := []struct {
		name        string
		body        models.PocketMessage
		method      string
		expectError error
	}{
		{
			name: "new_pocket_message-error_content_empty",
			body: models.PocketMessage{
				Title:   "asd",
				Content: "",
			},
			method:      http.MethodPost,
			expectError: errors.New("error, content should not be empty"),
		},
	}
	for _, v := range testCase {
		s.T().Run(v.name, func(t *testing.T) {
			res, _ := json.Marshal(v.body)
			r := httptest.NewRequest(v.method, "/", bytes.NewBuffer(res))
			w := httptest.NewRecorder()
			c := echo.New().NewContext(r, w)
			c.Request().Header.Set("Content-Type", "application/json")
			token, err := middleware.GetToken(uuid.Nil, "super")
			if err != nil {
				s.Error(err, "error get token")
			}
			bearer := fmt.Sprintf("Bearer %s", token)
			c.Request().Header.Set("Authorization", bearer)

			err = s.service.NewPocketMessage(c)
			s.Equal(v.expectError, err)
		})
	}
}
func (s *PocketMessageSuite) TestNewPocketMessageErrorSavePocketMessage() {
	testCase := []struct {
		name        string
		body        models.PocketMessage
		method      string
		expectError error
	}{
		{
			name: "new_pocket_message-error_save_pocket_message",
			body: models.PocketMessage{
				Title:   "super",
				Content: "asd",
			},
			method:      http.MethodPost,
			expectError: errors.New("database error"),
		},
	}
	for _, v := range testCase {
		s.T().Run(v.name, func(t *testing.T) {
			res, _ := json.Marshal(v.body)
			r := httptest.NewRequest(v.method, "/", bytes.NewBuffer(res))
			w := httptest.NewRecorder()
			c := echo.New().NewContext(r, w)
			c.Request().Header.Set("Content-Type", "application/json")
			token, err := middleware.GetToken(uuid.Nil, "super")
			if err != nil {
				s.Error(err, "error get token")
			}
			bearer := fmt.Sprintf("Bearer %s", token)
			c.Request().Header.Set("Authorization", bearer)

			err = s.service.NewPocketMessage(c)
			s.Equal(v.expectError, err)
		})
	}
}
func (s *PocketMessageSuite) TestNewPocketMessageErrorSaveRandomID() {
	testCase := []struct {
		name        string
		body        models.PocketMessage
		method      string
		expectError error
	}{
		{
			name: "new_pocket_message-error_save_random_id",
			body: models.PocketMessage{
				Title:   "super",
				Content: "asd",
			},
			method:      http.MethodPost,
			expectError: errors.New("database error"),
		},
	}
	for _, v := range testCase {
		s.T().Run(v.name, func(t *testing.T) {
			res, _ := json.Marshal(v.body)
			r := httptest.NewRequest(v.method, "/", bytes.NewBuffer(res))
			w := httptest.NewRecorder()
			c := echo.New().NewContext(r, w)
			c.Request().Header.Set("Content-Type", "application/json")
			token, err := middleware.GetToken(uuid.Nil, "super")
			if err != nil {
				s.Error(err, "error get token")
			}
			bearer := fmt.Sprintf("Bearer %s", token)
			c.Request().Header.Set("Authorization", bearer)
			c.Request().Header.Set("test", "true")

			err = s.service.NewPocketMessage(c)
			s.Equal(v.expectError, err)
		})
	}
}

// GetPocketMessageByRandomID
func (s *PocketMessageSuite) TestGetPocketMessageByRandomID() {
	testCase := []struct {
		name        string
		paramName   string
		paramValue  string
		expectBody  dto.PocketMessageWithRandomID
		expectError error
	}{
		{
			name:       "get_pocket_message_by_random_id-normal",
			paramName:  "random_id",
			paramValue: "asdfghjk",
			expectBody: dto.PocketMessageWithRandomID{
				UUID:  uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				Title: "selsya bahagia",
			},
			expectError: nil,
		},
	}
	for _, v := range testCase {
		s.T().Run(v.name, func(t *testing.T) {
			r := httptest.NewRequest(http.MethodGet, "/", nil)
			w := httptest.NewRecorder()
			c := echo.New().NewContext(r, w)
			c.SetParamNames(v.paramName)
			c.SetParamValues(v.paramValue)
			c.Request().Header.Set("Content-Type", "application/json")

			result, err := s.service.GetPocketMessageByRandomID(c)
			s.Equal(v.expectBody, result)
			s.Equal(v.expectError, err)
		})
	}
}
func (s *PocketMessageSuite) TestGetPocketMessageByRandomIDErrorParamRandomIDEmpty() {
	testCase := []struct {
		name        string
		paramName   string
		paramValue  string
		expectError error
	}{
		{
			name:        "get_pocket_message_by_random_id-error_param_random_id_empty",
			paramName:   "random_id",
			paramValue:  "",
			expectError: errors.New("error, random_id parameter can not be empty"),
		},
	}
	for _, v := range testCase {
		s.T().Run(v.name, func(t *testing.T) {
			r := httptest.NewRequest(http.MethodGet, "/", nil)
			w := httptest.NewRecorder()
			c := echo.New().NewContext(r, w)
			c.SetParamNames(v.paramName)
			c.SetParamValues(v.paramValue)
			c.Request().Header.Set("Content-Type", "application/json")

			_, err := s.service.GetPocketMessageByRandomID(c)
			s.Equal(v.expectError, err)
		})
	}
}
func (s *PocketMessageSuite) TestGetPocketMessageByRandomIDErrorRepo() {
	testCase := []struct {
		name        string
		paramName   string
		paramValue  string
		expectError error
	}{
		{
			name:        "get_pocket_message_by_random_id-error_repo",
			paramName:   "random_id",
			paramValue:  "superidol",
			expectError: errors.New("record not found"),
		},
	}
	for _, v := range testCase {
		s.T().Run(v.name, func(t *testing.T) {
			r := httptest.NewRequest(http.MethodGet, "/", nil)
			w := httptest.NewRecorder()
			c := echo.New().NewContext(r, w)
			c.SetParamNames(v.paramName)
			c.SetParamValues(v.paramValue)
			c.Request().Header.Set("Content-Type", "application/json")

			_, err := s.service.GetPocketMessageByRandomID(c)
			s.Equal(v.expectError, err)
		})
	}
}
func (s *PocketMessageSuite) TestGetPocketMessageByRandomIDErrorUpdateVisitCount() {
	testCase := []struct {
		name        string
		paramName   string
		paramValue  string
		expectError error
	}{
		{
			name:        "get_pocket_message_by_random_id-error_update_visit_count",
			paramName:   "random_id",
			paramValue:  "igantenk",
			expectError: errors.New("record not found"),
		},
	}
	for _, v := range testCase {
		s.T().Run(v.name, func(t *testing.T) {
			r := httptest.NewRequest(http.MethodGet, "/", nil)
			w := httptest.NewRecorder()
			c := echo.New().NewContext(r, w)
			c.SetParamNames(v.paramName)
			c.SetParamValues(v.paramValue)
			c.Request().Header.Set("Content-Type", "application/json")

			_, err := s.service.GetPocketMessageByRandomID(c)
			s.Equal(v.expectError, err)
		})
	}
}

// UpdatePocketMessage
func (s *PocketMessageSuite) TestUpdatePocketMessage() {
	testCase := []struct {
		name        string
		paramName   string
		paramValue  string
		body        models.PocketMessage
		expectError error
	}{
		{
			name:       "update_pocket_message-normal",
			paramName:  "uuid",
			paramValue: uuid.Nil.String(),
			body: models.PocketMessage{
				Title:   "damn",
				Content: "idol",
			},
			expectError: nil,
		},
	}
	for _, v := range testCase {
		s.T().Run(v.name, func(t *testing.T) {
			res, _ := json.Marshal(v.body)
			r := httptest.NewRequest(http.MethodGet, "/", bytes.NewBuffer(res))
			w := httptest.NewRecorder()
			c := echo.New().NewContext(r, w)
			c.SetParamNames(v.paramName)
			c.SetParamValues(v.paramValue)
			c.Request().Header.Set("Content-Type", "application/json")

			err := s.service.UpdatePocketMessage(c)
			s.Equal(v.expectError, err)
		})
	}
}
func (s *PocketMessageSuite) TestUpdatePocketMessageErrorBinding() {
	testCase := []struct {
		name        string
		paramName   string
		paramValue  string
		body        models.PocketMessage
		expectError error
	}{
		{
			name:       "update_pocket_message-error_binding",
			paramName:  "uuid",
			paramValue: uuid.Nil.String(),
			body: models.PocketMessage{
				Title:   "damn",
				Content: "idol",
			},
			expectError: echo.NewHTTPError(415, "Unsupported Media Type"),
		},
	}
	for _, v := range testCase {
		s.T().Run(v.name, func(t *testing.T) {
			res, _ := json.Marshal(v.body)
			r := httptest.NewRequest(http.MethodGet, "/", bytes.NewBuffer(res))
			w := httptest.NewRecorder()
			c := echo.New().NewContext(r, w)
			c.SetParamNames(v.paramName)
			c.SetParamValues(v.paramValue)

			err := s.service.UpdatePocketMessage(c)
			s.Equal(v.expectError, err)
		})
	}
}
func (s *PocketMessageSuite) TestUpdatePocketMessageErrorTitleEmpty() {
	testCase := []struct {
		name        string
		paramName   string
		paramValue  string
		body        models.PocketMessage
		expectError error
	}{
		{
			name:       "update_pocket_message-error_empty_title",
			paramName:  "uuid",
			paramValue: uuid.Nil.String(),
			body: models.PocketMessage{
				Title:   "",
				Content: "idol",
			},
			expectError: errors.New("error, title should not be empty"),
		},
	}
	for _, v := range testCase {
		s.T().Run(v.name, func(t *testing.T) {
			res, _ := json.Marshal(v.body)
			r := httptest.NewRequest(http.MethodGet, "/", bytes.NewBuffer(res))
			w := httptest.NewRecorder()
			c := echo.New().NewContext(r, w)
			c.SetParamNames(v.paramName)
			c.SetParamValues(v.paramValue)
			c.Request().Header.Set("Content-Type", "application/json")

			err := s.service.UpdatePocketMessage(c)
			s.Equal(v.expectError, err)
		})
	}
}
func (s *PocketMessageSuite) TestUpdatePocketMessageErrorContentEmpty() {
	testCase := []struct {
		name        string
		paramName   string
		paramValue  string
		body        models.PocketMessage
		expectError error
	}{
		{
			name:       "update_pocket_message-error_empty_content",
			paramName:  "uuid",
			paramValue: uuid.Nil.String(),
			body: models.PocketMessage{
				Title:   "asd",
				Content: "",
			},
			expectError: errors.New("error, content should not be empty"),
		},
	}
	for _, v := range testCase {
		s.T().Run(v.name, func(t *testing.T) {
			res, _ := json.Marshal(v.body)
			r := httptest.NewRequest(http.MethodGet, "/", bytes.NewBuffer(res))
			w := httptest.NewRecorder()
			c := echo.New().NewContext(r, w)
			c.SetParamNames(v.paramName)
			c.SetParamValues(v.paramValue)
			c.Request().Header.Set("Content-Type", "application/json")

			err := s.service.UpdatePocketMessage(c)
			s.Equal(v.expectError, err)
		})
	}
}
func (s *PocketMessageSuite) TestUpdatePocketMessageErrorParsingUUID() {
	testCase := []struct {
		name        string
		paramName   string
		paramValue  string
		body        models.PocketMessage
		expectError error
	}{
		{
			name:       "update_pocket_message-error_parsing",
			paramName:  "uuid",
			paramValue: "00000000-000-00-0-00000",
			body: models.PocketMessage{
				Title:   "asd",
				Content: "asd",
			},
			expectError: errors.New("uuid invalid"),
		},
	}
	for _, v := range testCase {
		s.T().Run(v.name, func(t *testing.T) {
			res, _ := json.Marshal(v.body)
			r := httptest.NewRequest(http.MethodGet, "/", bytes.NewBuffer(res))
			w := httptest.NewRecorder()
			c := echo.New().NewContext(r, w)
			c.SetParamNames(v.paramName)
			c.SetParamValues(v.paramValue)
			c.Request().Header.Set("Content-Type", "application/json")

			err := s.service.UpdatePocketMessage(c)
			s.Equal(v.expectError, err)
		})
	}
}
func (s *PocketMessageSuite) TestUpdatePocketMessageErrorRepo() {
	testCase := []struct {
		name        string
		paramName   string
		paramValue  string
		body        models.PocketMessage
		expectError error
	}{
		{
			name:       "update_pocket_message-error_repo",
			paramName:  "uuid",
			paramValue: uuid.Nil.String(),
			body: models.PocketMessage{
				Title:   "super",
				Content: "asd",
			},
			expectError: errors.New("database error"),
		},
	}
	for _, v := range testCase {
		s.T().Run(v.name, func(t *testing.T) {
			res, _ := json.Marshal(v.body)
			r := httptest.NewRequest(http.MethodGet, "/", bytes.NewBuffer(res))
			w := httptest.NewRecorder()
			c := echo.New().NewContext(r, w)
			c.SetParamNames(v.paramName)
			c.SetParamValues(v.paramValue)
			c.Request().Header.Set("Content-Type", "application/json")

			err := s.service.UpdatePocketMessage(c)
			s.Equal(v.expectError, err)
		})
	}
}

// DeletePocketMessage
func (s *PocketMessageSuite) TestDeletePocketMessage() {
	testCase := []struct {
		name        string
		paramName   string
		paramValue  string
		expectError error
	}{
		{
			name:        "delete_pocket_message-normal",
			paramName:   "uuid",
			paramValue:  "00000000-0000-0000-0000-000000000001",
			expectError: nil,
		},
	}
	for _, v := range testCase {
		s.T().Run(v.name, func(t *testing.T) {
			r := httptest.NewRequest(http.MethodGet, "/", nil)
			w := httptest.NewRecorder()
			c := echo.New().NewContext(r, w)
			c.SetParamNames(v.paramName)
			c.SetParamValues(v.paramValue)
			c.Request().Header.Set("Content-Type", "application/json")

			err := s.service.DeletePocketMessage(c)
			s.Equal(v.expectError, err)
		})
	}
}
func (s *PocketMessageSuite) TestDeletePocketMessageErrorParsing() {
	testCase := []struct {
		name        string
		paramName   string
		paramValue  string
		expectError error
	}{
		{
			name:        "delete_pocket_message-error_oarsing",
			paramName:   "uuid",
			paramValue:  "00000000-0000-000000-000000000000",
			expectError: errors.New("uuid invalid"),
		},
	}
	for _, v := range testCase {
		s.T().Run(v.name, func(t *testing.T) {
			r := httptest.NewRequest(http.MethodGet, "/", nil)
			w := httptest.NewRecorder()
			c := echo.New().NewContext(r, w)
			c.SetParamNames(v.paramName)
			c.SetParamValues(v.paramValue)
			c.Request().Header.Set("Content-Type", "application/json")

			err := s.service.DeletePocketMessage(c)
			s.Equal(v.expectError, err)
		})
	}
}
func (s *PocketMessageSuite) TestDeletePocketMessageErrorDB() {
	testCase := []struct {
		name        string
		paramName   string
		paramValue  string
		expectError error
	}{
		{
			name:        "delete_pocket_message-error_db",
			paramName:   "uuid",
			paramValue:  "00000000-0000-0000-0000-000000000000",
			expectError: errors.New("record not found"),
		},
	}
	for _, v := range testCase {
		s.T().Run(v.name, func(t *testing.T) {
			r := httptest.NewRequest(http.MethodGet, "/", nil)
			w := httptest.NewRecorder()
			c := echo.New().NewContext(r, w)
			c.SetParamNames(v.paramName)
			c.SetParamValues(v.paramValue)
			c.Request().Header.Set("Content-Type", "application/json")

			err := s.service.DeletePocketMessage(c)
			s.Equal(v.expectError, err)
		})
	}
}

// GetUserPocketMessage
func (s *PocketMessageSuite) TestGetUserPocketmessage() {
	testCase := []struct {
		name        string
		paramName   string
		paramValue  string
		expectBody  []dto.OwnedMessage
		expectError error
	}{
		{
			name:       "get_pocket_message_by_random_id-normal",
			paramName:  "random_id",
			paramValue: "asdfghjk",
			expectBody: []dto.OwnedMessage{
				{
					RandomID: "akasupas",
					Title:    "vtuber",
					Content:  "donation",
					Visit:    1000,
				},
			},
			expectError: nil,
		},
	}
	for _, v := range testCase {
		s.T().Run(v.name, func(t *testing.T) {
			r := httptest.NewRequest(http.MethodGet, "/", nil)
			w := httptest.NewRecorder()
			c := echo.New().NewContext(r, w)
			c.SetParamNames(v.paramName)
			c.SetParamValues(v.paramValue)
			c.Request().Header.Set("Content-Type", "application/json")

			token, err := middleware.GetToken(uuid.MustParse("00000000-0000-0000-0000-000000000001"), "super")
			if err != nil {
				s.Error(err, "error get token")
			}
			bearer := fmt.Sprintf("Bearer %s", token)
			c.Request().Header.Set("Authorization", bearer)

			result, err := s.service.GetUserPocketMessage(c)
			s.Equal(v.expectBody, result)
			s.Equal(v.expectError, err)
		})
	}
}
func (s *PocketMessageSuite) TestGetUserPocketmessageErrorDecodeJWT() {
	testCase := []struct {
		name        string
		paramName   string
		paramValue  string
		expectBody  []dto.OwnedMessage
		expectError error
	}{
		{
			name:        "get_pocket_message_by_random_id-error_decode_jwt",
			paramName:   "random_id",
			paramValue:  "asdfghjk",
			expectBody:  nil,
			expectError: errors.New("authorization header not found"),
		},
	}
	for _, v := range testCase {
		s.T().Run(v.name, func(t *testing.T) {
			r := httptest.NewRequest(http.MethodGet, "/", nil)
			w := httptest.NewRecorder()
			c := echo.New().NewContext(r, w)
			c.SetParamNames(v.paramName)
			c.SetParamValues(v.paramValue)
			c.Request().Header.Set("Content-Type", "application/json")

			result, err := s.service.GetUserPocketMessage(c)
			s.Equal(v.expectBody, result)
			s.Equal(v.expectError, err)
		})
	}
}
func (s *PocketMessageSuite) TestGetUserPocketmessageErrorDB() {
	testCase := []struct {
		name        string
		paramName   string
		paramValue  string
		expectBody  []dto.OwnedMessage
		expectError error
	}{
		{
			name:        "get_pocket_message_by_random_id-error_db",
			paramName:   "random_id",
			paramValue:  "asdfghjk",
			expectBody:  nil,
			expectError: errors.New("record not found"),
		},
	}
	for _, v := range testCase {
		s.T().Run(v.name, func(t *testing.T) {
			r := httptest.NewRequest(http.MethodGet, "/", nil)
			w := httptest.NewRecorder()
			c := echo.New().NewContext(r, w)
			c.SetParamNames(v.paramName)
			c.SetParamValues(v.paramValue)
			c.Request().Header.Set("Content-Type", "application/json")
			token, err := middleware.GetToken(uuid.MustParse("00000000-0000-0000-0000-000000000000"), "super")
			if err != nil {
				s.Error(err, "error get token")
			}
			bearer := fmt.Sprintf("Bearer %s", token)
			c.Request().Header.Set("Authorization", bearer)

			result, err := s.service.GetUserPocketMessage(c)
			s.Equal(v.expectBody, result)
			s.Equal(v.expectError, err)
		})
	}
}
