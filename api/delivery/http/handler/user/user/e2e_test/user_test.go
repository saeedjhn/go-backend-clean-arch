//go:build e2e
// +build e2e

package e2e_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/saeedjhn/go-backend-clean-arch/internal/sharedkernel/models"

	userdto "github.com/saeedjhn/go-backend-clean-arch/internal/dto/user"

	setuptest "github.com/saeedjhn/go-backend-clean-arch/test/setup_test"

	"github.com/stretchr/testify/suite"
)

const (
	_baseURL         = "http://localhost:8000"
	_correctMobile   = "09120000001"
	_correctPassword = "password123"
)

//go:generate go test -v -count=1 -race -tags=e2e -run UserTestSuite

type UserTestSuite struct {
	suite.Suite
	baseURL     string
	tinyRequest *setuptest.TinyRequest
}

func Test_UserTestSuite(t *testing.T) {
	suite.Run(t, new(UserTestSuite))
}

// SetupTest runs before each test in the suite.
func (s *UserTestSuite) SetupTest() {
	s.baseURL = _baseURL
	s.tinyRequest = setuptest.NewTinyRequest()
}

func (s *UserTestSuite) Test_UserLogin() {
	s.T().Run("UserLogin_Success", func(_ *testing.T) {
		rl := userdto.LoginRequest{
			Mobile:   _correctMobile,
			Password: _correctPassword,
		}
		setBodyErr := s.tinyRequest.SetBody(rl)
		s.Require().NoError(setBodyErr)

		s.tinyRequest.SetHeader(map[string]string{
			"Content-Type": "application/json",
			// "Authorization": fmt.Sprintf("Bearer %s"),
		})

		req, err := s.tinyRequest.Fetch(
			http.MethodPost,
			fmt.Sprintf("%s/users/auth/login", s.baseURL),
		)

		s.Require().NoError(err)
		s.Require().Equal(http.StatusOK, req.StatusCode())

		var resp models.SuccessResponse[userdto.LoginResponse]

		unmarshalErr := req.UnmarshallBody(&resp)
		s.Require().NoError(unmarshalErr)

		s.Require().NotZero(resp.Data.UserInfo.ID)
		s.Require().Equal(_correctMobile, resp.Data.UserInfo.Mobile)
		s.Require().NotEmpty(resp.Data.Tokens)
		s.Require().Nil(resp.Data.FieldErrors)
	})

	s.T().Run("UserLogin_FailureUserNotFound", func(_ *testing.T) {
		rl := userdto.LoginRequest{
			Mobile:   "09000000000",
			Password: _correctPassword,
		}
		setBodyErr := s.tinyRequest.SetBody(rl)
		s.Require().NoError(setBodyErr)

		s.tinyRequest.SetHeader(map[string]string{
			"Content-Type": "application/json",
			// "Authorization": fmt.Sprintf("Bearer %s"),
		})

		req, err := s.tinyRequest.Fetch(
			http.MethodPost,
			fmt.Sprintf("%s/users/auth/login", s.baseURL),
		)

		s.Require().NoError(err)
		s.Require().Equal(http.StatusNotFound, req.StatusCode())
	})

	s.T().Run("Failure_FailureBadRequest", func(_ *testing.T) {
		rl := userdto.LoginRequest{
			Mobile:   "",
			Password: "",
		}
		setBodyErr := s.tinyRequest.SetBody(rl)
		s.Require().NoError(setBodyErr)

		s.tinyRequest.SetHeader(map[string]string{
			"Content-Type": "application/json",
			// "Authorization": fmt.Sprintf("Bearer %s"),
		})

		req, err := s.tinyRequest.Fetch(
			http.MethodPost,
			fmt.Sprintf("%s/users/auth/login", s.baseURL),
		)

		s.Require().NoError(err)
		s.Require().Equal(http.StatusUnprocessableEntity, req.StatusCode())
	})
}

func (s *UserTestSuite) Test_UserRegister() {

}
