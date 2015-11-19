// Copyright 2015 go-swagger maintainers
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package store

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-swagger/go-swagger/examples/generated/models"
	"github.com/go-swagger/go-swagger/httpkit/middleware"
)

// GetOrderByIDHandlerFunc turns a function with the right signature into a get order by id handler
type GetOrderByIDHandlerFunc func(GetOrderByIDParams) (*models.Order, error)

func (fn GetOrderByIDHandlerFunc) Handle(params GetOrderByIDParams) (*models.Order, error) {
	return fn(params)
}

// GetOrderByIDHandler interface for that can handle valid get order by id params
type GetOrderByIDHandler interface {
	Handle(GetOrderByIDParams) (*models.Order, error)
}

// NewGetOrderByID creates a new http.Handler for the get order by id operation
func NewGetOrderByID(ctx *middleware.Context, handler GetOrderByIDHandler) *GetOrderByID {
	return &GetOrderByID{Context: ctx, Handler: handler}
}

/*
Find purchase order by ID

For valid response try integer IDs with value <= 5 or > 10. Other values will generated exceptions
*/
type GetOrderByID struct {
	Context *middleware.Context
	Params  GetOrderByIDParams
	Handler GetOrderByIDHandler
}

func (o *GetOrderByID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)

	if err := o.Context.BindValidRequest(r, route, &o.Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res, err := o.Handler.Handle(o.Params) // actually handle the request
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	o.Context.Respond(rw, r, route.Produces, route, res)

}
