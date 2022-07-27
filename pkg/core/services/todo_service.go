package services

import (
	"context"

	req "github.com/leoguilen/simple-go-api/pkg/api/models/requests"
	res "github.com/leoguilen/simple-go-api/pkg/api/models/responses"
)

type ITodoService interface {
	CreateNew(ctx context.Context, r req.CreateTodoRequest) error
	Update(ctx context.Context, todoId string, r req.UpdateTodoRequest) (*res.TodoResponse, error)
	SetStatus(ctx context.Context, todoId string, r req.SetTodoStatusRequest) (*res.TodoResponse, error)
	GetDetails(ctx context.Context, todoId string) (*res.TodoResponse, error)
	ListAll(ctx context.Context) (*[]res.TodoResponse, error)
	Delete(ctx context.Context, todoId string) error
}
