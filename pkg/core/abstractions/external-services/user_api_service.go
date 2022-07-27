package abstractions

import (
	"context"

	"github.com/leoguilen/simple-go-api/pkg/core/models"
)

type IUserApiService interface {
	GetUserByEmail(ctx context.Context, userEmail string) (*models.UserModel, error)
}
