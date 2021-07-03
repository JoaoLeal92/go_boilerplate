// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// HashProvider is an autogenerated mock type for the HashProvider type
type HashProvider struct {
	mock.Mock
}

// CompareHashAndPassword provides a mock function with given fields: userPassword, inputPassword
func (_m *HashProvider) CompareHashAndPassword(userPassword string, inputPassword string) error {
	ret := _m.Called(userPassword, inputPassword)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(userPassword, inputPassword)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GenerateHash provides a mock function with given fields: password
func (_m *HashProvider) GenerateHash(password string) (string, error) {
	ret := _m.Called(password)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(password)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}