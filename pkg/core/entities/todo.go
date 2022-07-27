package entities

import (
	"time"
)

type Todo struct {
	Id           string
	Title        string
	Description  string
	LimitDate    time.Time
	AssignedUser string
	Status       string
}

func (t *Todo) UpdateWith(newTitle string, newDescription string, newLimitDate time.Time, newAssignedUser string) {
	t.Title = newTitle
	t.Description = newDescription
	t.LimitDate = newLimitDate
	t.AssignedUser = newAssignedUser
}

func (t *Todo) SetNewStatus(newStatus string) {
	t.Status = newStatus
}
