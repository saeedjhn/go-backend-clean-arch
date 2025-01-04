package user_test

import (
	"context"
	"testing"
	"time"

	userdto "github.com/saeedjhn/go-backend-clean-arch/internal/dto/user"
	"github.com/saeedjhn/go-backend-clean-arch/internal/entity"
	userusecase "github.com/saeedjhn/go-backend-clean-arch/internal/usecase/user"
	"github.com/saeedjhn/go-backend-clean-arch/internal/usecase/user/user_test/doubles"
	"github.com/saeedjhn/go-backend-clean-arch/internal/usecase/user/user_test/mocks"
	"github.com/saeedjhn/go-backend-clean-arch/pkg/security/bcrypt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	password    = "password123"
	genPassword = func() string {
		gp, _ := userusecase.GenerateHash(password)

		return gp
	}()
	correctMobile     = "09123456789"
	incorrectMobile   = "incorrect-mobile"
	nonExistentMobile = "123456789"
)

//go:generate go test -v -count=1 -race -run Test_UserInterator_Login

//go:generate go test -v -count=1 -run Test_UserInterator_Login_ValidationSection
func Test_UserInterator_Login_ValidationSection(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	config, err := getConfig()
	if err != nil {
		t.Fatalf("Error loading configuration: %v", err)
	}

	testCases := []struct {
		name           string
		req            userdto.LoginRequest
		fieldsErr      map[string]string
		expectedError  error
		expectedResult userdto.LoginResponse
	}{{
		name:           "EmptyMobile_InvalidInputReturn",
		req:            userdto.LoginRequest{Mobile: "", Password: "password123"},
		fieldsErr:      map[string]string{"mobile": "required"},
		expectedError:  errInvalidInput,
		expectedResult: userdto.LoginResponse{FieldErrors: map[string]string{"mobile": "required"}},
	}, {
		name:           "EmptyPassword_InvalidInputReturn",
		req:            userdto.LoginRequest{Mobile: correctMobile, Password: ""},
		fieldsErr:      map[string]string{"password": "required"},
		expectedError:  errInvalidInput,
		expectedResult: userdto.LoginResponse{FieldErrors: map[string]string{"password": "required"}},
	}, {
		name:           "EmptyEmailAndPassword_InvalidInputReturn",
		req:            userdto.LoginRequest{Mobile: "", Password: ""},
		fieldsErr:      map[string]string{"mobile": "required", "password": "required"},
		expectedError:  errInvalidInput,
		expectedResult: userdto.LoginResponse{FieldErrors: map[string]string{"mobile": "required", "password": "required"}},
	}, {
		name:           "InvalidMobileFormat_InvalidInputReturn",
		req:            userdto.LoginRequest{Mobile: incorrectMobile, Password: password},
		fieldsErr:      map[string]string{"mobile": "invalid format"},
		expectedError:  errInvalidInput,
		expectedResult: userdto.LoginResponse{FieldErrors: map[string]string{"mobile": "invalid format"}},
	}, {
		name:          "ValidInput_NullReturn",
		req:           userdto.LoginRequest{Mobile: correctMobile, Password: password},
		fieldsErr:     nil,
		expectedError: nil,
		expectedResult: userdto.LoginResponse{
			Data: userdto.Data{
				ID:        1,
				Name:      "",
				Mobile:    correctMobile,
				Email:     "",
				CreatedAt: time.Time{},
				UpdatedAt: time.Time{},
			},
			Tokens: userdto.Tokens{
				AccessToken:  "access-token",
				RefreshToken: "refresh-token",
			},
			FieldErrors: nil,
		},
	}}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			mockVld, mockRepo, mockAuth, mockCache := setupMock(t)

			mockVld.EXPECT().ValidateLoginRequest(tc.req).Return(tc.fieldsErr, tc.expectedError)

			if tc.expectedError == nil {
				generatePassword, _ := bcrypt.Generate(tc.req.Password, bcrypt.DefaultCost)

				mockRepo.On("GetByMobile", ctx, tc.req.Mobile).Return(entity.User{
					ID:       1,
					Mobile:   tc.req.Mobile,
					Password: generatePassword,
				}, nil)

				userFetch, _ := mockRepo.GetByMobile(ctx, tc.req.Mobile)

				authenticable := entity.Authenticable{ID: userFetch.ID}

				mockAuth.On("CreateAccessToken", authenticable).Return("access-token", nil)
				mockAuth.On("CreateRefreshToken", authenticable).Return("refresh-token", nil)
			}

			usecase := userusecase.New(
				config,
				setupTracer(),
				mockVld,
				mockAuth,
				mockCache,
				mockRepo,
			)

			// Call the Login method
			resp, errL := usecase.Login(ctx, tc.req)

			assert.Equal(t, tc.expectedResult, resp)

			if tc.expectedError != nil {
				assert.Equal(t, tc.expectedError, errL)
				assert.Equal(t, tc.fieldsErr, resp.FieldErrors)
				assert.ObjectsAreEqualValues(tc.expectedResult, resp)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

//go:generate go test -v -count=1 -run Test_UserInterator_LoginRepositorySection
func Test_UserInterator_LoginRepositorySection(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	config, err := getConfig()
	if err != nil {
		t.Fatalf("Error loading configuration: %v", err)
	}

	testCases := []struct {
		name string
		req  userdto.LoginRequest
		repo struct {
			user entity.User
			err  error
		}
		expectedError  error
		expectedResult userdto.LoginResponse
	}{{
		name: "Login_WithNonExistentMobile_ReturnsUserNotFoundError",
		req:  userdto.LoginRequest{Mobile: nonExistentMobile, Password: correctMobile},
		repo: struct {
			user entity.User
			err  error
		}{user: entity.User{}, err: errUserNotFound},
		expectedError:  errUserNotFound,
		expectedResult: userdto.LoginResponse{},
	}, {
		name: "Login_WithRepositoryError_ReturnsUnexpectedError",
		req:  userdto.LoginRequest{Mobile: correctMobile, Password: password},
		repo: struct {
			user entity.User
			err  error
		}{user: entity.User{}, err: errUnexpected},
		expectedError:  errUnexpected,
		expectedResult: userdto.LoginResponse{},
	}, {
		name: "Login_WithValidMobile_ReturnsTokensAndUserData",
		req:  userdto.LoginRequest{Mobile: correctMobile, Password: password},
		repo: struct {
			user entity.User
			err  error
		}{user: entity.User{
			ID:        1,
			Name:      "",
			Mobile:    correctMobile,
			Email:     "",
			Password:  genPassword,
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
		}, err: nil},
		expectedError: nil,
		expectedResult: userdto.LoginResponse{
			Data: userdto.Data{
				ID:        1,
				Name:      "",
				Mobile:    correctMobile,
				Email:     "",
				CreatedAt: time.Time{},
				UpdatedAt: time.Time{},
			},
			Tokens: userdto.Tokens{
				AccessToken:  "access-token",
				RefreshToken: "refresh-token",
			},
			FieldErrors: nil,
		},
	}}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			mockVld, mockRepo, mockAuth, mockCache := setupMock(t)

			mockVld.EXPECT().ValidateLoginRequest(tc.req).Return(nil, nil)

			mockRepo.On("GetByMobile", ctx, tc.req.Mobile).
				Return(tc.repo.user, tc.repo.err)

			userFetch, _ := mockRepo.GetByMobile(ctx, tc.req.Mobile)

			if tc.expectedError == nil {
				authenticable := entity.Authenticable{ID: userFetch.ID}

				mockAuth.On("CreateAccessToken", authenticable).Return("access-token", nil)
				mockAuth.On("CreateRefreshToken", authenticable).Return("refresh-token", nil)
			}

			usecase := userusecase.New(
				config,
				setupTracer(),
				mockVld,
				mockAuth,
				mockCache,
				mockRepo,
			)

			resp, errL := usecase.Login(ctx, tc.req)

			if tc.expectedError != nil {
				require.Error(t, errL)
				assert.Equal(t, tc.expectedResult, resp)
				// assert.Equal(t, tc.expectedError, errL)
				// assert.Empty(t, resp)
			} else {
				require.NoError(t, errL)
				assert.Equal(t, tc.expectedResult, resp)
			}
		})
	}
}

