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

type UserSuite struct {
	suite.Suite
	service UserServices
}

func TestSuiteUser(t *testing.T) {
	suite.Run(t, new(UserSuite))
}
func (s *UserSuite) SetupSuite() {
	service := NewUserServices(&m.MockGorm{})
	s.service = service
}
func (s *UserSuite) TearDownSuite() {}

// Signup
func (s *UserSuite) TestSignup() {
	testCase := []struct {
		name        string
		body        models.User
		method      string
		expectError error
	}{
		{
			name: "signup-normal",
			body: models.User{
				Username: "superman",
				Password: "12345678",
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

			err = s.service.SignUp(c)
			s.Equal(v.expectError, err)
		})
	}
}
func (s *UserSuite) TestSignupErrorUsernameEmpty() {
	testCase := []struct {
		name        string
		body        models.User
		method      string
		expectError error
	}{
		{
			name: "signup-error_username_empty",
			body: models.User{
				Username: "",
				Password: "12345678",
			},
			method:      http.MethodPost,
			expectError: errors.New("username should not be empty"),
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

			err = s.service.SignUp(c)
			s.Equal(v.expectError, err)
		})
	}
}
func (s *UserSuite) TestSignupErrorPasswordEmpty() {
	testCase := []struct {
		name        string
		body        models.User
		method      string
		expectError error
	}{
		{
			name: "signup-error_password_empty",
			body: models.User{
				Username: "asd",
				Password: "",
			},
			method:      http.MethodPost,
			expectError: errors.New("password should not be empty"),
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

			err = s.service.SignUp(c)
			s.Equal(v.expectError, err)
		})
	}
}
func (s *UserSuite) TestSignupErrorDB() {
	testCase := []struct {
		name        string
		body        models.User
		method      string
		expectError error
	}{
		{
			name: "signup-error_db",
			body: models.User{
				Username: "admin",
				Password: "asd",
			},
			method:      http.MethodPost,
			expectError: errors.New("username has been taken"),
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

			err = s.service.SignUp(c)
			s.Equal(v.expectError, err)
		})
	}
}

// Login
func (s *UserSuite) TestLogin() {
	testCase := []struct {
		name        string
		body        models.User
		expectBody  dto.Login
		method      string
		expectError error
	}{
		{
			name: "login-normal",
			body: models.User{
				Username: "udin",
				Password: "12345678",
			},
			expectBody: dto.Login{
				Username: "udin",
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

			result, err := s.service.Login(c)
			s.Equal(v.expectBody.Username, result.Username)
			s.Equal(v.expectError, err)
		})
	}
}
func (s *UserSuite) TestLoginErrorBinding() {
	testCase := []struct {
		name        string
		body        models.User
		expectBody  dto.Login
		method      string
		expectError error
	}{
		{
			name: "login-error_binding",
			body: models.User{
				Username: "superman",
				Password: "12345678",
			},
			expectBody: dto.Login{
				Username: "",
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

			result, err := s.service.Login(c)
			s.Equal(v.expectBody.Username, result.Username)
			s.Equal(v.expectError, err)
		})
	}
}
func (s *UserSuite) TestLoginErrorUsernameEmpty() {
	testCase := []struct {
		name        string
		body        models.User
		expectBody  dto.Login
		method      string
		expectError error
	}{
		{
			name: "login-error_username_empty",
			body: models.User{
				Username: "",
				Password: "12345678",
			},
			expectBody: dto.Login{
				Username: "",
			},
			method:      http.MethodPost,
			expectError: errors.New("username should not be empty"),
		},
	}
	for _, v := range testCase {
		s.T().Run(v.name, func(t *testing.T) {
			res, _ := json.Marshal(v.body)
			r := httptest.NewRequest(v.method, "/", bytes.NewBuffer(res))
			w := httptest.NewRecorder()
			c := echo.New().NewContext(r, w)
			c.Request().Header.Set("Content-Type", "application/json")
			result, err := s.service.Login(c)
			s.Equal(v.expectBody.Username, result.Username)
			s.Equal(v.expectError, err)
		})
	}
}
func (s *UserSuite) TestLoginErrorPasswordEmpty() {
	testCase := []struct {
		name        string
		body        models.User
		expectBody  dto.Login
		method      string
		expectError error
	}{
		{
			name: "login-error_password_empty",
			body: models.User{
				Username: "asd",
				Password: "",
			},
			expectBody: dto.Login{
				Username: "",
			},
			method:      http.MethodPost,
			expectError: errors.New("password should not be empty"),
		},
	}
	for _, v := range testCase {
		s.T().Run(v.name, func(t *testing.T) {
			res, _ := json.Marshal(v.body)
			r := httptest.NewRequest(v.method, "/", bytes.NewBuffer(res))
			w := httptest.NewRecorder()
			c := echo.New().NewContext(r, w)
			c.Request().Header.Set("Content-Type", "application/json")
			result, err := s.service.Login(c)
			s.Equal(v.expectBody.Username, result.Username)
			s.Equal(v.expectError, err)
		})
	}
}
func (s *UserSuite) TestLoginErrorPasswordDB() {
	testCase := []struct {
		name        string
		body        models.User
		expectBody  dto.Login
		method      string
		expectError error
	}{
		{
			name: "login-error_password_db",
			body: models.User{
				Username: "suneo",
				Password: "asde",
			},
			expectBody: dto.Login{
				Username: "",
			},
			method:      http.MethodPost,
			expectError: errors.New("record not found"),
		},
	}
	for _, v := range testCase {
		s.T().Run(v.name, func(t *testing.T) {
			res, _ := json.Marshal(v.body)
			r := httptest.NewRequest(v.method, "/", bytes.NewBuffer(res))
			w := httptest.NewRecorder()
			c := echo.New().NewContext(r, w)
			c.Request().Header.Set("Content-Type", "application/json")
			result, err := s.service.Login(c)
			s.Equal(v.expectBody.Username, result.Username)
			s.Equal(v.expectError, err)
		})
	}
}

