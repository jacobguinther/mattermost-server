// Code generated by mockery v1.0.0. DO NOT EDIT.

// Regenerate this file using `make store-mocks`.

package mocks

import (
	model "github.com/mattermost/mattermost-server/v5/model"
	mock "github.com/stretchr/testify/mock"
)

// ThreadStore is an autogenerated mock type for the ThreadStore type
type ThreadStore struct {
	mock.Mock
}

// Delete provides a mock function with given fields: postId
func (_m *ThreadStore) Delete(postId string) error {
	ret := _m.Called(postId)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(postId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteMembershipForUser provides a mock function with given fields: userId, postId
func (_m *ThreadStore) DeleteMembershipForUser(userId string, postId string) error {
	ret := _m.Called(userId, postId)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(userId, postId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: id
func (_m *ThreadStore) Get(id string) (*model.Thread, error) {
	ret := _m.Called(id)

	var r0 *model.Thread
	if rf, ok := ret.Get(0).(func(string) *model.Thread); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Thread)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMembershipForUser provides a mock function with given fields: userId, postId
func (_m *ThreadStore) GetMembershipForUser(userId string, postId string) (*model.ThreadMembership, error) {
	ret := _m.Called(userId, postId)

	var r0 *model.ThreadMembership
	if rf, ok := ret.Get(0).(func(string, string) *model.ThreadMembership); ok {
		r0 = rf(userId, postId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.ThreadMembership)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(userId, postId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMembershipsForUser provides a mock function with given fields: userId
func (_m *ThreadStore) GetMembershipsForUser(userId string) ([]*model.ThreadMembership, error) {
	ret := _m.Called(userId)

	var r0 []*model.ThreadMembership
	if rf, ok := ret.Get(0).(func(string) []*model.ThreadMembership); ok {
		r0 = rf(userId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.ThreadMembership)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(userId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Save provides a mock function with given fields: thread
func (_m *ThreadStore) Save(thread *model.Thread) (*model.Thread, error) {
	ret := _m.Called(thread)

	var r0 *model.Thread
	if rf, ok := ret.Get(0).(func(*model.Thread) *model.Thread); ok {
		r0 = rf(thread)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Thread)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*model.Thread) error); ok {
		r1 = rf(thread)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SaveMembership provides a mock function with given fields: membership
func (_m *ThreadStore) SaveMembership(membership *model.ThreadMembership) (*model.ThreadMembership, error) {
	ret := _m.Called(membership)

	var r0 *model.ThreadMembership
	if rf, ok := ret.Get(0).(func(*model.ThreadMembership) *model.ThreadMembership); ok {
		r0 = rf(membership)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.ThreadMembership)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*model.ThreadMembership) error); ok {
		r1 = rf(membership)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SaveMultiple provides a mock function with given fields: thread
func (_m *ThreadStore) SaveMultiple(thread []*model.Thread) ([]*model.Thread, int, error) {
	ret := _m.Called(thread)

	var r0 []*model.Thread
	if rf, ok := ret.Get(0).(func([]*model.Thread) []*model.Thread); ok {
		r0 = rf(thread)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Thread)
		}
	}

	var r1 int
	if rf, ok := ret.Get(1).(func([]*model.Thread) int); ok {
		r1 = rf(thread)
	} else {
		r1 = ret.Get(1).(int)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func([]*model.Thread) error); ok {
		r2 = rf(thread)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Update provides a mock function with given fields: thread
func (_m *ThreadStore) Update(thread *model.Thread) (*model.Thread, error) {
	ret := _m.Called(thread)

	var r0 *model.Thread
	if rf, ok := ret.Get(0).(func(*model.Thread) *model.Thread); ok {
		r0 = rf(thread)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Thread)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*model.Thread) error); ok {
		r1 = rf(thread)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateMembership provides a mock function with given fields: membership
func (_m *ThreadStore) UpdateMembership(membership *model.ThreadMembership) (*model.ThreadMembership, error) {
	ret := _m.Called(membership)

	var r0 *model.ThreadMembership
	if rf, ok := ret.Get(0).(func(*model.ThreadMembership) *model.ThreadMembership); ok {
		r0 = rf(membership)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.ThreadMembership)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*model.ThreadMembership) error); ok {
		r1 = rf(membership)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateMembershipFromMention provides a mock function with given fields: userId, postId
func (_m *ThreadStore) UpdateMembershipFromMention(userId string, postId string) error {
	ret := _m.Called(userId, postId)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(userId, postId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