//go:generate go test -v -count=1 -run Test_UserInterator_LoginCreateTokenSection
func Test_UserInterator_LoginCreateTokenSection(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	config, err := getConfig()
	if err != nil {
		t.Fatalf("Error loading configuration: %v", err)
	}

	testCases := []struct {
		name           string
		user           entity.User
		req            userdto.LoginRequest
		accessToken    string
		refreshToken   string
		expectedError  error
		expectedResult userdto.LoginResponse
	}{
		{
			name: "Login_WithNotValidRequest_TokenNotGenerated",
			user: entity.User{
				ID:        1,
				Name:      "",
				Mobile:    correctMobile,
				Email:     "",
				Password:  genPassword,
				CreatedAt: time.Time{},
				UpdatedAt: time.Time{},
			},
			req:            userdto.LoginRequest{Mobile: correctMobile, Password: password},
			accessToken:    "",
			refreshToken:   "",
			expectedError:  errAccessTokenCreationFailed,
			expectedResult: userdto.LoginResponse{},
		},
		{
			name: "Login_WithValidRequest_TokenGenerated",
			user: entity.User{
				ID:        1,
				Name:      "",
				Mobile:    "09123456789",
				Email:     "",
				Password:  genPassword,
				CreatedAt: time.Time{},
				UpdatedAt: time.Time{},
			},
			req:           userdto.LoginRequest{Mobile: correctMobile, Password: password},
			accessToken:   "access-token",
			refreshToken:  "refresh-token",
			expectedError: nil,
			expectedResult: userdto.LoginResponse{
				Data: userdto.Data{
					ID:        1,
					Name:      "",
					Mobile:    correctMobile,
					Email:     "",
					CreatedAt: time.Time{},
					UpdatedAt: time.Time{},
				},
				Tokens: userdto.Tokens{
					AccessToken:  "access-token",
					RefreshToken: "refresh-token",
				},
				FieldErrors: nil,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			mockVld, mockRepo, mockAuth, mockCache := setupMock(t)

			mockVld.EXPECT().ValidateLoginRequest(tc.req).Return(nil, nil)

			mockRepo.On("GetByMobile", ctx, tc.req.Mobile).Return(tc.user, nil)
			userFetch, _ := mockRepo.GetByMobile(ctx, tc.req.Mobile)

			authenticable := entity.Authenticable{ID: userFetch.ID}

			mockAuth.On("CreateAccessToken", authenticable).Return(tc.accessToken, tc.expectedError)
			if tc.expectedError == nil {
				mockAuth.On("CreateRefreshToken", authenticable).Return(tc.refreshToken, tc.expectedError)
			}

			usecase := userusecase.New(
				config,
				setupTracer(),
				mockVld,
				mockAuth,
				mockCache,
				mockRepo,
			)

			resp, errL := usecase.Login(ctx, tc.req)

			if tc.expectedError != nil {
				require.Error(t, errL)
				assert.Equal(t, tc.expectedResult, resp)
				// assert.Equal(t, tc.expectedError, errL)
				// assert.Empty(t, resp)
			} else {
				require.NoError(t, errL)
				assert.Equal(t, tc.expectedResult, resp)
			}
		})
	}
}

func setupTracer() *doubles.DummyTracer {
	return doubles.NewDummyTracer()
}

func setupMock(t *testing.T) (
	*mocks.MockValidator,
	*mocks.MockRepository,
	*mocks.MockAuthInteractor,
	*mocks.MockCache,
) {
	mockVld := mocks.NewMockValidator(t)
	mockRepo := mocks.NewMockRepository(t)
	mockAuth := mocks.NewMockAuthInteractor(t)
	mockCache := mocks.NewMockCache(t)

	return mockVld, mockRepo, mockAuth, mockCache
}
