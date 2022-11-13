package pocket_messages

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"pocket-message/configs"
	"pocket-message/dto"
	"pocket-message/middleware"
	"pocket-message/models"
	mocks "pocket-message/services/pocket_messages/mocks"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type PocketMessageSuite struct {
	suite.Suite
	handler PocketMessageHandler
	service *mocks.IPocketMessageServices
}

func (s *PocketMessageSuite) SetupSuite() {
	// mock service
	service := new(mocks.IPocketMessageServices)
	handler := PocketMessageHandler{
		IPocketMessageServices: service,
	}

	s.service = service
	s.handler = handler

	// Init Config
	cfg := &configs.Config{}
	cfg.TokenSecret = "test"
	configs.Cfg = cfg
}
func (s *PocketMessageSuite) TearDownSuite() {
}

func TestPocketMessageSuite(t *testing.T) {
	suite.Run(t, new(PocketMessageSuite))
}

func (s *PocketMessageSuite) TestNewPocketMessage() {
	testCases := []struct {
		name             string
		method           string
		path             string
		body             models.PocketMessage
		contentType      string
		wantJWTError     bool
		wantServiceError bool
		expectStatusCode int
		expectMessage    string
	}{
		{
			name:   "test_new_pocket_message-normal",
			method: http.MethodPost,
			path:   "/api/v1/pocket-messages",
			body: models.PocketMessage{
				Title:   "Tentang Hari yang Cerah",
				Content: "Kenapa datang hari dimana aku tidak semangat?",
			},
			contentType:      "application/json",
			expectStatusCode: http.StatusCreated,
			expectMessage:    "created",
		},
		{
			name:   "test_new_pocket_message-error_binding",
			method: http.MethodPost,
			path:   "/api/v1/pocket-messages",
			body: models.PocketMessage{
				Title:   "Tentang Hari yang Cerah",
				Content: "Kenapa datang hari dimana aku tidak semangat?",
			},
			contentType:      "",
			expectStatusCode: http.StatusBadRequest,
			expectMessage:    "code=415, message=Unsupported Media Type",
		},
		{
			name:   "test_new_pocket_message-error_title_empty",
			method: http.MethodPost,
			path:   "/api/v1/pocket-messages",
			body: models.PocketMessage{
				Title:   "",
				Content: "Kenapa datang hari dimana aku tidak semangat?",
			},
			contentType:      "application/json",
			expectStatusCode: http.StatusBadRequest,
			expectMessage:    "",
		},
		{
			name:   "test_new_pocket_message-error_content_empty",
			method: http.MethodPost,
			path:   "/api/v1/pocket-messages",
			body: models.PocketMessage{
				Title:   "Super idol",
				Content: "",
			},
			contentType:      "application/json",
			expectStatusCode: http.StatusBadRequest,
			expectMessage:    "",
		},
		{
			name:   "test_new_pocket_message-error_auth",
			method: http.MethodPost,
			path:   "/api/v1/pocket-messages",
			body: models.PocketMessage{
				Title:   "Tentang Hari yang Cerah",
				Content: "Kenapa datang hari dimana aku tidak semangat?",
			},
			contentType:      "application/json",
			wantJWTError:     true,
			expectStatusCode: http.StatusBadRequest,
			expectMessage:    "authorization header not found",
		},
		// {
		// 	name:   "test_new_pocket_message-error_service",
		// 	method: http.MethodPost,
		// 	path:   "/api/v1/pocket-messages",
		// 	body: models.PocketMessage{
		// 		Title:   "Tentang Hari yang Cerah",
		// 		Content: "Kenapa datang hari dimana aku tidak semangat?",
		// 	},
		// 	contentType:      "application/json",
		// 	wantServiceError: true,
		// 	expectStatusCode: http.StatusCreated,
		// 	expectMessage:    "created",
		// },
	}

	for _, v := range testCases {
		s.T().Run(v.name, func(t *testing.T) {
			if v.wantServiceError {
				s.service.
					On("NewPocketMessage",
						v.body,
						dto.Token{}).
					Return(errors.New("service error"))
			} else {
				s.service.
					On("NewPocketMessage",
						v.body,
						dto.Token{
							UUID:     uuid.Nil,
							Username: "userTest",
						}).
					Return(nil)
			}

			res, _ := json.Marshal(v.body)
			r := httptest.NewRequest(v.method, "/", bytes.NewBuffer(res))
			w := httptest.NewRecorder()
			c := echo.New().NewContext(r, w)
			c.SetPath(v.path)
			c.Request().Header.Set("Content-Type", v.contentType)

			if !v.wantJWTError {
				tokenStr, err := middleware.GetToken(uuid.Nil, "userTest")
				if err != nil {
					s.Error(err, "error get token")
				}
				jwtoken := fmt.Sprintf("Bearer %s", tokenStr)
				c.Request().Header.Set("Authorization", jwtoken)
			}

			if s.NoError(s.handler.NewPocketMessage(c)) {
				body := w.Body.Bytes()

				type response struct {
					Message string `json:"message"`
				}
				var resp response
				err := json.Unmarshal(body, &resp)
				if err != nil {
					s.Error(err, "error unmarshaling")
				}

				s.Equal(v.expectStatusCode, w.Result().StatusCode)
				s.Equal(v.expectMessage, resp.Message)
			}
		})
	}
}
func (s *PocketMessageSuite) TestGetPocketMessageByRandomID() {
	testCases := []struct {
		name             string
		method           string
		path             string
		paramValue       string
		contentType      string
		expectBody       dto.MsgForPublic
		wantServiceError bool
		expectStatusCode int
		expectMessage    string
	}{
		{
			name:        "test_get_pocket_message_by_random_id-normal",
			method:      http.MethodPost,
			path:        "/api/v1/pocket-messages",
			paramValue:  "testRand",
			contentType: "application/json",
			expectBody: dto.MsgForPublic{
				Title:   "Super idol",
				Content: "Idol idol",
			},
			wantServiceError: false,
			expectStatusCode: http.StatusOK,
			expectMessage:    "success",
		},
		{
			name:             "test_get_pocket_message_by_random_id-error_random_id_empty",
			method:           http.MethodPost,
			path:             "/api/v1/pocket-messages",
			paramValue:       "",
			contentType:      "application/json",
			expectBody:       dto.MsgForPublic{},
			wantServiceError: false,
			expectStatusCode: http.StatusBadRequest,
			expectMessage:    "",
		},
		// {
		// 	name:             "test_get_pocket_message_by_random_id-error_service",
		// 	method:           http.MethodPost,
		// 	path:             "/api/v1/pocket-messages",
		// 	paramValue:       "testRand",
		// 	contentType:      "application/json",
		// 	wantServiceError: true,
		// 	expectStatusCode: http.StatusInternalServerError,
		// 	expectMessage:    "service error",
		// },
	}

	for _, v := range testCases {
		s.T().Run(v.name, func(t *testing.T) {
			if v.wantServiceError {
				s.service.
					On("GetPocketMessageByRandomID",
						mock.Anything).
					Return(dto.MsgForPublic{}, errors.New("service error"))
			} else {
				s.service.
					On("GetPocketMessageByRandomID",
						"testRand").
					Return(dto.MsgForPublic{
						Title:   "Super idol",
						Content: "Idol idol",
					}, nil)
			}

			r := httptest.NewRequest(v.method, "/", nil)
			w := httptest.NewRecorder()
			c := echo.New().NewContext(r, w)
			c.SetPath(v.path)
			c.SetParamNames("random_id")
			c.SetParamValues(v.paramValue)
			c.Request().Header.Set("Content-Type", v.contentType)

			if s.NoError(s.handler.GetPocketMessageByRandomID(c)) {
				body := w.Body.Bytes()

				type response struct {
					Message string           `json:"message"`
					Data    dto.MsgForPublic `json:"data"`
				}
				var resp response
				err := json.Unmarshal(body, &resp)
				if err != nil {
					s.Error(err, "error unmarshaling")
				}

				s.Equal(v.expectBody, resp.Data)
				s.Equal(v.expectStatusCode, w.Result().StatusCode)
				s.Equal(v.expectMessage, resp.Message)
			}
		})
	}
}

