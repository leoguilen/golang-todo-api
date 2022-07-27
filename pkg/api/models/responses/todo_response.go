package models

import (
	"time"

	"github.com/leoguilen/simple-go-api/pkg/core/entities"
)

// TodoResponse
//
// A TodoResponse is a model to return Todo data to client.
//
// swagger:model
type TodoResponse struct {
	// Id of the TODO
	// Example: 421f701e-7ae8-4a9e-a40f-a6ad9ffd478b
	Id string `json:"id"`
	// Title of the TODO
	// Example: Do something
	Title string `json:"title"`
	// Description of the TODO
	// Example: Do something today
	Description string `json:"description"`
	// LimitDate of the TODO
	// Example: 2022-07-24T23:00:00
	LimitDate time.Time `json:"limitDate"`
	// AssignedTo of the TODO
	// Example: user@email.com
	AssignedTo string `json:"assignedTo"`
	// Status of the TODO
	// Example: PENDING
	Status string `json:"status"`
}

func NewTodoResponseFrom(t *entities.Todo) *TodoResponse {
	return &TodoResponse{
		Id:          t.Id,
		Title:       t.Title,
		Description: t.Description,
		LimitDate:   t.LimitDate,
		AssignedTo:  t.AssignedUser,
		Status:      t.Status,
	}
}
