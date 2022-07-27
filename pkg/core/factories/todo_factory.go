package factories

import (
	"context"
	"log"

	"github.com/google/uuid"
	req "github.com/leoguilen/simple-go-api/pkg/api/models/requests"
	abstractions "github.com/leoguilen/simple-go-api/pkg/core/abstractions/external-services"
	"github.com/leoguilen/simple-go-api/pkg/core/entities"
	"github.com/leoguilen/simple-go-api/pkg/infra/services"
)

type TodoFactory struct {
	UserApiService abstractions.IUserApiService
}

func NewTodoFactory() *TodoFactory {
	return &TodoFactory{
		UserApiService: services.NewUserApiService(),
	}
}

func (f *TodoFactory) Create(ctx context.Context, r req.CreateTodoRequest) (*entities.Todo, error) {
	user, err := f.UserApiService.GetUserByEmail(ctx, r.AssignedTo)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	todoUser := entities.TodoAssignedUser{
		Id:       user.Id,
		UserName: user.UserName,
		Email:    user.Email,
	}

	var todoStatus string
	if r.Started {
		todoStatus = "IN_PROGRESS"
	} else {
		todoStatus = "PENDING"
	}

	return &entities.Todo{
		Id:           uuid.New().String(),
		Title:        r.Title,
		Description:  r.Description,
		LimitDate:    r.LimitDate,
		AssignedUser: todoUser.Email,
		Status:       todoStatus,
	}, nil
}
