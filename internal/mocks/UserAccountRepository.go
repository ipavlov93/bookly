// Code generated by mockery v2.53.3. DO NOT EDIT.

package mocks

import (
	context "context"
	domain "event-calendar/internal/domain"

	mock "github.com/stretchr/testify/mock"
)

// UserAccountRepository is an autogenerated mock type for the UserAccountRepository type
type UserAccountRepository struct {
	mock.Mock
}

// CreateUserAccount provides a mock function with given fields: ctx, user, ignoreDuplicate
func (_m *UserAccountRepository) CreateUserAccount(ctx context.Context, user domain.UserAccount, ignoreDuplicate bool) (int64, error) {
	ret := _m.Called(ctx, user, ignoreDuplicate)

	if len(ret) == 0 {
		panic("no return value specified for CreateUserAccount")
	}

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.UserAccount, bool) (int64, error)); ok {
		return rf(ctx, user, ignoreDuplicate)
	}
	if rf, ok := ret.Get(0).(func(context.Context, domain.UserAccount, bool) int64); ok {
		r0 = rf(ctx, user, ignoreDuplicate)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, domain.UserAccount, bool) error); ok {
		r1 = rf(ctx, user, ignoreDuplicate)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListUserAccountsByUserID provides a mock function with given fields: ctx, userID
func (_m *UserAccountRepository) ListUserAccountsByUserID(ctx context.Context, userID int64) ([]domain.UserAccount, error) {
	ret := _m.Called(ctx, userID)

	if len(ret) == 0 {
		panic("no return value specified for ListUserAccountsByUserID")
	}

	var r0 []domain.UserAccount
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) ([]domain.UserAccount, error)); ok {
		return rf(ctx, userID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) []domain.UserAccount); ok {
		r0 = rf(ctx, userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.UserAccount)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewUserAccountRepository creates a new instance of UserAccountRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUserAccountRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *UserAccountRepository {
	mock := &UserAccountRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
