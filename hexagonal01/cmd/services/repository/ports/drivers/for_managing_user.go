package ports

import (
	"github.com/wlopezob/hexagonal01/cmd/services/repository/models"
)

type ForManagingUser interface {
	GetUser(username string) (*models.User, error)
}