func (s *PocketMessageSuite) TestUpdatePocketMessage() {
	testCases := []struct {
		name             string
		method           string
		expectMessage    string
		uuid             string
		newBody          models.PocketMessage
		expectStatusCode int
		wantBindingError bool
	}{
		{
			name:          "update_pocket_message-normal",
			method:        http.MethodPut,
			expectMessage: "updated",
			newBody: models.PocketMessage{
				Title:   "new title",
				Content: "new content",
			},
			uuid:             uuid.Nil.String(),
			expectStatusCode: http.StatusOK,
		},
		{
			name:          "update_pocket_message-error_uuid_invalid",
			method:        http.MethodPut,
			expectMessage: "",
			newBody: models.PocketMessage{
				Title:   "new title",
				Content: "new content",
			},
			uuid:             "sdasea-ase-fa-wr",
			expectStatusCode: http.StatusBadRequest,
		},
		{
			name:          "update_pocket_message-error_binding",
			method:        http.MethodPut,
			expectMessage: "",
			newBody: models.PocketMessage{
				Title:   "new title",
				Content: "new content",
			},
			wantBindingError: true,
			uuid:             "sdasea-ase-fa-wr",
			expectStatusCode: http.StatusBadRequest,
		},
		{
			name:          "update_pocket_message-error_title_empty",
			method:        http.MethodPut,
			expectMessage: "",
			newBody: models.PocketMessage{
				Title:   "",
				Content: "new content",
			},
			uuid:             uuid.Nil.String(),
			expectStatusCode: http.StatusBadRequest,
		},
		{
			name:          "update_pocket_message-error_content_empty",
			method:        http.MethodPut,
			expectMessage: "",
			newBody: models.PocketMessage{
				Title:   "new title",
				Content: "",
			},
			uuid:             uuid.Nil.String(),
			expectStatusCode: http.StatusBadRequest,
		},
	}
	for _, v := range testCases {
		s.T().Run(v.name, func(t *testing.T) {
			s.service.On("UpdatePocketMessage", v.newBody).Return(nil)

			res, _ := json.Marshal(v.newBody)
			r := httptest.NewRequest(v.method, "/", bytes.NewBuffer(res))
			w := httptest.NewRecorder()
			c := echo.New().NewContext(r, w)
			c.SetPath("/")
			c.SetParamNames("uuid")
			c.SetParamValues(v.uuid)
			if !v.wantBindingError {
				c.Request().Header.Set("Content-Type", "application/json")
			}

			if s.NoError(s.handler.UpdatePocketMessage(c)) {
				body := w.Body.Bytes()
				type response struct {
					Message string `json:"message"`
				}
				var resp response
				err := json.Unmarshal(body, &resp)
				if err != nil {
					s.Error(err, "error unmarshaling")
				}

				s.Equal(v.expectStatusCode, w.Result().StatusCode)
				s.Equal(v.expectMessage, resp.Message)
			}
		})
	}
}

