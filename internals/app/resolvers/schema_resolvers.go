package resolvers

import (
	"context"
	"fmt"

	"github.com/HEEPOKE/fiber-graphql/internals/domains/generated"
	"github.com/HEEPOKE/fiber-graphql/internals/domains/models"
)

func (r *queryResolver) GetAllAccounts(ctx context.Context) ([]*models.Account, error) {
	panic(fmt.Errorf("not implemented: GetAllAccounts - getAllAccounts"))
}

func (r *queryResolver) GetAccount(ctx context.Context, id string) (*models.Account, error) {
	panic(fmt.Errorf("not implemented: GetAccount - getAccount"))
}

func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
