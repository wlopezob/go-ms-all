package adapters

import (
	"context"

	"github.com/wlopezob/hexagonal01/cmd/services/repository"
	"github.com/wlopezob/hexagonal01/cmd/services/repository/models"
)

type UserManagerProxyAdapter struct {
	ctx context.Context
	repository *repository.Repository
}

func (a *UserManagerProxyAdapter) GetUser(username string) (*models.User, error) {
	return a.repository.GetUser(a.ctx, username)
}

func NewUserManagerProxyAdapter(ctx context.Context, repository *repository.Repository) UserManagerProxyAdapter {
	return UserManagerProxyAdapter{
		repository: repository,
		ctx: ctx,
	}
}