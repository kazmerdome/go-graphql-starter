// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	licence "github.com/kazmerdome/go-graphql-starter/pkg/domain/licence"
	mock "github.com/stretchr/testify/mock"

	primitive "go.mongodb.org/mongo-driver/bson/primitive"
)

// LicenceRepository is an autogenerated mock type for the LicenceRepository type
type LicenceRepository struct {
	mock.Mock
}

// Count provides a mock function with given fields: filter
func (_m *LicenceRepository) Count(filter *licence.LicenceWhereDTO) (*int, error) {
	ret := _m.Called(filter)

	var r0 *int
	if rf, ok := ret.Get(0).(func(*licence.LicenceWhereDTO) *int); ok {
		r0 = rf(filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*licence.LicenceWhereDTO) error); ok {
		r1 = rf(filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Create provides a mock function with given fields: data
func (_m *LicenceRepository) Create(data *licence.LicenceCreateDTO) (*licence.Licence, error) {
	ret := _m.Called(data)

	var r0 *licence.Licence
	if rf, ok := ret.Get(0).(func(*licence.LicenceCreateDTO) *licence.Licence); ok {
		r0 = rf(data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*licence.Licence)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*licence.LicenceCreateDTO) error); ok {
		r1 = rf(data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: where
func (_m *LicenceRepository) Delete(where primitive.ObjectID) (*licence.Licence, error) {
	ret := _m.Called(where)

	var r0 *licence.Licence
	if rf, ok := ret.Get(0).(func(primitive.ObjectID) *licence.Licence); ok {
		r0 = rf(where)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*licence.Licence)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(primitive.ObjectID) error); ok {
		r1 = rf(where)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: filter, orderBy, skip, limit, customQuery
func (_m *LicenceRepository) List(filter *licence.LicenceWhereDTO, orderBy *licence.LicenceOrderByENUM, skip *int, limit *int, customQuery *primitive.M) ([]*licence.Licence, error) {
	ret := _m.Called(filter, orderBy, skip, limit, customQuery)

	var r0 []*licence.Licence
	if rf, ok := ret.Get(0).(func(*licence.LicenceWhereDTO, *licence.LicenceOrderByENUM, *int, *int, *primitive.M) []*licence.Licence); ok {
		r0 = rf(filter, orderBy, skip, limit, customQuery)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*licence.Licence)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*licence.LicenceWhereDTO, *licence.LicenceOrderByENUM, *int, *int, *primitive.M) error); ok {
		r1 = rf(filter, orderBy, skip, limit, customQuery)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// One provides a mock function with given fields: filter
func (_m *LicenceRepository) One(filter *licence.LicenceWhereDTO) (*licence.Licence, error) {
	ret := _m.Called(filter)

	var r0 *licence.Licence
	if rf, ok := ret.Get(0).(func(*licence.LicenceWhereDTO) *licence.Licence); ok {
		r0 = rf(filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*licence.Licence)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*licence.LicenceWhereDTO) error); ok {
		r1 = rf(filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: where, data
func (_m *LicenceRepository) Update(where primitive.ObjectID, data *licence.LicenceUpdateDTO) (*licence.Licence, error) {
	ret := _m.Called(where, data)

	var r0 *licence.Licence
	if rf, ok := ret.Get(0).(func(primitive.ObjectID, *licence.LicenceUpdateDTO) *licence.Licence); ok {
		r0 = rf(where, data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*licence.Licence)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(primitive.ObjectID, *licence.LicenceUpdateDTO) error); ok {
		r1 = rf(where, data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
