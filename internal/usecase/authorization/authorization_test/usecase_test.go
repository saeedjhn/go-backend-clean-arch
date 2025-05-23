package authorization_test

import (
	"context"
	"errors"
	"testing"

	"github.com/saeedjhn/go-backend-clean-arch/internal/sharedkernel/models"

	"github.com/saeedjhn/go-backend-clean-arch/internal/sharedkernel/types"

	"github.com/saeedjhn/go-backend-clean-arch/internal/usecase/authorization"
	"github.com/saeedjhn/go-backend-clean-arch/internal/usecase/authorization/authorization_test/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

//go:generate go test -v -race -count=1 ./...

func TestCheckAccess_NoActionsGiven_ReturnsFalse(t *testing.T) {
	i := authorization.New(authorization.Config{}, nil, nil)

	allowed, err := i.CheckAccess(context.Background(), []types.ID{1}, "resource", []models.Action{}...)
	require.NoError(t, err)
	assert.False(t, allowed)
}

func TestCheckAccess_ResourceNotFound_ReturnsError(t *testing.T) {
	mockResource := new(mocks.MockResourceInteractor)
	mockRoleResource := new(mocks.MockRoleResourcePermissionInteractor)

	mockResource.On("GetIDByName", mock.Anything, "resource").
		Return(uint64(0), errors.New("not found"))

	i := authorization.New(authorization.Config{}, mockResource, mockRoleResource)

	allowed, err := i.CheckAccess(context.Background(), []types.ID{1}, "resource", models.ReadAction)
	require.Error(t, err)
	assert.False(t, allowed)
}

// func TestCheckAccess_PermissionNotFound_ReturnsError(t *testing.T) {
// 	mockResource := new(mocks.MockResourceInteractor)
// 	mockRoleResource := new(mocks.MockRoleResourcePermissionInteractor)
//
// 	mockResource.On("GetIDByName", mock.Anything, "resource").
// 		Return(uint64(10), nil)
// 	mockRoleResource.On("GetByRoleIDAndResourceID", mock.Anything, uint64(1), uint64(10)).
// 		Return(models.RoleResourcePermission{}, errors.New("not found"))
//
// 	i := authorization.New(authorization.Config{}, mockResource, mockRoleResource)
//
// 	allowed, err := i.CheckAccess(context.Background(), []types.ID{1}, "resource", models.ReadAction)
// 	assert.Error(t, err)
// 	assert.False(t, allowed)
// }

func TestCheckAccess_AllowedRead_ReturnsTrue(t *testing.T) {
	mockResource := new(mocks.MockResourceInteractor)
	mockRoleResource := new(mocks.MockRoleResourcePermissionInteractor)

	mockResource.On("GetIDByName", mock.Anything, "resource").
		Return(uint64(10), nil)
	mockRoleResource.On("GetByRoleIDAndResourceID", mock.Anything, uint64(1), uint64(10)).
		Return(models.RoleResourcePermission{
			Permissions: models.Permission{
				Allow: models.RWXD{R: true},
				Deny:  models.RWXD{R: false},
			},
		}, nil)

	i := authorization.New(authorization.Config{}, mockResource, mockRoleResource)

	allowed, err := i.CheckAccess(context.Background(), []types.ID{1}, "resource", models.ReadAction)
	require.NoError(t, err)
	assert.True(t, allowed)
}

func TestCheckAccess_DeniedRead_ReturnsFalse(t *testing.T) {
	mockResource := new(mocks.MockResourceInteractor)
	mockRoleResource := new(mocks.MockRoleResourcePermissionInteractor)

	mockResource.On("GetIDByName", mock.Anything, "resource").Return(uint64(10), nil)
	mockRoleResource.On(
		"GetByRoleIDAndResourceID",
		mock.Anything,
		uint64(1),
		uint64(10)).Return(models.RoleResourcePermission{
		Permissions: models.Permission{
			Allow: models.RWXD{R: false},
			Deny:  models.RWXD{R: true},
		},
	}, nil)

	i := authorization.New(authorization.Config{}, mockResource, mockRoleResource)

	allowed, err := i.CheckAccess(context.Background(), []types.ID{1}, "resource", models.ReadAction)
	require.NoError(t, err)
	assert.False(t, allowed)
}

func TestCheckAccess_MultipleActions_AllAllowed_ReturnsTrue(t *testing.T) {
	mockResource := new(mocks.MockResourceInteractor)
	mockRoleResource := new(mocks.MockRoleResourcePermissionInteractor)

	mockResource.On("GetIDByName", mock.Anything, "resource").Return(uint64(10), nil)
	mockRoleResource.On(
		"GetByRoleIDAndResourceID",
		mock.Anything,
		uint64(1),
		uint64(10),
	).Return(models.RoleResourcePermission{
		Permissions: models.Permission{
			Allow: models.RWXD{R: true, W: true},
			Deny:  models.RWXD{R: false, W: false},
		},
	}, nil)

	i := authorization.New(authorization.Config{}, mockResource, mockRoleResource)

	allowed, err := i.CheckAccess(context.Background(), []types.ID{1}, "resource", models.ReadAction, models.WriteAction)
	require.NoError(t, err)
	assert.True(t, allowed)
}

func TestCheckAccess_MultipleActions_OneDenied_ReturnsFalse(t *testing.T) {
	mockResource := new(mocks.MockResourceInteractor)
	mockRoleResource := new(mocks.MockRoleResourcePermissionInteractor)

	mockResource.On("GetIDByName", mock.Anything, "resource").Return(uint64(10), nil)
	mockRoleResource.On(
		"GetByRoleIDAndResourceID",
		mock.Anything,
		uint64(1),
		uint64(10)).Return(models.RoleResourcePermission{
		Permissions: models.Permission{
			Allow: models.RWXD{R: true, W: true},
			Deny:  models.RWXD{R: false, W: true}, // WriteAction explicitly denied
		},
	}, nil)

	i := authorization.New(authorization.Config{}, mockResource, mockRoleResource)

	allowed, err := i.CheckAccess(context.Background(), []types.ID{1}, "resource", models.ReadAction, models.WriteAction)
	require.NoError(t, err)
	assert.False(t, allowed)
}

func TestCheckAccess_MultipleRolesAndActions_OneAllowed_ReturnsTrue(t *testing.T) {
	mockResource := new(mocks.MockResourceInteractor)
	mockRoleResource := new(mocks.MockRoleResourcePermissionInteractor)

	mockResource.On("GetIDByName", mock.Anything, "resource").Return(uint64(10), nil)
	mockRoleResource.On("GetByRoleIDAndResourceID", mock.Anything, uint64(1), uint64(10)).
		Return(models.RoleResourcePermission{}, errors.New("not found"))
	mockRoleResource.On(
		"GetByRoleIDAndResourceID",
		mock.Anything,
		uint64(2),
		uint64(10),
	).Return(models.RoleResourcePermission{
		Permissions: models.Permission{
			Allow: models.RWXD{R: true, W: true},
			Deny:  models.RWXD{R: false, W: false}, // WriteAction explicitly denied
		},
	}, nil)

	i := authorization.New(authorization.Config{}, mockResource, mockRoleResource)

	allowed, err := i.CheckAccess(
		context.Background(),
		[]types.ID{1, 2},
		"resource",
		models.ReadAction,
		models.WriteAction,
	)
	require.NoError(t, err)
	assert.True(t, allowed)
}

func TestCheckAccess_MultipleRoles_OneAllowed_ReturnsTrue(t *testing.T) {
	mockResource := new(mocks.MockResourceInteractor)
	mockRoleResource := new(mocks.MockRoleResourcePermissionInteractor)

	mockResource.On("GetIDByName", mock.Anything, "resource").
		Return(uint64(10), nil)
	mockRoleResource.On("GetByRoleIDAndResourceID", mock.Anything, uint64(1), uint64(10)).
		Return(models.RoleResourcePermission{}, errors.New("not found"))
	mockRoleResource.On("GetByRoleIDAndResourceID", mock.Anything, uint64(2), uint64(10)).
		Return(models.RoleResourcePermission{
			Permissions: models.Permission{
				Allow: models.RWXD{R: true},
				Deny:  models.RWXD{R: false},
			},
		}, nil)

	i := authorization.New(authorization.Config{}, mockResource, mockRoleResource)

	allowed, err := i.CheckAccess(context.Background(), []types.ID{1, 2}, "resource", models.ReadAction)
	require.NoError(t, err)
	assert.True(t, allowed)
}

func TestCheckAccess_MultipleRoles_AllDenied_ReturnsFalse(t *testing.T) {
	mockResource := new(mocks.MockResourceInteractor)
	mockRoleResource := new(mocks.MockRoleResourcePermissionInteractor)

	mockResource.On("GetIDByName", mock.Anything, "resource").
		Return(uint64(10), nil)
	mockRoleResource.On("GetByRoleIDAndResourceID", mock.Anything, uint64(1), uint64(10)).
		Return(models.RoleResourcePermission{
			Permissions: models.Permission{
				Allow: models.RWXD{R: false},
				Deny:  models.RWXD{R: true},
			},
		}, nil)
	mockRoleResource.On("GetByRoleIDAndResourceID", mock.Anything, uint64(2), uint64(10)).
		Return(models.RoleResourcePermission{
			Permissions: models.Permission{
				Allow: models.RWXD{R: false},
				Deny:  models.RWXD{R: true},
			},
		}, nil)

	i := authorization.New(authorization.Config{}, mockResource, mockRoleResource)

	allowed, err := i.CheckAccess(context.Background(), []types.ID{1, 2}, "resource", models.ReadAction)
	require.NoError(t, err)
	assert.False(t, allowed)
}
