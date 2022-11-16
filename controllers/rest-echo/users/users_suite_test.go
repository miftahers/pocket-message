package users

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"pocket-message/configs"
	"pocket-message/dto"
	"pocket-message/middleware"
	"pocket-message/models"
	"pocket-message/services/users/mocks"
	"testing"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"
)

type UserSuite struct {
	suite.Suite
	service *mocks.IUserServices
	handler UserHandler
}

func (s *UserSuite) SetupSuite() {
	service := new(mocks.IUserServices)
	handler := UserHandler{
		IUserServices: service,
	}

	s.service = service
	s.handler = handler

	// Init Config
	cfg := &configs.Config{}
	cfg.TokenSecret = "test"
	configs.Cfg = cfg
}

func TestUserSuite(t *testing.T) {
	suite.Run(t, new(UserSuite))
}

func (s *UserSuite) TestSignup() {
	testCases := []struct {
		name             string
		method           string
		User             models.User
		expectStatusCode int
		expectMessage    string
		wantErrorBinding bool
	}{
		{
			name: "signup-normal",
			User: models.User{
				Username: "user",
				Password: "password",
			},
			method:           http.MethodPost,
			expectStatusCode: http.StatusCreated,
			expectMessage:    "created",
		},
		{
			name: "signup-error_binding",
			User: models.User{
				Username: "user",
				Password: "password",
			},
			wantErrorBinding: true,
			method:           http.MethodPost,
			expectStatusCode: http.StatusBadRequest,
			expectMessage:    "code=415, message=Unsupported Media Type",
		},
		{
			name: "signup-error_password_empty",
			User: models.User{
				Username: "user",
				Password: "",
			},
			method:           http.MethodPost,
			expectStatusCode: http.StatusBadRequest,
			expectMessage:    "",
		},
		{
			name: "signup-error_username_empty",
			User: models.User{
				Username: "",
				Password: "password",
			},
			method:           http.MethodPost,
			expectStatusCode: http.StatusBadRequest,
			expectMessage:    "",
		},
	}
	for _, v := range testCases {
		s.T().Run(v.name, func(t *testing.T) {
			s.service.On("SignUp", v.User).Return(nil)

			res, _ := json.Marshal(v.User)
			r := httptest.NewRequest(v.method, "/", bytes.NewBuffer(res))
			w := httptest.NewRecorder()
			c := echo.New().NewContext(r, w)
			c.SetPath("/")
			if !v.wantErrorBinding {
				c.Request().Header.Set("Content-Type", "application/json")
			}

			if s.NoError(s.handler.SignUp(c)) {
				body := w.Body.Bytes()
				type response struct {
					Message string `json:"message"`
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

func (s *UserSuite) TestLogin() {
	testCases := []struct {
		name             string
		method           string
		newUser          models.User
		expectStatusCode int
		expectMessage    string
		wantErrorBinding bool
	}{
		{
			name: "login-normal",
			newUser: models.User{
				Username: "user",
				Password: "password",
			},
			method:           http.MethodPost,
			expectStatusCode: http.StatusOK,
			expectMessage:    "success",
		},
		{
			name: "login-error_binding",
			newUser: models.User{
				Username: "user",
				Password: "password",
			},
			wantErrorBinding: true,
			method:           http.MethodPost,
			expectStatusCode: http.StatusBadRequest,
			expectMessage:    "code=415, message=Unsupported Media Type",
		},
		{
			name: "login-error_password_empty",
			newUser: models.User{
				Username: "user",
				Password: "",
			},
			method:           http.MethodPost,
			expectStatusCode: http.StatusBadRequest,
			expectMessage:    "",
		},
		{
			name: "login-error_username_empty",
			newUser: models.User{
				Username: "",
				Password: "password",
			},
			method:           http.MethodPost,
			expectStatusCode: http.StatusBadRequest,
			expectMessage:    "",
		},
	}
	for _, v := range testCases {
		s.T().Run(v.name, func(t *testing.T) {
			s.service.On("Login", v.newUser).Return(dto.Login{}, nil)

			res, _ := json.Marshal(v.newUser)
			r := httptest.NewRequest(v.method, "/", bytes.NewBuffer(res))
			w := httptest.NewRecorder()
			c := echo.New().NewContext(r, w)
			c.SetPath("/")
			if !v.wantErrorBinding {
				c.Request().Header.Set("Content-Type", "application/json")
			}

			if s.NoError(s.handler.Login(c)) {
				body := w.Body.Bytes()
				type response struct {
					Message string `json:"message"`
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

func (s *UserSuite) TestUpdateUsername() {
	testCases := []struct {
		name             string
		method           string
		User             models.User
		expectStatusCode int
		expectMessage    string
		wantErrorAuth    bool
		wantErrorBinding bool
	}{
		{
			name: "update_username-normal",
			User: models.User{
				Username: "user",
				Password: "password",
			},
			method:           http.MethodPut,
			expectStatusCode: http.StatusOK,
			expectMessage:    "success",
		},
		{
			name: "update_username-error_binding",
			User: models.User{
				Username: "user",
				Password: "password",
			},
			wantErrorBinding: true,
			method:           http.MethodPut,
			expectStatusCode: http.StatusBadRequest,
			expectMessage:    "code=415, message=Unsupported Media Type",
		},
		{
			name: "update_username-error_username_empty",
			User: models.User{
				Username: "",
				Password: "password",
			},
			method:           http.MethodPut,
			expectStatusCode: http.StatusBadRequest,
			expectMessage:    "",
		},
		{
			name: "update_username-error_auth",
			User: models.User{
				Username: "user",
				Password: "password",
			},
			method:           http.MethodPut,
			expectStatusCode: http.StatusBadRequest,
			wantErrorAuth:    true,
			expectMessage:    "authorization header not found",
		},
	}
	for _, v := range testCases {
		s.T().Run(v.name, func(t *testing.T) {
			s.service.On("UpdateUsername", v.User, dto.Token{UUID: uuid.Nil, Username: "user"}).Return(nil)

			res, _ := json.Marshal(v.User)
			r := httptest.NewRequest(v.method, "/", bytes.NewBuffer(res))
			w := httptest.NewRecorder()
			c := echo.New().NewContext(r, w)
			c.SetPath("/")
			if !v.wantErrorBinding {
				c.Request().Header.Set("Content-Type", "application/json")
			}
			if !v.wantErrorAuth {
				tokenStr, _ := middleware.GetToken(uuid.Nil, "user")
				authStr := fmt.Sprintf("Bearer %s", tokenStr)
				c.Request().Header.Set("Authorization", authStr)
			}

			if s.NoError(s.handler.UpdateUsername(c)) {
				body := w.Body.Bytes()
				type response struct {
					Message string `json:"message"`
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

func (s *UserSuite) TestUpdatePassword() {
	testCases := []struct {
		name             string
		method           string
		User             models.User
		expectStatusCode int
		expectMessage    string
		wantErrorBinding bool
	}{
		{
			name: "update_password-normal",
			User: models.User{
				Username: "user",
				Password: "password",
			},
			method:           http.MethodPut,
			expectStatusCode: http.StatusOK,
			expectMessage:    "success",
		},
		{
			name: "update_password-error_binding",
			User: models.User{
				Username: "user",
				Password: "password",
			},
			wantErrorBinding: true,
			method:           http.MethodPut,
			expectStatusCode: http.StatusBadRequest,
			expectMessage:    "code=415, message=Unsupported Media Type",
		},
		{
			name: "update_password-error_username_empty",
			User: models.User{
				Username: "",
				Password: "password",
			},
			method:           http.MethodPut,
			expectStatusCode: http.StatusBadRequest,
			expectMessage:    "",
		},
		{
			name: "update_password-error_password_empty",
			User: models.User{
				Username: "user",
				Password: "",
			},
			method:           http.MethodPut,
			expectStatusCode: http.StatusBadRequest,
			expectMessage:    "",
		},
	}
	for _, v := range testCases {
		s.T().Run(v.name, func(t *testing.T) {
			s.service.On("UpdatePassword", v.User).Return(nil)

			res, _ := json.Marshal(v.User)
			r := httptest.NewRequest(v.method, "/", bytes.NewBuffer(res))
			w := httptest.NewRecorder()
			c := echo.New().NewContext(r, w)
			c.SetPath("/")
			if !v.wantErrorBinding {
				c.Request().Header.Set("Content-Type", "application/json")
			}

			if s.NoError(s.handler.UpdatePassword(c)) {
				body := w.Body.Bytes()
				type response struct {
					Message string `json:"message"`
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
