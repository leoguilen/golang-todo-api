package services

import (
	"context"
	"database/sql"
	"log"

	req "github.com/leoguilen/simple-go-api/pkg/api/models/requests"
	res "github.com/leoguilen/simple-go-api/pkg/api/models/responses"
	repos "github.com/leoguilen/simple-go-api/pkg/core/abstractions/repositories"
	"github.com/leoguilen/simple-go-api/pkg/core/factories"
	"github.com/leoguilen/simple-go-api/pkg/core/models"
	"github.com/leoguilen/simple-go-api/pkg/core/services"
	repos_impl "github.com/leoguilen/simple-go-api/pkg/infra/repositories"
)

type TodoService struct {
	TodoFactory    factories.TodoFactory
	TodoRepository repos.ITodoRepository
}

func NewTodoService() services.ITodoService {
	return &TodoService{
		TodoFactory:    *factories.NewTodoFactory(),
		TodoRepository: repos_impl.NewTodoRepository(),
	}
}

func (t *TodoService) CreateNew(ctx context.Context, r req.CreateTodoRequest) error {
	todo, err := t.TodoFactory.Create(ctx, r)
	if err != nil {
		return err
	}

	if err := t.TodoRepository.Insert(ctx, *todo); err != nil {
		return err
	}

	return nil
}

func (t *TodoService) Update(ctx context.Context, todoId string, r req.UpdateTodoRequest) (*res.TodoResponse, error) {
	todo, err := t.TodoRepository.GetById(ctx, todoId)
	if err != nil {
		switch {
		case err == sql.ErrNoRows:
			return nil, models.ErrTodoNotFound
		default:
			return nil, err
		}
	}

	todo.UpdateWith(r.Title, r.Description, r.LimitDate, r.AssignedTo)
	if err := t.TodoRepository.Update(ctx, *todo); err != nil {
		log.Printf("update resource failed: %v", err)
		return nil, err
	}

	return res.NewTodoResponseFrom(todo), nil
}

func (t *TodoService) SetStatus(ctx context.Context, todoId string, r req.SetTodoStatusRequest) (*res.TodoResponse, error) {
	todo, err := t.TodoRepository.GetById(ctx, todoId)
	if err != nil {
		switch {
		case err == sql.ErrNoRows:
			return nil, models.ErrTodoNotFound
		default:
			return nil, err
		}
	}

	if todo.Status == "DONE" {
		return nil, models.ErrTodoClosed
	}

	todo.SetNewStatus(r.Status)
	if err := t.TodoRepository.UpdateStatus(ctx, todoId, todo.Status); err != nil {
		log.Printf("update status failed: %v", err)
		return nil, err
	}

	return res.NewTodoResponseFrom(todo), nil
}

func (t *TodoService) GetDetails(ctx context.Context, todoId string) (*res.TodoResponse, error) {
	todo, err := t.TodoRepository.GetById(ctx, todoId)
	if err != nil {
		switch {
		case err == sql.ErrNoRows:
			return nil, models.ErrTodoNotFound
		default:
			return nil, err
		}
	}
	return res.NewTodoResponseFrom(todo), nil
}

func (t *TodoService) ListAll(ctx context.Context) (*[]res.TodoResponse, error) {
	todos, err := t.TodoRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	var r []res.TodoResponse
	for _, t := range *todos {
		r = append(r, *res.NewTodoResponseFrom(&t))
	}
	return &r, nil
}

func (t *TodoService) Delete(ctx context.Context, todoId string) error {
	todo, err := t.TodoRepository.GetById(ctx, todoId)
	if err != nil {
		switch {
		case err == sql.ErrNoRows:
			return models.ErrTodoNotFound
		default:
			return err
		}
	}

	if err := t.TodoRepository.Delete(ctx, todo.Id); err != nil {
		return err
	}
	return nil
}