func (s *PocketMessageSuite) TestDeletePocketMessage() {
	testCases := []struct {
		name             string
		method           string
		uuid             string
		expectMessage    string
		expectStatusCode int
	}{
		{
			name:             "delete_pocket_message-normal",
			method:           http.MethodDelete,
			uuid:             uuid.Nil.String(),
			expectMessage:    "deleted",
			expectStatusCode: http.StatusOK,
		},
		{
			name:             "delete_pocket_message-error_invalid_uuid",
			method:           http.MethodDelete,
			uuid:             "asdb-ea-wdi-ba",
			expectMessage:    "",
			expectStatusCode: http.StatusBadRequest,
		},
		// {
		// 	name:             "delete_pocket_message-error_service",
		// 	method:           http.MethodDelete,
		// 	uuid:             uuid.Nil.String(),
		// 	expectMessage:    "",
		// 	expectStatusCode: http.StatusBadRequest,
		// },
	}
	for _, v := range testCases {
		s.T().Run(v.name, func(t *testing.T) {
			s.service.On("DeletePocketMessage", uuid.Nil).Return(nil)

			w := httptest.NewRecorder()
			c := echo.New().NewContext(httptest.NewRequest(v.method, "/", nil), w)
			c.SetPath("/")
			c.Request().Header.Set("Content-Type", "application/json")
			c.SetParamNames("uuid")
			c.SetParamValues(v.uuid)

			if s.NoError(s.handler.DeletePocketMessage(c)) {
				body := w.Body.Bytes()
				type response struct {
					Message string `json:"message"`
				}
				var resp response
				err := json.Unmarshal(body, &resp)
				if err != nil {
					s.Error(err, "error unmarshaling")
				}

				s.Equal(v.expectStatusCode, w.Result().StatusCode)
				s.Equal(v.expectMessage, resp.Message)
			}
		})
	}
}

func (s *PocketMessageSuite) TestGetOwnedPocketMessage() {
	testCases := []struct {
		name             string
		method           string
		expectData       dto.OwnedMessage
		expectMessage    string
		expectStatusCode int
		wantJWTError     bool
	}{
		{
			name:   "get_owned_pocket_message-normal",
			method: http.MethodGet,
			// expectData: dto.OwnedMessage{
			// 	Title: "my title",
			// },
			expectMessage:    "success",
			expectStatusCode: http.StatusOK,
		},
		{
			name:   "get_owned_pocket_message-error_decode_jwt",
			method: http.MethodGet,
			// expectData: dto.OwnedMessage{
			// 	Title: "my title",
			// },
			expectMessage:    "authorization header not found",
			wantJWTError:     true,
			expectStatusCode: http.StatusBadRequest,
		},
	}
	for _, v := range testCases {
		s.T().Run(v.name, func(t *testing.T) {
			s.service.On("GetUserPocketMessage", dto.Token{
				UUID:     uuid.Nil,
				Username: "userTest",
			}).
				Return([]dto.OwnedMessage{
					{
						Title: "my title",
					},
				}, nil)

			r := httptest.NewRequest(v.method, "/", nil)
			w := httptest.NewRecorder()
			c := echo.New().NewContext(r, w)
			c.Request().Header.Set("Content-Type", "application/json")
			c.SetPath("/")
			if !v.wantJWTError {
				tokenStr, _ := middleware.GetToken(uuid.Nil, "userTest")
				auth := fmt.Sprintf("Bearer %s", tokenStr)
				c.Request().Header.Set("Authorization", auth)
			}

			if s.NoError(s.handler.GetOwnedPocketMessage(c)) {
				body := w.Body.Bytes()
				type response struct {
					Message string             `json:"message"`
					Data    []dto.OwnedMessage `json:"data"`
				}
				var resp response
				err := json.Unmarshal(body, &resp)
				if err != nil {
					s.Error(err, "error unmarshaling")
				}

				s.Equal(v.expectMessage, resp.Message)
				s.Equal(v.expectStatusCode, w.Result().StatusCode)
			}
		})
	}
}
