package adapters

import (
	"context"

	"github.com/wlopezob/hexagonal01/cmd/services/dashboard-api/models"
	repository_drivers "github.com/wlopezob/hexagonal01/cmd/services/repository/ports/drivers"
)

type UserQueryerMockAdapter struct {
	ctx    context.Context
	driver repository_drivers.ForManagingUser
}

func (a *UserQueryerMockAdapter) GetUser(username string) (*models.User, error) {
	mockedUser := models.User{
		Username: "mockedUser",
		Password: "mockedPassword",
	}

	return &mockedUser, nil
}

func NewUserQueryerMockAdapter(ctx context.Context, driver repository_drivers.ForManagingUser) *UserQueryerMockAdapter {
	return &UserQueryerMockAdapter{
		ctx:    ctx,
		driver: driver,
	}
}
