//go:generate go run github.com/99designs/gqlgen generate

package public

import (
	"github.com/kytruong0712/goffee-shop/api-gateway/internal/controller/user"

	"github.com/99designs/gqlgen/graphql"
)

func NewSchema(usrCtrl user.Controller) graphql.ExecutableSchema {
	cfg := Config{
		Resolvers: &resolver{
			usrCtrl: usrCtrl,
		},
	}

	return NewExecutableSchema(cfg)
}

type resolver struct {
	usrCtrl user.Controller
}

func (r *resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}

func (r *resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type queryResolver struct {
	*resolver
}

type mutationResolver struct {
	*resolver
}
