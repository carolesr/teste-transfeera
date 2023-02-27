// Code generated by mockery v2.20.2. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	entity "github.com/teste-transfeera/internal/entity"
)

// ReceiverRepository is an autogenerated mock type for the ReceiverRepository type
type ReceiverRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: receiver
func (_m *ReceiverRepository) Create(receiver entity.Receiver) (*entity.Receiver, error) {
	ret := _m.Called(receiver)

	var r0 *entity.Receiver
	var r1 error
	if rf, ok := ret.Get(0).(func(entity.Receiver) (*entity.Receiver, error)); ok {
		return rf(receiver)
	}
	if rf, ok := ret.Get(0).(func(entity.Receiver) *entity.Receiver); ok {
		r0 = rf(receiver)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Receiver)
		}
	}

	if rf, ok := ret.Get(1).(func(entity.Receiver) error); ok {
		r1 = rf(receiver)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ids
func (_m *ReceiverRepository) Delete(ids []string) error {
	ret := _m.Called(ids)

	var r0 error
	if rf, ok := ret.Get(0).(func([]string) error); ok {
		r0 = rf(ids)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindById provides a mock function with given fields: id
func (_m *ReceiverRepository) FindById(id string) (*entity.Receiver, error) {
	ret := _m.Called(id)

	var r0 *entity.Receiver
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*entity.Receiver, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(string) *entity.Receiver); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Receiver)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: filter
func (_m *ReceiverRepository) List(filter map[string]string) ([]entity.Receiver, error) {
	ret := _m.Called(filter)

	var r0 []entity.Receiver
	var r1 error
	if rf, ok := ret.Get(0).(func(map[string]string) ([]entity.Receiver, error)); ok {
		return rf(filter)
	}
	if rf, ok := ret.Get(0).(func(map[string]string) []entity.Receiver); ok {
		r0 = rf(filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.Receiver)
		}
	}

	if rf, ok := ret.Get(1).(func(map[string]string) error); ok {
		r1 = rf(filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: id, fields
func (_m *ReceiverRepository) Update(id string, fields map[string]string) error {
	ret := _m.Called(id, fields)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, map[string]string) error); ok {
		r0 = rf(id, fields)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewReceiverRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewReceiverRepository creates a new instance of ReceiverRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewReceiverRepository(t mockConstructorTestingTNewReceiverRepository) *ReceiverRepository {
	mock := &ReceiverRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
