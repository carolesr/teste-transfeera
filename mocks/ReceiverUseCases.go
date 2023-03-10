// Code generated by mockery v2.20.2. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	entity "github.com/teste-transfeera/internal/entity"

	usecase "github.com/teste-transfeera/internal/usecase"
)

// ReceiverUseCases is an autogenerated mock type for the ReceiverUseCases type
type ReceiverUseCases struct {
	mock.Mock
}

// Create provides a mock function with given fields: input
func (_m *ReceiverUseCases) Create(input *usecase.CreateReceiverInput) (*entity.Receiver, error) {
	ret := _m.Called(input)

	var r0 *entity.Receiver
	var r1 error
	if rf, ok := ret.Get(0).(func(*usecase.CreateReceiverInput) (*entity.Receiver, error)); ok {
		return rf(input)
	}
	if rf, ok := ret.Get(0).(func(*usecase.CreateReceiverInput) *entity.Receiver); ok {
		r0 = rf(input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Receiver)
		}
	}

	if rf, ok := ret.Get(1).(func(*usecase.CreateReceiverInput) error); ok {
		r1 = rf(input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: input
func (_m *ReceiverUseCases) Delete(input *usecase.DeleteReceiverInput) error {
	ret := _m.Called(input)

	var r0 error
	if rf, ok := ret.Get(0).(func(*usecase.DeleteReceiverInput) error); ok {
		r0 = rf(input)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// List provides a mock function with given fields: filter
func (_m *ReceiverUseCases) List(filter map[string]string) ([]entity.Receiver, error) {
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

// ListById provides a mock function with given fields: input
func (_m *ReceiverUseCases) ListById(input *usecase.ListReceiverByIdInput) (*entity.Receiver, error) {
	ret := _m.Called(input)

	var r0 *entity.Receiver
	var r1 error
	if rf, ok := ret.Get(0).(func(*usecase.ListReceiverByIdInput) (*entity.Receiver, error)); ok {
		return rf(input)
	}
	if rf, ok := ret.Get(0).(func(*usecase.ListReceiverByIdInput) *entity.Receiver); ok {
		r0 = rf(input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Receiver)
		}
	}

	if rf, ok := ret.Get(1).(func(*usecase.ListReceiverByIdInput) error); ok {
		r1 = rf(input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: input
func (_m *ReceiverUseCases) Update(input *usecase.UpdateReceiverInput) error {
	ret := _m.Called(input)

	var r0 error
	if rf, ok := ret.Get(0).(func(*usecase.UpdateReceiverInput) error); ok {
		r0 = rf(input)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewReceiverUseCases interface {
	mock.TestingT
	Cleanup(func())
}

// NewReceiverUseCases creates a new instance of ReceiverUseCases. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewReceiverUseCases(t mockConstructorTestingTNewReceiverUseCases) *ReceiverUseCases {
	mock := &ReceiverUseCases{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
