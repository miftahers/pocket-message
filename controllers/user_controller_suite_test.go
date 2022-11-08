package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	m "pocket-message/controllers/mock"
	"pocket-message/dto"
	"pocket-message/models"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"
)

type UserSuite struct {
	suite.Suite
	handler UserHandler
}

func (s *UserSuite) SetupSuite() {

	handler := NewUserHandler(&m.MockUserServices{})
	s.handler = handler
}
func (s *UserSuite) TearDownSuite() {

}
func TestSuiteUser(t *testing.T) {
	suite.Run(t, new(UserSuite))
}

func (s *UserSuite) TestSignup() {
	testCase := []struct {
		name          string
		method        string
		path          string
		body          models.User
		expectCode    int
		expectMessage string
	}{
		{
			name:   "signup-normal",
			method: http.MethodPost,
			path:   "/api/v1/signup",
			body: models.User{
				Username: "miftah",
				Password: "test",
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

			if s.NoError(s.handler.SignUp(c)) {
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
func (s *UserSuite) TestSignupError() {
	testCase := []struct {
		name          string
		method        string
		path          string
		body          models.User
		expectCode    int
		expectMessage string
	}{
		{
			name:   "signup-error",
			method: http.MethodPost,
			path:   "/api/v1/signup",
			body: models.User{
				Username: "miftah",
			},
			expectCode:    http.StatusInternalServerError,
			expectMessage: "error, password should not be empty",
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

			if s.NoError(s.handler.SignUp(c)) {
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

func (s *UserSuite) TestLogin() {
	testCase := []struct {
		name          string
		method        string
		path          string
		body          models.User
		expectBody    dto.Login
		expectCode    int
		expectMessage string
	}{
		{
			name:   "login-normal",
			method: http.MethodPost,
			path:   "/api/v1/login",
			body: models.User{
				Username: "Super",
				Password: "test",
			},
			expectBody: dto.Login{
				Username: "Super",
				Token:    "Idol",
			},
			expectCode:    http.StatusOK,
			expectMessage: "success",
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

			if s.NoError(s.handler.Login(c)) {
				body := w.Body.Bytes()

				type response struct {
					Message string    `json:"message"`
					Data    dto.Login `json:"data"`
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
func (s *UserSuite) TestLoginError() {
	testCase := []struct {
		name          string
		method        string
		path          string
		body          models.User
		expectBody    dto.Login
		expectCode    int
		expectMessage string
	}{
		{
			name:   "login-error",
			method: http.MethodPost,
			path:   "/api/v1/login",
			body: models.User{
				Username: "",
				Password: "test",
			},
			expectBody:    dto.Login{},
			expectCode:    http.StatusInternalServerError,
			expectMessage: "error, username should not be empty",
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

			if s.NoError(s.handler.Login(c)) {
				body := w.Body.Bytes()

				type response struct {
					Message string    `json:"message"`
					Data    dto.Login `json:"data"`
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

func (s *UserSuite) TestUpdateUsername() {
	testCase := []struct {
		name          string
		method        string
		path          string
		body          models.User
		expectCode    int
		expectMessage string
	}{
		{
			name:   "update_username-normal",
			method: http.MethodPut,
			path:   "/api/v1/users/change-username",
			body: models.User{
				Username: "Super",
				Password: "test",
			},
			expectCode:    http.StatusOK,
			expectMessage: "success",
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

			if s.NoError(s.handler.UpdateUsername(c)) {
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
func (s *UserSuite) TestUpdateUsernameError() {
	testCase := []struct {
		name          string
		method        string
		path          string
		body          models.User
		expectCode    int
		expectMessage string
	}{
		{
			name:   "update_username-error",
			method: http.MethodPut,
			path:   "/api/v1/users/change-username",
			body: models.User{
				Username: "",
				Password: "test",
			},
			expectCode:    http.StatusInternalServerError,
			expectMessage: "error, username should not be empty",
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

			if s.NoError(s.handler.UpdateUsername(c)) {
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
func (s *UserSuite) TestUpdatePassword() {
	testCase := []struct {
		name          string
		method        string
		path          string
		body          models.User
		expectCode    int
		expectMessage string
	}{
		{
			name:   "update_password-normal",
			method: http.MethodPut,
			path:   "/api/v1/users/reset-password",
			body: models.User{
				Username: "Super",
				Password: "test",
			},
			expectCode:    http.StatusOK,
			expectMessage: "success",
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

			if s.NoError(s.handler.UpdatePassword(c)) {
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
func (s *UserSuite) TestUpdatePasswordError() {
	testCase := []struct {
		name          string
		method        string
		path          string
		body          models.User
		expectCode    int
		expectMessage string
	}{
		{
			name:   "update_password-error",
			method: http.MethodPut,
			path:   "/api/v1/users/reset-password",
			body: models.User{
				Username: "",
				Password: "asdw",
			},
			expectCode:    http.StatusInternalServerError,
			expectMessage: "error, username should not be empty",
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

			if s.NoError(s.handler.UpdatePassword(c)) {
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
