package resolvers

import (
	"context"
	"fmt"

	"github.com/HEEPOKE/fiber-graphql/internals/domains/generated"
	"github.com/HEEPOKE/fiber-graphql/internals/domains/generated/model"
	"github.com/HEEPOKE/fiber-graphql/internals/domains/models"
)

func (r *accountResolver) ID(ctx context.Context, obj *models.Account) (string, error) {
	panic(fmt.Errorf("not implemented: ID - id"))
}

func (r *accountResolver) Role(ctx context.Context, obj *models.Account) (model.Role, error) {
	panic(fmt.Errorf("not implemented: Role - role"))
}

func (r *Resolver) Account() generated.AccountResolver { return &accountResolver{r} }

type accountResolver struct{ *Resolver }
