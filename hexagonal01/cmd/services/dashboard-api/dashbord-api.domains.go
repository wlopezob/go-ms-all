package dashboard_api

import (
	"context"

	"github.com/wlopezob/hexagonal01/cmd/services/dashboard-api/models"
	ports "github.com/wlopezob/hexagonal01/cmd/services/dashboard-api/ports/drivens"
)

type DashboardApi struct {
	userQueryer ports.ForQueryingUser
}

func (d *DashboardApi) GetUser(ctx context.Context, username string) (*models.User, error) {
	return &models.User{}, nil
}

func NewDashboardService(userQueryer ports.ForQueryingUser) *DashboardApi {
	return &DashboardApi{userQueryer: userQueryer}
}
