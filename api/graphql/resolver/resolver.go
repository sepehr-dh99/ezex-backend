package resolver

import (
	"github.com/ezex-io/ezex-gateway/api/graphql/gen"
	"github.com/ezex-io/ezex-gateway/internal/interactor/auth"
)

type Resolver struct {
	auth *auth.Auth
}

func NewResolver(auth *auth.Auth) *Resolver {
	return &Resolver{
		auth: auth,
	}
}

func (r *Resolver) Mutation() gen.MutationResolver {
	return &mutationResolver{r}
}

func (r *Resolver) Query() gen.QueryResolver {
	return &queryResolver{r}
}

type (
	mutationResolver struct{ *Resolver }
	queryResolver    struct{ *Resolver }
)
