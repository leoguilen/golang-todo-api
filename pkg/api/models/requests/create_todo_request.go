package models

import "time"

// CreateTodoRequest
//
// A request model to create todo item.
//
// swagger:model
type CreateTodoRequest struct {
	// Todo title.
	// Example: Do anything
	Title string `json:"title" validate:"required,min=3,max=30"`
	// Todo description.
	// Example: Do anything after wake up.
	Description string `json:"description" validate:"required,min=3,max=100"`
	// Todo limitDate.
	// Example: 2022-07-25T09:00:00Z
	LimitDate time.Time `json:"limitDate" validate:"required"`
	// Todo assigned user.
	// Example: user@email.com
	AssignedTo string `json:"assignedTo" validate:"email"`
	// Todo is started.
	// Example: false
	Started bool `json:"started" default:"false"`
}
