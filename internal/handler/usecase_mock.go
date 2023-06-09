// Code generated by mockery v2.23.2. DO NOT EDIT.

package handler

import (
	context "context"

	domain "github.com/brisk84/gophkeeper/domain"
	mock "github.com/stretchr/testify/mock"
)

// useCaseMock is an autogenerated mock type for the useCase type
type useCaseMock struct {
	mock.Mock
}

// Auth provides a mock function with given fields: ctx, token
func (_m *useCaseMock) Auth(ctx context.Context, token string) (domain.User, error) {
	ret := _m.Called(ctx, token)

	var r0 domain.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (domain.User, error)); ok {
		return rf(ctx, token)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) domain.User); ok {
		r0 = rf(ctx, token)
	} else {
		r0 = ret.Get(0).(domain.User)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, token)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetData provides a mock function with given fields: ctx, userId, dataId
func (_m *useCaseMock) GetData(ctx context.Context, userId int, dataId int) ([]byte, error) {
	ret := _m.Called(ctx, userId, dataId)

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int, int) ([]byte, error)); ok {
		return rf(ctx, userId, dataId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int, int) []byte); ok {
		r0 = rf(ctx, userId, dataId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int, int) error); ok {
		r1 = rf(ctx, userId, dataId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListData provides a mock function with given fields: ctx, userId
func (_m *useCaseMock) ListData(ctx context.Context, userId int) ([]domain.Data, error) {
	ret := _m.Called(ctx, userId)

	var r0 []domain.Data
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int) ([]domain.Data, error)); ok {
		return rf(ctx, userId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int) []domain.Data); ok {
		r0 = rf(ctx, userId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Data)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, userId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Login provides a mock function with given fields: ctx, user
func (_m *useCaseMock) Login(ctx context.Context, user domain.User) (bool, string, error) {
	ret := _m.Called(ctx, user)

	var r0 bool
	var r1 string
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.User) (bool, string, error)); ok {
		return rf(ctx, user)
	}
	if rf, ok := ret.Get(0).(func(context.Context, domain.User) bool); ok {
		r0 = rf(ctx, user)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context, domain.User) string); ok {
		r1 = rf(ctx, user)
	} else {
		r1 = ret.Get(1).(string)
	}

	if rf, ok := ret.Get(2).(func(context.Context, domain.User) error); ok {
		r2 = rf(ctx, user)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Register provides a mock function with given fields: ctx, user
func (_m *useCaseMock) Register(ctx context.Context, user domain.User) (string, error) {
	ret := _m.Called(ctx, user)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.User) (string, error)); ok {
		return rf(ctx, user)
	}
	if rf, ok := ret.Get(0).(func(context.Context, domain.User) string); ok {
		r0 = rf(ctx, user)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, domain.User) error); ok {
		r1 = rf(ctx, user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SaveBankCard provides a mock function with given fields: ctx, userId, title, card
func (_m *useCaseMock) SaveBankCard(ctx context.Context, userId int, title string, card domain.CardInfo) (int, error) {
	ret := _m.Called(ctx, userId, title, card)

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int, string, domain.CardInfo) (int, error)); ok {
		return rf(ctx, userId, title, card)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int, string, domain.CardInfo) int); ok {
		r0 = rf(ctx, userId, title, card)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int, string, domain.CardInfo) error); ok {
		r1 = rf(ctx, userId, title, card)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SaveData provides a mock function with given fields: ctx, userId, title, data
func (_m *useCaseMock) SaveData(ctx context.Context, userId int, title string, data []byte) (int, error) {
	ret := _m.Called(ctx, userId, title, data)

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int, string, []byte) (int, error)); ok {
		return rf(ctx, userId, title, data)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int, string, []byte) int); ok {
		r0 = rf(ctx, userId, title, data)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int, string, []byte) error); ok {
		r1 = rf(ctx, userId, title, data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SaveLogin provides a mock function with given fields: ctx, userId, title, login, pass
func (_m *useCaseMock) SaveLogin(ctx context.Context, userId int, title string, login string, pass string) (int, error) {
	ret := _m.Called(ctx, userId, title, login, pass)

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int, string, string, string) (int, error)); ok {
		return rf(ctx, userId, title, login, pass)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int, string, string, string) int); ok {
		r0 = rf(ctx, userId, title, login, pass)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int, string, string, string) error); ok {
		r1 = rf(ctx, userId, title, login, pass)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SaveText provides a mock function with given fields: ctx, userId, title, text
func (_m *useCaseMock) SaveText(ctx context.Context, userId int, title string, text string) (int, error) {
	ret := _m.Called(ctx, userId, title, text)

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int, string, string) (int, error)); ok {
		return rf(ctx, userId, title, text)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int, string, string) int); ok {
		r0 = rf(ctx, userId, title, text)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int, string, string) error); ok {
		r1 = rf(ctx, userId, title, text)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTnewUseCaseMock interface {
	mock.TestingT
	Cleanup(func())
}

// newUseCaseMock creates a new instance of useCaseMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func newUseCaseMock(t mockConstructorTestingTnewUseCaseMock) *useCaseMock {
	mock := &useCaseMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
