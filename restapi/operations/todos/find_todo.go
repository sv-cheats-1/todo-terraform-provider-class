// Code generated by go-swagger; DO NOT EDIT.

package todos

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// FindTodoHandlerFunc turns a function with the right signature into a find todo handler
type FindTodoHandlerFunc func(FindTodoParams) middleware.Responder

// Handle executing the request and returning a response
func (fn FindTodoHandlerFunc) Handle(params FindTodoParams) middleware.Responder {
	return fn(params)
}

// FindTodoHandler interface for that can handle valid find todo params
type FindTodoHandler interface {
	Handle(FindTodoParams) middleware.Responder
}

// NewFindTodo creates a new http.Handler for the find todo operation
func NewFindTodo(ctx *middleware.Context, handler FindTodoHandler) *FindTodo {
	return &FindTodo{Context: ctx, Handler: handler}
}

/*FindTodo swagger:route GET /{id} todos findTodo

FindTodo find todo API

*/
type FindTodo struct {
	Context *middleware.Context
	Handler FindTodoHandler
}

func (o *FindTodo) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewFindTodoParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
