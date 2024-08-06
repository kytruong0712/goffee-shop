package main

import (
	"context"
	"github.com/kytruong0712/goffee-shop/api-gateway/internal/controller/user"
	"github.com/kytruong0712/goffee-shop/api-gateway/internal/infra/protocols/gql"

	"github.com/kytruong0712/goffee-shop/api-gateway/internal/handler/gql/public"

	"github.com/go-chi/chi/v5"
)

type router struct {
	ctx         context.Context
	corsOrigins []string
	userCtrl    user.Controller
}

func (rtr router) routes(r chi.Router) {
	r.Group(rtr.public)
}

func (rtr router) public(r chi.Router) {
	const prefix = "/gateway/public"

	r.Handle(prefix+"/graphql", gql.Handler(public.NewSchema(rtr.userCtrl), true))
}
