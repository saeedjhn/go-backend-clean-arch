// Code generated by mockery v2.52.4. DO NOT EDIT.

package mocks

import (
	context "context"
	time "time"

	models "github.com/saeedjhn/go-backend-clean-arch/internal/sharedkernel/models"
	mock "github.com/stretchr/testify/mock"

	types "github.com/saeedjhn/go-backend-clean-arch/internal/sharedkernel/types"
)

// MockRepository is an autogenerated mock type for the Repository type
type MockRepository struct {
	mock.Mock
}

type MockRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *MockRepository) EXPECT() *MockRepository_Expecter {
	return &MockRepository_Expecter{mock: &_m.Mock}
}

// GetUnPublished provides a mock function with given fields: ctx, offset, limit, retryThreshold
func (_m *MockRepository) GetUnPublished(ctx context.Context, offset int, limit int, retryThreshold int) ([]models.OutboxEvent, error) {
	ret := _m.Called(ctx, offset, limit, retryThreshold)

	if len(ret) == 0 {
		panic("no return value specified for GetUnPublished")
	}

	var r0 []models.OutboxEvent
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int, int, int) ([]models.OutboxEvent, error)); ok {
		return rf(ctx, offset, limit, retryThreshold)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int, int, int) []models.OutboxEvent); ok {
		r0 = rf(ctx, offset, limit, retryThreshold)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.OutboxEvent)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int, int, int) error); ok {
		r1 = rf(ctx, offset, limit, retryThreshold)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRepository_GetUnPublished_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetUnPublished'
type MockRepository_GetUnPublished_Call struct {
	*mock.Call
}

// GetUnPublished is a helper method to define mock.On call
//   - ctx context.Context
//   - offset int
//   - limit int
//   - retryThreshold int
func (_e *MockRepository_Expecter) GetUnPublished(ctx interface{}, offset interface{}, limit interface{}, retryThreshold interface{}) *MockRepository_GetUnPublished_Call {
	return &MockRepository_GetUnPublished_Call{Call: _e.mock.On("GetUnPublished", ctx, offset, limit, retryThreshold)}
}

func (_c *MockRepository_GetUnPublished_Call) Run(run func(ctx context.Context, offset int, limit int, retryThreshold int)) *MockRepository_GetUnPublished_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int), args[2].(int), args[3].(int))
	})
	return _c
}

func (_c *MockRepository_GetUnPublished_Call) Return(_a0 []models.OutboxEvent, _a1 error) *MockRepository_GetUnPublished_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRepository_GetUnPublished_Call) RunAndReturn(run func(context.Context, int, int, int) ([]models.OutboxEvent, error)) *MockRepository_GetUnPublished_Call {
	_c.Call.Return(run)
	return _c
}

// UnpublishedCount provides a mock function with given fields: ctx, retryThreshold
func (_m *MockRepository) UnpublishedCount(ctx context.Context, retryThreshold int64) (int64, error) {
	ret := _m.Called(ctx, retryThreshold)

	if len(ret) == 0 {
		panic("no return value specified for UnpublishedCount")
	}

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) (int64, error)); ok {
		return rf(ctx, retryThreshold)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) int64); ok {
		r0 = rf(ctx, retryThreshold)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, retryThreshold)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRepository_UnpublishedCount_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UnpublishedCount'
type MockRepository_UnpublishedCount_Call struct {
	*mock.Call
}

// UnpublishedCount is a helper method to define mock.On call
//   - ctx context.Context
//   - retryThreshold int64
func (_e *MockRepository_Expecter) UnpublishedCount(ctx interface{}, retryThreshold interface{}) *MockRepository_UnpublishedCount_Call {
	return &MockRepository_UnpublishedCount_Call{Call: _e.mock.On("UnpublishedCount", ctx, retryThreshold)}
}

func (_c *MockRepository_UnpublishedCount_Call) Run(run func(ctx context.Context, retryThreshold int64)) *MockRepository_UnpublishedCount_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64))
	})
	return _c
}