// UpdateUsername
func (s *UserSuite) TestUpdateUsername() {
	testCase := []struct {
		name        string
		body        models.User
		method      string
		expectError error
	}{
		{
			name: "update_username-normal",
			body: models.User{
				Username: "udin",
				Password: "12345678",
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

			token, err := middleware.GetToken(uuid.Nil, "udin")
			if err != nil {
				s.Error(err, "error get token")
			}
			bearer := fmt.Sprintf("Bearer %s", token)
			c.Request().Header.Set("Authorization", bearer)

			err = s.service.UpdateUsername(c)
			s.Equal(v.expectError, err)
		})
	}
}
func (s *UserSuite) TestUpdateUsernameErrorBinding() {
	testCase := []struct {
		name        string
		body        models.User
		method      string
		expectError error
	}{
		{
			name: "update_username-error_binding",
			body: models.User{
				Username: "udin",
				Password: "12345678",
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

			token, err := middleware.GetToken(uuid.Nil, "udin")
			if err != nil {
				s.Error(err, "error get token")
			}
			bearer := fmt.Sprintf("Bearer %s", token)
			c.Request().Header.Set("Authorization", bearer)

			err = s.service.UpdateUsername(c)
			s.Equal(v.expectError, err)
		})
	}
}
func (s *UserSuite) TestUpdateUsernameErrorUsernameEmpty() {
	testCase := []struct {
		name        string
		body        models.User
		method      string
		expectError error
	}{
		{
			name: "update_username-error_username_empty",
			body: models.User{
				Username: "",
				Password: "12345678",
			},
			method:      http.MethodPost,
			expectError: errors.New("error, username should not be empty"),
		},
	}
	for _, v := range testCase {
		s.T().Run(v.name, func(t *testing.T) {
			res, _ := json.Marshal(v.body)
			r := httptest.NewRequest(v.method, "/", bytes.NewBuffer(res))
			w := httptest.NewRecorder()
			c := echo.New().NewContext(r, w)
			c.Request().Header.Set("Content-Type", "application/json")

			token, err := middleware.GetToken(uuid.Nil, "udin")
			if err != nil {
				s.Error(err, "error get token")
			}
			bearer := fmt.Sprintf("Bearer %s", token)
			c.Request().Header.Set("Authorization", bearer)

			err = s.service.UpdateUsername(c)
			s.Equal(v.expectError, err)
		})
	}
}
func (s *UserSuite) TestUpdateUsernameErrorAuth() {
	testCase := []struct {
		name        string
		body        models.User
		method      string
		expectError error
	}{
		{
			name: "update_username-error_auth",
			body: models.User{
				Username: "aseasd",
				Password: "12345678",
			},
			method:      http.MethodPost,
			expectError: errors.New("authorization header not found"),
		},
	}
	for _, v := range testCase {
		s.T().Run(v.name, func(t *testing.T) {
			res, _ := json.Marshal(v.body)
			r := httptest.NewRequest(v.method, "/", bytes.NewBuffer(res))
			w := httptest.NewRecorder()
			c := echo.New().NewContext(r, w)
			c.Request().Header.Set("Content-Type", "application/json")

			err := s.service.UpdateUsername(c)
			s.Equal(v.expectError, err)
		})
	}
}
func (s *UserSuite) TestUpdateUsernameErrorDB() {
	testCase := []struct {
		name        string
		body        models.User
		method      string
		expectError error
	}{
		{
			name: "update_username-error_db",
			body: models.User{
				Username: "suneo",
				Password: "12345678",
			},
			method:      http.MethodPost,
			expectError: errors.New("username has taken"),
		},
	}
	for _, v := range testCase {
		s.T().Run(v.name, func(t *testing.T) {
			res, _ := json.Marshal(v.body)
			r := httptest.NewRequest(v.method, "/", bytes.NewBuffer(res))
			w := httptest.NewRecorder()
			c := echo.New().NewContext(r, w)
			c.Request().Header.Set("Content-Type", "application/json")

			token, err := middleware.GetToken(uuid.Nil, "udin")
			if err != nil {
				s.Error(err, "error get token")
			}
			bearer := fmt.Sprintf("Bearer %s", token)
			c.Request().Header.Set("Authorization", bearer)

			err = s.service.UpdateUsername(c)
			s.Equal(v.expectError, err)
		})
	}
}

