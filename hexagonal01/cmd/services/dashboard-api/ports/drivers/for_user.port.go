package ports

import (
	"github.com/wlopezob/hexagonal01/cmd/services/dashboard-api/models"
)

type ForUser interface {
	GetUser(username string)(*models.User, error)
}