package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	m "pocket-message/controllers/mock"
	"pocket-message/dto"
	"pocket-message/middleware"
	"pocket-message/models"
	"testing"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"
)

type PocketMessageSuite struct {
	suite.Suite
	handler PocketMessageHandler
}

func TestSuitePocketMessage(t *testing.T) {
	suite.Run(t, new(PocketMessageSuite))
}
func (s *PocketMessageSuite) SetupSuite() {

	handler := NewPocketMessageHandler(&m.MockPocketMessageServices{})
	s.handler = handler
}
func (s *PocketMessageSuite) TearDownSuite() {
}

// NewPocketMessage Unit Test
func (s *PocketMessageSuite) TestNewPocketMessage() {
	testCase := []struct {
		name          string
		method        string
		path          string
		body          models.PocketMessage
		expectCode    int
		expectMessage string
	}{
		{
			name:   "new_pocket_message-normal",
			method: http.MethodPost,
			path:   "/api/v1/pocket-messages",
			body: models.PocketMessage{
				Title:   "si anak lalai",
				Content: "iya kamu",
			},
			expectCode:    http.StatusCreated,
			expectMessage: "created",
		},
	}
	for _, v := range testCase {
		s.T().Run(v.name, func(t *testing.T) {
			res, _ := json.Marshal(v.body)

			r := httptest.NewRequest(v.method, "/", bytes.NewBuffer(res))
			w := httptest.NewRecorder()
			c := echo.New().NewContext(r, w)
			c.SetPath(v.path)
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
func (s *PocketMessageSuite) TestNewPocketMessageError() {
	testCase := []struct {
		name          string
		method        string
		path          string
		body          models.PocketMessage
		expectCode    int
		expectMessage string
	}{
		{
			name:   "new_pocket_message-error",
			method: http.MethodPost,
			path:   "/api/v1/pocket-messages",
			body: models.PocketMessage{
				Title:   "",
				Content: "iya kamu",
			},
			expectCode:    http.StatusInternalServerError,
			expectMessage: "error, title should not be empty",
		},
	}
	for _, v := range testCase {
		s.T().Run(v.name, func(t *testing.T) {
			res, _ := json.Marshal(v.body)

			r := httptest.NewRequest(v.method, "/", bytes.NewBuffer(res))
			w := httptest.NewRecorder()
			c := echo.New().NewContext(r, w)
			c.SetPath(v.path)
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

// GetPocketMessageByRandomID
func (s *PocketMessageSuite) TestGetPocketMessageByRandomID() {
	testCase := []struct {
		name          string
		method        string
		path          string
		expectBody    models.PocketMessage
		expectCode    int
		expectMessage string
	}{
		{
			name:   "get_pocket_message_by_random_id-normal",
			method: http.MethodPost,
			path:   "/api/v1/pocket-messages",
			expectBody: models.PocketMessage{
				Title:   "Ini Test",
				Content: "Ini juga Test",
			},
			expectCode:    http.StatusOK,
			expectMessage: "success",
		},
	}
	for _, v := range testCase {
		s.T().Run(v.name, func(t *testing.T) {

			r := httptest.NewRequest(v.method, "/", nil)
			w := httptest.NewRecorder()
			c := echo.New().NewContext(r, w)
			c.SetPath(v.path)
			c.SetParamNames("random_id")
			c.SetParamValues("ini_param_test")
			c.Request().Header.Set("Content-Type", "application/json")

			if s.NoError(s.handler.GetPocketMessageByRandomID(c)) {
				body := w.Body.Bytes()

				type response struct {
					Message string                        `json:"message"`
					Data    dto.PocketMessageWithRandomID `json:"data"`
				}
				var resp response
				err := json.Unmarshal(body, &resp)
				if err != nil {
					s.Error(err, "error unmarshalling")
				}

				s.Equal(v.expectCode, w.Result().StatusCode)
				s.Equal(v.expectMessage, resp.Message)
				s.Equal(v.expectBody.Title, resp.Data.Title)
				s.Equal(v.expectBody.Content, resp.Data.Content)
			}
		})
	}
}
func (s *PocketMessageSuite) TestGetPocketMessageByRandomIDError() {
	testCase := []struct {
		name          string
		method        string
		path          string
		expectBody    models.PocketMessage
		expectCode    int
		expectMessage string
	}{
		{
			name:   "get_pocket_message_by_random_id-error",
			method: http.MethodPost,
			path:   "/api/v1/pocket-messages",
			expectBody: models.PocketMessage{
				Title:   "",
				Content: "",
			},
			expectCode:    http.StatusInternalServerError,
			expectMessage: "error, random_id parameter can not be empty",
		},
	}
	for _, v := range testCase {
		s.T().Run(v.name, func(t *testing.T) {

			r := httptest.NewRequest(v.method, "/", nil)
			w := httptest.NewRecorder()
			c := echo.New().NewContext(r, w)
			c.SetPath(v.path)
			c.SetParamNames("random_id")
			c.SetParamValues("")
			c.Request().Header.Set("Content-Type", "application/json")

			if s.NoError(s.handler.GetPocketMessageByRandomID(c)) {
				body := w.Body.Bytes()

				type response struct {
					Message string                        `json:"message"`
					Data    dto.PocketMessageWithRandomID `json:"data"`
				}
				var resp response
				err := json.Unmarshal(body, &resp)
				if err != nil {
					s.Error(err, "error unmarshalling")
				}

				s.Equal(v.expectCode, w.Result().StatusCode)
				s.Equal(v.expectMessage, resp.Message)
				s.Equal(v.expectBody.Title, resp.Data.Title)
				s.Equal(v.expectBody.Content, resp.Data.Content)
			}
		})
	}
}

// UpdatePocketMessage
func (s *PocketMessageSuite) TestUpdatePocketMessage() {
	testCase := []struct {
		name          string
		method        string
		path          string
		body          models.PocketMessage
		expectCode    int
		expectMessage string
		paramName     string
		paramValue    string
	}{
		{
			name:   "update_pocket_message-normal",
			method: http.MethodPut,
			path:   "/api/v1/pocket-messages",
			body: models.PocketMessage{
				Title:   "untuk kamu",
				Content: "apakah kamu sehat?",
			},
			expectCode:    http.StatusOK,
			expectMessage: "updated",
			paramName:     "uuid",
			paramValue:    "00000000-0000-0000-0000-000000000000",
		},
	}
	for _, v := range testCase {
		s.T().Run(v.name, func(t *testing.T) {
			res, err := json.Marshal(v.body)
			if err != nil {
				s.Error(err, "error marshalling")
			}

			r := httptest.NewRequest(v.method, "/", bytes.NewBuffer(res))
			w := httptest.NewRecorder()
			c := echo.New().NewContext(r, w)
			c.SetParamNames(v.paramName)
			c.SetParamValues(v.paramValue)
			c.Request().Header.Set("Content-Type", "application/json")
			c.Request().Header.Set("test", "true")

			if s.NoError(s.handler.UpdatePocketMessage(c)) {
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
func (s *PocketMessageSuite) TestUpdatePocketMessageError() {
	testCase := []struct {
		name          string
		method        string
		path          string
		body          models.PocketMessage
		expectCode    int
		expectMessage string
		paramName     string
		paramValue    string
	}{
		{
			name:   "update_pocket_message-error",
			method: http.MethodPut,
			path:   "/api/v1/pocket-messages",
			body: models.PocketMessage{
				Title:   "",
				Content: "apakah kamu sehat?",
			},
			expectCode:    http.StatusInternalServerError,
			expectMessage: "error, title should not be empty",
			paramName:     "uuid",
			paramValue:    "00000000-0000-0000-0000-000000000000",
		},
	}
	for _, v := range testCase {
		s.T().Run(v.name, func(t *testing.T) {
			res, err := json.Marshal(v.body)
			if err != nil {
				s.Error(err, "error marshalling")
			}

			r := httptest.NewRequest(v.method, "/", bytes.NewBuffer(res))
			w := httptest.NewRecorder()
			c := echo.New().NewContext(r, w)
			c.SetParamNames(v.paramName)
			c.SetParamValues(v.paramValue)
			c.Request().Header.Set("Content-Type", "application/json")
			c.Request().Header.Set("test", "true")

			if s.NoError(s.handler.UpdatePocketMessage(c)) {
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

// DeletePocketMessage
func (s *PocketMessageSuite) TestDeletePocketMessage() {
	testCase := []struct {
		name   string
		method string
		path   string
		// body          models.PocketMessage
		expectCode    int
		expectMessage string
		paramName     string
		paramValue    string
	}{
		{
			name:   "delete_pocket_message-normal",
			method: http.MethodDelete,
			path:   "/api/v1/pocket-messages",
			// body: models.PocketMessage{
			// 	Title:   "untuk kamu",
			// 	Content: "apakah kamu sehat?",
			// },
			expectCode:    http.StatusOK,
			expectMessage: "deleted",
			paramName:     "uuid",
			paramValue:    "00000000-0000-0000-0000-000000000000",
		},
	}
	for _, v := range testCase {
		s.T().Run(v.name, func(t *testing.T) {

			r := httptest.NewRequest(v.method, "/", nil)
			w := httptest.NewRecorder()
			c := echo.New().NewContext(r, w)
			c.SetParamNames(v.paramName)
			c.SetParamValues(v.paramValue)
			c.Request().Header.Set("Content-Type", "application/json")
			c.Request().Header.Set("test", "true")

			if s.NoError(s.handler.DeletePocketMessage(c)) {
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
func (s *PocketMessageSuite) TestDeletePocketMessageError() {
	testCase := []struct {
		name   string
		method string
		path   string
		// body          models.PocketMessage
		expectCode    int
		expectMessage string
		paramName     string
		paramValue    string
	}{
		{
			name:   "delete_pocket_message-normal",
			method: http.MethodDelete,
			path:   "/api/v1/pocket-messages",
			// body: models.PocketMessage{
			// 	Title:   "untuk kamu",
			// 	Content: "apakah kamu sehat?",
			// },
			expectCode:    http.StatusInternalServerError,
			expectMessage: "invalid UUID length: 8",
			paramName:     "uuid",
			paramValue:    "00000000",
		},
	}
	for _, v := range testCase {
		s.T().Run(v.name, func(t *testing.T) {

			r := httptest.NewRequest(v.method, "/", nil)
			w := httptest.NewRecorder()
			c := echo.New().NewContext(r, w)
			c.SetParamNames(v.paramName)
			c.SetParamValues(v.paramValue)
			c.Request().Header.Set("Content-Type", "application/json")
			c.Request().Header.Set("test", "true")

			if s.NoError(s.handler.DeletePocketMessage(c)) {
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

// GetOwnedPocketMessage
func (s *PocketMessageSuite) TestGetOwnedPocketMessage() {
	testCase := []struct {
		name          string
		method        string
		path          string
		uuid          uuid.UUID
		username      string
		expectBody    []dto.OwnedMessage
		expectCode    int
		expectMessage string
	}{
		{
			name:     "get_owned_pocket_message-normal",
			method:   http.MethodGet,
			path:     "/api/v1/pocket-messages",
			uuid:     uuid.New(),
			username: "udin",
			expectBody: []dto.OwnedMessage{
				{
					Title:   "halo dunia",
					Content: "halo kamu",
					Visit:   1,
				},
			},
			expectCode:    http.StatusOK,
			expectMessage: "success",
		},
	}
	for _, v := range testCase {
		s.T().Run(v.name, func(t *testing.T) {

			r := httptest.NewRequest(v.method, "/", nil)
			w := httptest.NewRecorder()
			c := echo.New().NewContext(r, w)
			tok, err := middleware.GetToken(v.uuid, v.username)
			if err != nil {
				s.Error(err, "error get token")
			}
			token := fmt.Sprintf("Bearer %s", tok)
			c.Request().Header.Set("Authorization", token)
			c.Request().Header.Set("Content-Type", "application/json")

			if s.NoError(s.handler.GetOwnedPocketMessage(c)) {
				body := w.Body.Bytes()

				type response struct {
					Message string             `json:"message"`
					Data    []dto.OwnedMessage `json:"data"`
				}
				var resp response
				err := json.Unmarshal(body, &resp)
				if err != nil {
					s.Error(err, "error unmarshalling")
				}

				s.Equal(v.expectCode, w.Result().StatusCode)
				s.Equal(v.expectMessage, resp.Message)
				s.Equal(v.expectBody, resp.Data)
			}
		})
	}
}
func (s *PocketMessageSuite) TestGetOwnedPocketMessageError() {
	testCase := []struct {
		name          string
		method        string
		path          string
		uuid          uuid.UUID
		username      string
		expectBody    []dto.OwnedMessage
		expectCode    int
		expectMessage string
	}{
		{
			name:          "get_owned_pocket_message-normal",
			method:        http.MethodGet,
			path:          "/api/v1/pocket-messages",
			uuid:          uuid.New(),
			username:      "udin",
			expectBody:    nil,
			expectCode:    http.StatusInternalServerError,
			expectMessage: "authorization header not found",
		},
	}
	for _, v := range testCase {
		s.T().Run(v.name, func(t *testing.T) {

			r := httptest.NewRequest(v.method, "/", nil)
			w := httptest.NewRecorder()
			c := echo.New().NewContext(r, w)
			c.Request().Header.Set("Authorization", "")
			c.Request().Header.Set("Content-Type", "application/json")

			if s.NoError(s.handler.GetOwnedPocketMessage(c)) {
				body := w.Body.Bytes()

				type response struct {
					Message string             `json:"message"`
					Data    []dto.OwnedMessage `json:"data"`
				}
				var resp response
				err := json.Unmarshal(body, &resp)
				if err != nil {
					s.Error(err, "error unmarshalling")
				}

				s.Equal(v.expectCode, w.Result().StatusCode)
				s.Equal(v.expectMessage, resp.Message)
				s.Equal(v.expectBody, resp.Data)
			}
		})
	}
}
