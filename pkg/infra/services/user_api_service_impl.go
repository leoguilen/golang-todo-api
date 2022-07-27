package services

import (
	"context"

	"github.com/google/uuid"
	abstractions "github.com/leoguilen/simple-go-api/pkg/core/abstractions/external-services"
	"github.com/leoguilen/simple-go-api/pkg/core/models"
)

type UserApiService struct{}

func NewUserApiService() abstractions.IUserApiService {
	return &UserApiService{}
}

func (*UserApiService) GetUserByEmail(ctx context.Context, userEmail string) (*models.UserModel, error) {
	// TODO: Here, an integration can be implemented for another service

	mockedUser := models.UserModel{
		Id:       uuid.NewString(),
		Email:    userEmail,
		UserName: userEmail,
	}

	return &mockedUser, nil
}
