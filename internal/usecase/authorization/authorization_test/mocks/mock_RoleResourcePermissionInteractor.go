// Code generated by mockery v2.52.4. DO NOT EDIT.

package mocks

import (
	context "context"

	models "github.com/saeedjhn/go-backend-clean-arch/internal/sharedkernel/models"
	mock "github.com/stretchr/testify/mock"
)

// MockRoleResourcePermissionInteractor is an autogenerated mock type for the RoleResourcePermissionInteractor type
type MockRoleResourcePermissionInteractor struct {
	mock.Mock
}

type MockRoleResourcePermissionInteractor_Expecter struct {
	mock *mock.Mock
}

func (_m *MockRoleResourcePermissionInteractor) EXPECT() *MockRoleResourcePermissionInteractor_Expecter {
	return &MockRoleResourcePermissionInteractor_Expecter{mock: &_m.Mock}
}

// GetByRoleIDAndResourceID provides a mock function with given fields: ctx, roleID, resourceID
func (_m *MockRoleResourcePermissionInteractor) GetByRoleIDAndResourceID(ctx context.Context, roleID uint64, resourceID uint64) (models.RoleResourcePermission, error) {
	ret := _m.Called(ctx, roleID, resourceID)

	if len(ret) == 0 {
		panic("no return value specified for GetByRoleIDAndResourceID")
	}

	var r0 models.RoleResourcePermission
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64, uint64) (models.RoleResourcePermission, error)); ok {
		return rf(ctx, roleID, resourceID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64, uint64) models.RoleResourcePermission); ok {
		r0 = rf(ctx, roleID, resourceID)
	} else {
		r0 = ret.Get(0).(models.RoleResourcePermission)
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64, uint64) error); ok {
		r1 = rf(ctx, roleID, resourceID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRoleResourcePermissionInteractor_GetByRoleIDAndResourceID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByRoleIDAndResourceID'
type MockRoleResourcePermissionInteractor_GetByRoleIDAndResourceID_Call struct {
	*mock.Call
}

// GetByRoleIDAndResourceID is a helper method to define mock.On call
//   - ctx context.Context
//   - roleID uint64
//   - resourceID uint64
func (_e *MockRoleResourcePermissionInteractor_Expecter) GetByRoleIDAndResourceID(ctx interface{}, roleID interface{}, resourceID interface{}) *MockRoleResourcePermissionInteractor_GetByRoleIDAndResourceID_Call {
	return &MockRoleResourcePermissionInteractor_GetByRoleIDAndResourceID_Call{Call: _e.mock.On("GetByRoleIDAndResourceID", ctx, roleID, resourceID)}
}

func (_c *MockRoleResourcePermissionInteractor_GetByRoleIDAndResourceID_Call) Run(run func(ctx context.Context, roleID uint64, resourceID uint64)) *MockRoleResourcePermissionInteractor_GetByRoleIDAndResourceID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint64), args[2].(uint64))
	})
	return _c
}

func (_c *MockRoleResourcePermissionInteractor_GetByRoleIDAndResourceID_Call) Return(_a0 models.RoleResourcePermission, _a1 error) *MockRoleResourcePermissionInteractor_GetByRoleIDAndResourceID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRoleResourcePermissionInteractor_GetByRoleIDAndResourceID_Call) RunAndReturn(run func(context.Context, uint64, uint64) (models.RoleResourcePermission, error)) *MockRoleResourcePermissionInteractor_GetByRoleIDAndResourceID_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockRoleResourcePermissionInteractor creates a new instance of MockRoleResourcePermissionInteractor. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockRoleResourcePermissionInteractor(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockRoleResourcePermissionInteractor {
	mock := &MockRoleResourcePermissionInteractor{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