// UpdatePassword
func (s *UserSuite) TestUpdatePassword() {
	testCase := []struct {
		name        string
		body        models.User
		method      string
		expectError error
	}{
		{
			name: "update_password-normal",
			body: models.User{
				Username: "udin",
				Password: "12345678",
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

			err := s.service.UpdatePassword(c)
			s.Equal(v.expectError, err)
		})
	}
}
func (s *UserSuite) TestUpdatePasswordErrorBinding() {
	testCase := []struct {
		name        string
		body        models.User
		method      string
		expectError error
	}{
		{
			name: "update_password-error_binding",
			body: models.User{
				Username: "udin",
				Password: "12345678",
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

			err := s.service.UpdatePassword(c)
			s.Equal(v.expectError, err)
		})
	}
}
func (s *UserSuite) TestUpdatePasswordErrorUsernameEmpty() {
	testCase := []struct {
		name        string
		body        models.User
		method      string
		expectError error
	}{
		{
			name: "update_password-error_username_empty",
			body: models.User{
				Username: "",
				Password: "12345678",
			},
			method:      http.MethodPost,
			expectError: errors.New("username should not be empty"),
		},
	}
	for _, v := range testCase {
		s.T().Run(v.name, func(t *testing.T) {
			res, _ := json.Marshal(v.body)
			r := httptest.NewRequest(v.method, "/", bytes.NewBuffer(res))
			w := httptest.NewRecorder()
			c := echo.New().NewContext(r, w)
			c.Request().Header.Set("Content-Type", "application/json")

			err := s.service.UpdatePassword(c)
			s.Equal(v.expectError, err)
		})
	}
}
func (s *UserSuite) TestUpdatePasswordErrorPasswordEmpty() {
	testCase := []struct {
		name        string
		body        models.User
		method      string
		expectError error
	}{
		{
			name: "update_password-error_password_empty",
			body: models.User{
				Username: "asds",
				Password: "",
			},
			method:      http.MethodPost,
			expectError: errors.New("password should not be empty"),
		},
	}
	for _, v := range testCase {
		s.T().Run(v.name, func(t *testing.T) {
			res, _ := json.Marshal(v.body)
			r := httptest.NewRequest(v.method, "/", bytes.NewBuffer(res))
			w := httptest.NewRecorder()
			c := echo.New().NewContext(r, w)
			c.Request().Header.Set("Content-Type", "application/json")

			err := s.service.UpdatePassword(c)
			s.Equal(v.expectError, err)
		})
	}
}
func (s *UserSuite) TestUpdatePasswordErrorDB() {
	testCase := []struct {
		name        string
		body        models.User
		method      string
		expectError error
	}{
		{
			name: "update_password-error_db",
			body: models.User{
				Username: "asds",
				Password: "adwawea",
			},
			method:      http.MethodPost,
			expectError: errors.New("password too short"),
		},
	}
	for _, v := range testCase {
		s.T().Run(v.name, func(t *testing.T) {
			res, _ := json.Marshal(v.body)
			r := httptest.NewRequest(v.method, "/", bytes.NewBuffer(res))
			w := httptest.NewRecorder()
			c := echo.New().NewContext(r, w)
			c.Request().Header.Set("Content-Type", "application/json")

			err := s.service.UpdatePassword(c)
			s.Equal(v.expectError, err)
		})
	}
}