func (_c *MockRepository_UnpublishedCount_Call) Return(_a0 int64, _a1 error) *MockRepository_UnpublishedCount_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRepository_UnpublishedCount_Call) RunAndReturn(run func(context.Context, int64) (int64, error)) *MockRepository_UnpublishedCount_Call {
	_c.Call.Return(run)
	return _c
}

// UpdatePublished provides a mock function with given fields: ctx, eventIDs, publishedAt
func (_m *MockRepository) UpdatePublished(ctx context.Context, eventIDs []types.ID, publishedAt time.Time) error {
	ret := _m.Called(ctx, eventIDs, publishedAt)

	if len(ret) == 0 {
		panic("no return value specified for UpdatePublished")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []types.ID, time.Time) error); ok {
		r0 = rf(ctx, eventIDs, publishedAt)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockRepository_UpdatePublished_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdatePublished'
type MockRepository_UpdatePublished_Call struct {
	*mock.Call
}

// UpdatePublished is a helper method to define mock.On call
//   - ctx context.Context
//   - eventIDs []types.ID
//   - publishedAt time.Time
func (_e *MockRepository_Expecter) UpdatePublished(ctx interface{}, eventIDs interface{}, publishedAt interface{}) *MockRepository_UpdatePublished_Call {
	return &MockRepository_UpdatePublished_Call{Call: _e.mock.On("UpdatePublished", ctx, eventIDs, publishedAt)}
}

func (_c *MockRepository_UpdatePublished_Call) Run(run func(ctx context.Context, eventIDs []types.ID, publishedAt time.Time)) *MockRepository_UpdatePublished_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([]types.ID), args[2].(time.Time))
	})
	return _c
}

func (_c *MockRepository_UpdatePublished_Call) Return(_a0 error) *MockRepository_UpdatePublished_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockRepository_UpdatePublished_Call) RunAndReturn(run func(context.Context, []types.ID, time.Time) error) *MockRepository_UpdatePublished_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateUnpublished provides a mock function with given fields: ctx, eventIDs, lastRetriedAt
func (_m *MockRepository) UpdateUnpublished(ctx context.Context, eventIDs []types.ID, lastRetriedAt time.Time) error {
	ret := _m.Called(ctx, eventIDs, lastRetriedAt)

	if len(ret) == 0 {
		panic("no return value specified for UpdateUnpublished")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []types.ID, time.Time) error); ok {
		r0 = rf(ctx, eventIDs, lastRetriedAt)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockRepository_UpdateUnpublished_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateUnpublished'
type MockRepository_UpdateUnpublished_Call struct {
	*mock.Call
}

// UpdateUnpublished is a helper method to define mock.On call
//   - ctx context.Context
//   - eventIDs []types.ID
//   - lastRetriedAt time.Time
func (_e *MockRepository_Expecter) UpdateUnpublished(ctx interface{}, eventIDs interface{}, lastRetriedAt interface{}) *MockRepository_UpdateUnpublished_Call {
	return &MockRepository_UpdateUnpublished_Call{Call: _e.mock.On("UpdateUnpublished", ctx, eventIDs, lastRetriedAt)}
}

func (_c *MockRepository_UpdateUnpublished_Call) Run(run func(ctx context.Context, eventIDs []types.ID, lastRetriedAt time.Time)) *MockRepository_UpdateUnpublished_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([]types.ID), args[2].(time.Time))
	})
	return _c
}

func (_c *MockRepository_UpdateUnpublished_Call) Return(_a0 error) *MockRepository_UpdateUnpublished_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockRepository_UpdateUnpublished_Call) RunAndReturn(run func(context.Context, []types.ID, time.Time) error) *MockRepository_UpdateUnpublished_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockRepository creates a new instance of MockRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockRepository {
	mock := &MockRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
