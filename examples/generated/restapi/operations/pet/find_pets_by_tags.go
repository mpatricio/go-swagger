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

package pet

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-swagger/go-swagger/examples/generated/models"
	"github.com/go-swagger/go-swagger/httpkit/middleware"
)

// FindPetsByTagsHandlerFunc turns a function with the right signature into a find pets by tags handler
type FindPetsByTagsHandlerFunc func(FindPetsByTagsParams, *models.User) ([]models.Pet, error)

func (fn FindPetsByTagsHandlerFunc) Handle(params FindPetsByTagsParams, principal *models.User) ([]models.Pet, error) {
	return fn(params, principal)
}

// FindPetsByTagsHandler interface for that can handle valid find pets by tags params
type FindPetsByTagsHandler interface {
	Handle(FindPetsByTagsParams, *models.User) ([]models.Pet, error)
}

// NewFindPetsByTags creates a new http.Handler for the find pets by tags operation
func NewFindPetsByTags(ctx *middleware.Context, handler FindPetsByTagsHandler) *FindPetsByTags {
	return &FindPetsByTags{Context: ctx, Handler: handler}
}

/*
Finds Pets by tags

Muliple tags can be provided with comma seperated strings. Use tag1, tag2, tag3 for testing.
*/
type FindPetsByTags struct {
	Context *middleware.Context
	Params  FindPetsByTagsParams
	Handler FindPetsByTagsHandler
}

func (o *FindPetsByTags) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)

	uprinc, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	var principal *models.User
	if uprinc != nil {
		principal = uprinc.(*models.User) // this is really a models.User, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &o.Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res, err := o.Handler.Handle(o.Params, principal) // actually handle the request
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	o.Context.Respond(rw, r, route.Produces, route, res)

}
