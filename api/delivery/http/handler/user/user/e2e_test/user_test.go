//go:build e2e
// +build e2e

package e2e_test

import (
	"fmt"
	"net/http"
	"testing"

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
func (suite *UserTestSuite) SetupTest() {
	suite.baseURL = _baseURL
	suite.tinyRequest = setuptest.NewTinyRequest()
}

func (suite *UserTestSuite) Test_UserLogin() {
	suite.T().Run("UserLogin_Success", func(_ *testing.T) {
		rl := userdto.LoginRequest{
			Mobile:   _correctMobile,
			Password: _correctPassword,
		}
		setBodyErr := suite.tinyRequest.SetBody(rl)
		suite.Require().NoError(setBodyErr)

		suite.tinyRequest.SetHeader(map[string]string{
			"Content-Type": "application/json",
			// "Authorization": fmt.Sprintf("Bearer %s"),
		})

		req, err := suite.tinyRequest.Fetch(
			http.MethodPost,
			fmt.Sprintf("%s/users/auth/login", suite.baseURL),
		)

		suite.Require().NoError(err)
		suite.Require().Equal(http.StatusOK, req.StatusCode())

		var resp userdto.LoginResponse
		unmarshalErr := req.UnmarshallBody(&resp)
		suite.Require().NoError(unmarshalErr)

		suite.Require().NotZero(resp.Data.ID)
		suite.Require().Equal(_correctMobile, resp.Data.Mobile)
		suite.Require().NotEmpty(resp.Tokens)
		suite.Require().Nil(resp.FieldErrors)
	})

	suite.T().Run("UserLogin_FailureUserNotFound", func(_ *testing.T) {
		rl := userdto.LoginRequest{
			Mobile:   "09123456789",
			Password: _correctPassword,
		}
		setBodyErr := suite.tinyRequest.SetBody(rl)
		suite.Require().NoError(setBodyErr)

		suite.tinyRequest.SetHeader(map[string]string{
			"Content-Type": "application/json",
			// "Authorization": fmt.Sprintf("Bearer %s"),
		})

		req, err := suite.tinyRequest.Fetch(
			http.MethodPost,
			fmt.Sprintf("%s/users/auth/login", suite.baseURL),
		)

		suite.Require().NoError(err)
		suite.Require().Equal(http.StatusNotFound, req.StatusCode())
	})

	suite.T().Run("Failure_FailureBadRequest", func(_ *testing.T) {
		rl := userdto.LoginRequest{
			Mobile:   "",
			Password: "",
		}
		setBodyErr := suite.tinyRequest.SetBody(rl)
		suite.Require().NoError(setBodyErr)

		suite.tinyRequest.SetHeader(map[string]string{
			"Content-Type": "application/json",
			// "Authorization": fmt.Sprintf("Bearer %s"),
		})

		req, err := suite.tinyRequest.Fetch(
			http.MethodPost,
			fmt.Sprintf("%s/users/auth/login", suite.baseURL),
		)

		suite.Require().NoError(err)
		suite.Require().Equal(http.StatusUnprocessableEntity, req.StatusCode())
	})
}

func (suite *UserTestSuite) Test_UserRegister() {

}
