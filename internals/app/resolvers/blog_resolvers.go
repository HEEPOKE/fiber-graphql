package resolvers

import (
	"context"
	"fmt"

	"github.com/HEEPOKE/fiber-graphql/internals/domains/generated"
	"github.com/HEEPOKE/fiber-graphql/internals/domains/models"
)

func (r *blogResolver) ID(ctx context.Context, obj *models.Blog) (string, error) {
	panic(fmt.Errorf("not implemented: ID - id"))
}

func (r *blogResolver) AccountID(ctx context.Context, obj *models.Blog) (string, error) {
	panic(fmt.Errorf("not implemented: AccountID - accountId"))
}

func (r *Resolver) Blog() generated.BlogResolver { return &blogResolver{r} }

type blogResolver struct{ *Resolver }
