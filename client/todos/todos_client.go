// Code generated by go-swagger; DO NOT EDIT.

package todos

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new todos API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for todos API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
AddOne add one API
*/
func (a *Client) AddOne(params *AddOneParams) (*AddOneCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddOneParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "addOne",
		Method:             "POST",
		PathPattern:        "/",
		ProducesMediaTypes: []string{"application/spkane.todo-list.v1+json"},
		ConsumesMediaTypes: []string{"application/spkane.todo-list.v1+json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &AddOneReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AddOneCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*AddOneDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
DestroyOne destroy one API
*/
func (a *Client) DestroyOne(params *DestroyOneParams) (*DestroyOneNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDestroyOneParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "destroyOne",
		Method:             "DELETE",
		PathPattern:        "/{id}",
		ProducesMediaTypes: []string{"application/spkane.todo-list.v1+json"},
		ConsumesMediaTypes: []string{"application/spkane.todo-list.v1+json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DestroyOneReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DestroyOneNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DestroyOneDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
FindTodo find todo API
*/
func (a *Client) FindTodo(params *FindTodoParams) (*FindTodoOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFindTodoParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "findTodo",
		Method:             "GET",
		PathPattern:        "/{id}",
		ProducesMediaTypes: []string{"application/spkane.todo-list.v1+json"},
		ConsumesMediaTypes: []string{"application/spkane.todo-list.v1+json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &FindTodoReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*FindTodoOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*FindTodoDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
FindTodos find todos API
*/
func (a *Client) FindTodos(params *FindTodosParams) (*FindTodosOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFindTodosParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "findTodos",
		Method:             "GET",
		PathPattern:        "/",
		ProducesMediaTypes: []string{"application/spkane.todo-list.v1+json"},
		ConsumesMediaTypes: []string{"application/spkane.todo-list.v1+json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &FindTodosReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*FindTodosOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*FindTodosDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
UpdateOne update one API
*/
func (a *Client) UpdateOne(params *UpdateOneParams) (*UpdateOneOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateOneParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateOne",
		Method:             "PUT",
		PathPattern:        "/{id}",
		ProducesMediaTypes: []string{"application/spkane.todo-list.v1+json"},
		ConsumesMediaTypes: []string{"application/spkane.todo-list.v1+json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateOneReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateOneOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UpdateOneDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
