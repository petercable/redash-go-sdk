// Code generated by mockery v2.9.4. DO NOT EDIT.

package visualizationsmock

import (
	runtime "github.com/go-openapi/runtime"
	mock "github.com/stretchr/testify/mock"

	visualizations "github.com/recolabs/reco/redash-client/gen/client/visualizations"
)

// ClientService is an autogenerated mock type for the ClientService type
type ClientService struct {
	mock.Mock
}

// DeleteVisualizationsID provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) DeleteVisualizationsID(params *visualizations.DeleteVisualizationsIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...visualizations.ClientOption) (*visualizations.DeleteVisualizationsIDOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *visualizations.DeleteVisualizationsIDOK
	if rf, ok := ret.Get(0).(func(*visualizations.DeleteVisualizationsIDParams, runtime.ClientAuthInfoWriter, ...visualizations.ClientOption) *visualizations.DeleteVisualizationsIDOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*visualizations.DeleteVisualizationsIDOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*visualizations.DeleteVisualizationsIDParams, runtime.ClientAuthInfoWriter, ...visualizations.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PostVisualizations provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) PostVisualizations(params *visualizations.PostVisualizationsParams, authInfo runtime.ClientAuthInfoWriter, opts ...visualizations.ClientOption) (*visualizations.PostVisualizationsOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *visualizations.PostVisualizationsOK
	if rf, ok := ret.Get(0).(func(*visualizations.PostVisualizationsParams, runtime.ClientAuthInfoWriter, ...visualizations.ClientOption) *visualizations.PostVisualizationsOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*visualizations.PostVisualizationsOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*visualizations.PostVisualizationsParams, runtime.ClientAuthInfoWriter, ...visualizations.ClientOption) error); ok {
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
