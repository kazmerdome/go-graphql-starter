// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	options "go.mongodb.org/mongo-driver/mongo/options"

	repository "github.com/kazmerdome/go-graphql-starter/pkg/repository"
)

// MongoDatabase is an autogenerated mock type for the MongoDatabase type
type MongoDatabase struct {
	mock.Mock
}

// Collection provides a mock function with given fields: name, opts
func (_m *MongoDatabase) Collection(name string, opts ...*options.CollectionOptions) repository.MongoCollection {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, name)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 repository.MongoCollection
	if rf, ok := ret.Get(0).(func(string, ...*options.CollectionOptions) repository.MongoCollection); ok {
		r0 = rf(name, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(repository.MongoCollection)
		}
	}

	return r0
}

// Disconnect provides a mock function with given fields:
func (_m *MongoDatabase) Disconnect() {
	_m.Called()
}