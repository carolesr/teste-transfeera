// Code generated by mockery v2.20.2. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	graph "github.com/teste-transfeera/internal/graph"
)

// ResolverRoot is an autogenerated mock type for the ResolverRoot type
type ResolverRoot struct {
	mock.Mock
}

// Mutation provides a mock function with given fields:
func (_m *ResolverRoot) Mutation() graph.MutationResolver {
	ret := _m.Called()

	var r0 graph.MutationResolver
	if rf, ok := ret.Get(0).(func() graph.MutationResolver); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(graph.MutationResolver)
		}
	}

	return r0
}

// Query provides a mock function with given fields:
func (_m *ResolverRoot) Query() graph.QueryResolver {
	ret := _m.Called()

	var r0 graph.QueryResolver
	if rf, ok := ret.Get(0).(func() graph.QueryResolver); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(graph.QueryResolver)
		}
	}

	return r0
}

type mockConstructorTestingTNewResolverRoot interface {
	mock.TestingT
	Cleanup(func())
}

// NewResolverRoot creates a new instance of ResolverRoot. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewResolverRoot(t mockConstructorTestingTNewResolverRoot) *ResolverRoot {
	mock := &ResolverRoot{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
