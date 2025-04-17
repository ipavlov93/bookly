// Code generated by mockery v2.53.3. DO NOT EDIT.

package mocks

import (
	claims "event-calendar/internal/domain/claims"

	auth "firebase.google.com/go/v4/auth"

	context "context"

	mock "github.com/stretchr/testify/mock"
)

// FirebaseAuthService is an autogenerated mock type for the FirebaseAuthService type
type FirebaseAuthService struct {
	mock.Mock
}

// RevokeRefreshTokens provides a mock function with given fields: ctx, idToken
func (_m *FirebaseAuthService) RevokeRefreshTokens(ctx context.Context, idToken string) error {
	ret := _m.Called(ctx, idToken)

	if len(ret) == 0 {
		panic("no return value specified for RevokeRefreshTokens")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, idToken)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetRolePrivilegesToClaims provides a mock function with given fields: firebaseUID, roles
func (_m *FirebaseAuthService) SetRolePrivilegesToClaims(firebaseUID string, roles []claims.Role) error {
	ret := _m.Called(firebaseUID, roles)

	if len(ret) == 0 {
		panic("no return value specified for SetRolePrivilegesToClaims")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, []claims.Role) error); ok {
		r0 = rf(firebaseUID, roles)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// VerifyIDToken provides a mock function with given fields: idToken
func (_m *FirebaseAuthService) VerifyIDToken(idToken string) (*auth.Token, error) {
	ret := _m.Called(idToken)

	if len(ret) == 0 {
		panic("no return value specified for VerifyIDToken")
	}

	var r0 *auth.Token
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*auth.Token, error)); ok {
		return rf(idToken)
	}
	if rf, ok := ret.Get(0).(func(string) *auth.Token); ok {
		r0 = rf(idToken)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*auth.Token)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(idToken)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewFirebaseAuthService creates a new instance of FirebaseAuthService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewFirebaseAuthService(t interface {
	mock.TestingT
	Cleanup(func())
}) *FirebaseAuthService {
	mock := &FirebaseAuthService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
