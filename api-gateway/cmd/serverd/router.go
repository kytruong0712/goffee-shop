package main

import (
	"context"

	"github.com/kytruong0712/goffee-shop/api-gateway/internal/controller/user"
	"github.com/kytruong0712/goffee-shop/api-gateway/internal/handler/gql/authenticated"
	"github.com/kytruong0712/goffee-shop/api-gateway/internal/handler/gql/public"
	"github.com/kytruong0712/goffee-shop/api-gateway/internal/infra/iam"
	"github.com/kytruong0712/goffee-shop/api-gateway/internal/infra/protocols/gql"

	"github.com/go-chi/chi/v5"
)

type router struct {
	ctx         context.Context
	corsOrigins []string
	userCtrl    user.Controller
}

func (rtr router) routes(r chi.Router) {
	r.Group(rtr.public)
	r.Group(rtr.authenticated)
}

func (rtr router) public(r chi.Router) {
	const prefix = "/gateway/public"

	r.Handle(prefix+"/graphql", gql.Handler(public.NewSchema(rtr.userCtrl), true))
}

func (rtr router) authenticated(r chi.Router) {
	const prefix = "/gateway/authenticated"

	r.Use(iam.AuthenticateUserMiddleware(rtr.ctx))

	r.Handle(prefix+"/graphql", gql.Handler(authenticated.NewSchema(rtr.userCtrl), true))
}
