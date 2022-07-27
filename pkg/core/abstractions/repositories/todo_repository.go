package abstractions

import (
	"context"

	"github.com/leoguilen/simple-go-api/pkg/core/entities"
)

type ITodoRepository interface {
	Insert(ctx context.Context, t entities.Todo) error
	GetById(ctx context.Context, todoId string) (*entities.Todo, error)
	GetAll(ctx context.Context) (*[]entities.Todo, error)
	Update(ctx context.Context, t entities.Todo) error
	UpdateStatus(ctx context.Context, todoId string, newStatus string) error
	Delete(ctx context.Context, todoId string) error
}
