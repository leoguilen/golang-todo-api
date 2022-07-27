package repositories

import (
	"context"
	"database/sql"
	"fmt"

	repos "github.com/leoguilen/simple-go-api/pkg/core/abstractions/repositories"
	"github.com/leoguilen/simple-go-api/pkg/core/entities"
	"github.com/leoguilen/simple-go-api/pkg/infra/constants"
	dbcontext "github.com/leoguilen/simple-go-api/pkg/infra/context"
)

type TodoRepository struct {
	DB *sql.DB
}

func NewTodoRepository() repos.ITodoRepository {
	db, err := dbcontext.NewDbContext().GetConnection()
	if err != nil {
		panic(err)
	}
	return &TodoRepository{
		DB: db,
	}
}

func (r *TodoRepository) Insert(ctx context.Context, t entities.Todo) error {
	conn, err := r.DB.Conn(ctx)
	if err != nil {
		return err
	}

	defer conn.Close()
	res, err := conn.ExecContext(ctx, constants.InsertTodoSql, t.Id, t.Title, t.Description, t.LimitDate, t.AssignedUser, t.Status)
	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}

	fmt.Printf("Affected rows %v", rows)
	return nil
}

func (r *TodoRepository) GetById(ctx context.Context, todoId string) (*entities.Todo, error) {
	conn, err := r.DB.Conn(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	var todo entities.Todo
	rows, err := conn.QueryContext(ctx, constants.SelectSingleTodoSql, todoId)
	if err != nil {
		return nil, err
	}

	if !rows.Next() {
		return nil, sql.ErrNoRows
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	if err := rows.Scan(&todo.Id, &todo.Title, &todo.Description, &todo.LimitDate, &todo.AssignedUser, &todo.Status); err != nil {
		switch {
		case err == sql.ErrNoRows:
			return nil, sql.ErrNoRows
		default:
			return nil, err
		}
	}
	defer rows.Close()

	return &todo, nil
}

func (r *TodoRepository) GetAll(ctx context.Context) (*[]entities.Todo, error) {
	conn, err := r.DB.Conn(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	var todos []entities.Todo
	rows, err := conn.QueryContext(ctx, constants.SelectAllTodosSql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var todo entities.Todo
		if err := rows.Scan(&todo.Id, &todo.Title, &todo.Description, &todo.LimitDate, &todo.AssignedUser, &todo.Status); err != nil {
			switch {
			case err == sql.ErrNoRows:
				return nil, sql.ErrNoRows
			default:
				return nil, err
			}
		}
		todos = append(todos, todo)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &todos, nil
}

func (r *TodoRepository) Update(ctx context.Context, t entities.Todo) error {
	conn, err := r.DB.Conn(ctx)
	if err != nil {
		return err
	}

	defer conn.Close()
	res, err := conn.ExecContext(ctx, constants.UpdateTodoSql, t.Title, t.Description, t.LimitDate, t.AssignedUser, t.Status, t.Id)
	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}

	fmt.Printf("Affected rows %v", rows)
	return nil
}

func (r *TodoRepository) UpdateStatus(ctx context.Context, todoId string, newStatus string) error {
	conn, err := r.DB.Conn(ctx)
	if err != nil {
		return err
	}

	defer conn.Close()
	res, err := conn.ExecContext(ctx, constants.UpdateStatusSql, newStatus, todoId)
	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}

	fmt.Printf("Affected rows %v", rows)
	return nil
}

func (r *TodoRepository) Delete(ctx context.Context, todoId string) error {
	conn, err := r.DB.Conn(ctx)
	if err != nil {
		return err
	}

	defer conn.Close()
	res, err := conn.ExecContext(ctx, constants.DeleteTodoSql, todoId)
	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}

	fmt.Printf("Affected rows %v", rows)
	return nil
}
