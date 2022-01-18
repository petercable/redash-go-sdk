// Code generated by mockery v2.9.4. DO NOT EDIT.

package usersmock

import (
	runtime "github.com/go-openapi/runtime"
	mock "github.com/stretchr/testify/mock"

	users "github.com/recolabs/reco/redash-client/gen/client/users"
)

// ClientService is an autogenerated mock type for the ClientService type
type ClientService struct {
	mock.Mock
}

// GetUsersID provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) GetUsersID(params *users.GetUsersIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...users.ClientOption) (*users.GetUsersIDOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *users.GetUsersIDOK
	if rf, ok := ret.Get(0).(func(*users.GetUsersIDParams, runtime.ClientAuthInfoWriter, ...users.ClientOption) *users.GetUsersIDOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*users.GetUsersIDOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*users.GetUsersIDParams, runtime.ClientAuthInfoWriter, ...users.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetTransport provides a mock function with given fields: transport
func (_m *ClientService) SetTransport(transport runtime.ClientTransport) {
	_m.Called(transport)
}
