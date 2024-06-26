// Code generated by mockery v2.23.1. DO NOT EDIT.

package mocks

import (
	context "context"
	entity "recommendation-service/internal/module/recommendation/models/entity"

	mock "github.com/stretchr/testify/mock"

	response "recommendation-service/internal/module/recommendation/models/response"
)

// Repositories is an autogenerated mock type for the Repositories type
type Repositories struct {
	mock.Mock
}

// FindTicketByRegionName provides a mock function with given fields: ctx, regionName
func (_m *Repositories) FindTicketByRegionName(ctx context.Context, regionName string) ([]response.Ticket, error) {
	ret := _m.Called(ctx, regionName)

	var r0 []response.Ticket
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]response.Ticket, error)); ok {
		return rf(ctx, regionName)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []response.Ticket); ok {
		r0 = rf(ctx, regionName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]response.Ticket)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, regionName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindUserProfile provides a mock function with given fields: ctx, userID
func (_m *Repositories) FindUserProfile(ctx context.Context, userID int64) (response.UserProfile, error) {
	ret := _m.Called(ctx, userID)

	var r0 response.UserProfile
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) (response.UserProfile, error)); ok {
		return rf(ctx, userID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) response.UserProfile); ok {
		r0 = rf(ctx, userID)
	} else {
		r0 = ret.Get(0).(response.UserProfile)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindVenueByName provides a mock function with given fields: ctx, name
func (_m *Repositories) FindVenueByName(ctx context.Context, name string) (entity.Venues, error) {
	ret := _m.Called(ctx, name)

	var r0 entity.Venues
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (entity.Venues, error)); ok {
		return rf(ctx, name)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) entity.Venues); ok {
		r0 = rf(ctx, name)
	} else {
		r0 = ret.Get(0).(entity.Venues)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindVenues provides a mock function with given fields: ctx
func (_m *Repositories) FindVenues(ctx context.Context) ([]entity.Venues, error) {
	ret := _m.Called(ctx)

	var r0 []entity.Venues
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]entity.Venues, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []entity.Venues); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.Venues)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpsertVenue provides a mock function with given fields: ctx, payload
func (_m *Repositories) UpsertVenue(ctx context.Context, payload entity.Venues) error {
	ret := _m.Called(ctx, payload)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, entity.Venues) error); ok {
		r0 = rf(ctx, payload)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ValidateToken provides a mock function with given fields: ctx, token
func (_m *Repositories) ValidateToken(ctx context.Context, token string) (response.UserServiceValidate, error) {
	ret := _m.Called(ctx, token)

	var r0 response.UserServiceValidate
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (response.UserServiceValidate, error)); ok {
		return rf(ctx, token)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) response.UserServiceValidate); ok {
		r0 = rf(ctx, token)
	} else {
		r0 = ret.Get(0).(response.UserServiceValidate)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, token)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewRepositories interface {
	mock.TestingT
	Cleanup(func())
}

// NewRepositories creates a new instance of Repositories. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRepositories(t mockConstructorTestingTNewRepositories) *Repositories {
	mock := &Repositories{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
