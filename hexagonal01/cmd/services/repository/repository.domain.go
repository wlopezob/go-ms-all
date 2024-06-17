package repository

import (
	"context"

	"github.com/wlopezob/hexagonal01/cmd/services/repository/models"
)

type Repository struct {

}

func (r *Repository) GetUser(ctx context.Context, username string) (*models.User, error) {
	return &models.User{}, nil
}

func NewRepository() *Repository {
	return &Repository{}
}