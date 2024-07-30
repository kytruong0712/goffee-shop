//go:generate go run github.com/99designs/gqlgen generate

package public

import "github.com/99designs/gqlgen/graphql"

func NewSchema() graphql.ExecutableSchema {
	cfg := Config{
		Resolvers: &resolver{},
	}

	return NewExecutableSchema(cfg)
}

type resolver struct {
}

func (r *resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}

type queryResolver struct {
	*resolver
}

type mutationResolver struct {
	*resolver
}
