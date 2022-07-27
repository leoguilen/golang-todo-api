package models

import "time"

// UpdateTodoRequest
//
// A request model to update todo item.
//
// swagger:model
type UpdateTodoRequest struct {
	// New todo title.
	// Example: Do anything
	Title string `json:"title" validate:"required,min=3,max=30"`
	// New todo description.
	// Example: Do anything after wake up.
	Description string `json:"description" validate:"required,min=3,max=100"`
	// New todo limitDate.
	// Example: 2022-07-25T09:00:00Z
	LimitDate time.Time `json:"limitDate" validate:"required"`
	// New todo assigned user.
	// Example: user@email.com
	AssignedTo string `json:"assignedTo" validate:"email"`
}
