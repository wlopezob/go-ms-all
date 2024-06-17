package adapters

import (
	"context"

	dashboard_api "github.com/wlopezob/hexagonal01/cmd/services/dashboard-api"
	"github.com/wlopezob/hexagonal01/cmd/services/dashboard-api/models"
)

type UserAdapter struct {
	ctx         context.Context
	dasboardApi *dashboard_api.DashboardApi
}

func (a *UserAdapter) GetUser(username string) (*models.User, error) {
	return a.dasboardApi.GetUser(a.ctx, username)
}

func CreateUserAdapter(ctx context.Context, dasboardApi *dashboard_api.DashboardApi) *UserAdapter {
	return &UserAdapter{
		ctx:         ctx,
		dasboardApi: dasboardApi,
	}
}
