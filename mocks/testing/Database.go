// Code generated by mockery v2.34.2. DO NOT EDIT.

package mocks

import (
	seeder "github.com/goravel/framework/contracts/database/seeder"
	testing "github.com/goravel/framework/contracts/testing"
	mock "github.com/stretchr/testify/mock"
)

// Database is an autogenerated mock type for the Database type
type Database struct {
	mock.Mock
}

// Build provides a mock function with given fields:
func (_m *Database) Build() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Clear provides a mock function with given fields:
func (_m *Database) Clear() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Config provides a mock function with given fields:
func (_m *Database) Config() testing.DatabaseConfig {
	ret := _m.Called()

	var r0 testing.DatabaseConfig
	if rf, ok := ret.Get(0).(func() testing.DatabaseConfig); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(testing.DatabaseConfig)
	}

	return r0
}

// Image provides a mock function with given fields: _a0
func (_m *Database) Image(_a0 testing.Image) {
	_m.Called(_a0)
}

// Seed provides a mock function with given fields: seeds
func (_m *Database) Seed(seeds ...seeder.Seeder) {
	_va := make([]interface{}, len(seeds))
	for _i := range seeds {
		_va[_i] = seeds[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	_m.Called(_ca...)
}

// Stop provides a mock function with given fields:
func (_m *Database) Stop() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewDatabase creates a new instance of Database. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewDatabase(t interface {
	mock.TestingT
	Cleanup(func())
}) *Database {
	mock := &Database{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
