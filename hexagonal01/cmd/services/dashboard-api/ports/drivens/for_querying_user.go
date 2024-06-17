package ports

import "github.com/wlopezob/hexagonal01/cmd/services/dashboard-api/models"

type ForQueryingUser interface {
	GetUser(username string) (*models.User, error)
}