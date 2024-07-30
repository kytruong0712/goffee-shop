package main

import (
	"context"

	"github.com/go-chi/chi/v5"
)

type router struct {
	ctx         context.Context
	corsOrigins []string
}

func (rtr router) routes(r chi.Router) {
	r.Group(rtr.public)
}

func (rtr router) public(r chi.Router) {
}
