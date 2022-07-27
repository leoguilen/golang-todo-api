package models

// SetTodoStatusRequest
//
// A request model to set new todo status.
//
// swagger:model
type SetTodoStatusRequest struct {
	// New todo status. Allowed status: (PENDING, IN_PROGRESS, DONE)
	// Example: IN_PROGRESS
	Status string `json:"status" validate:"required,oneof=PENDING IN_PROGRESS DONE"`
}
