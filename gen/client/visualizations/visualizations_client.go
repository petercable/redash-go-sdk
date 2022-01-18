// Code generated by go-swagger; DO NOT EDIT.

package visualizations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new visualizations API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for visualizations API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteVisualizationsID(params *DeleteVisualizationsIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteVisualizationsIDOK, error)

	PostVisualizations(params *PostVisualizationsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostVisualizationsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  DeleteVisualizationsID delete visualizations ID API
*/
func (a *Client) DeleteVisualizationsID(params *DeleteVisualizationsIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteVisualizationsIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteVisualizationsIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteVisualizationsID",
		Method:             "DELETE",
		PathPattern:        "/visualizations/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/x-www-form-urlencoded"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteVisualizationsIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteVisualizationsIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteVisualizationsIDDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  PostVisualizations post visualizations API
*/
func (a *Client) PostVisualizations(params *PostVisualizationsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostVisualizationsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostVisualizationsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostVisualizations",
		Method:             "POST",
		PathPattern:        "/visualizations",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostVisualizationsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostVisualizationsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PostVisualizationsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
