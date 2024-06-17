package main

import (
	"context"

	dashboard_api "github.com/wlopezob/hexagonal01/cmd/services/dashboard-api"
	adapters "github.com/wlopezob/hexagonal01/cmd/services/dashboard-api/adapters/drivens"
	dashboard_adapters "github.com/wlopezob/hexagonal01/cmd/services/dashboard-api/adapters/drivers"
	dashboard_ports "github.com/wlopezob/hexagonal01/cmd/services/dashboard-api/ports/drivers"
	"github.com/wlopezob/hexagonal01/cmd/services/repository"
	repository_adapters "github.com/wlopezob/hexagonal01/cmd/services/repository/adapters/drivers"
)

func Compose() (dashboard_ports.ForUser) {
	// create ctx(Context)
	ctx := context.Background()

	// create repository
	repository := repository.NewRepository()

	// create repository drivers
	userManagerProxyAdapter := repository_adapters.NewUserManagerProxyAdapter(ctx, repository)

	// create dashboard api drivens
	userQueryerAdapter := adapters.NewUserQueryerAdapter(ctx, &userManagerProxyAdapter)

	// Create dashboard api
	dashboardApi := dashboard_api.NewDashboardService(userQueryerAdapter)

	// create dashboard api divers
	userAdapter := dashboard_adapters.CreateUserAdapter(ctx, dashboardApi)

	return userAdapter
}

func ComposeMock() (dashboard_ports.ForUser) {
	// create ctx(Context)
	ctx := context.Background()

	// create repository
	repository := repository.NewRepository()

	// create repository drivers
	userManagerProxyAdapter := repository_adapters.NewUserManagerProxyAdapter(ctx, repository)

	// create dashboard api drivens
	userQueryerAdapter := adapters.NewUserQueryerMockAdapter(ctx, &userManagerProxyAdapter)

	// Create dashboard api
	dashboardApi := dashboard_api.NewDashboardService(userQueryerAdapter)

	// create dashboard api divers
	userAdapter := dashboard_adapters.CreateUserAdapter(ctx, dashboardApi)

	return userAdapter
}