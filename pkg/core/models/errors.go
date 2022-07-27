package models

import "errors"

var ErrTodoNotFound = errors.New("todo not exists")
var ErrTodoClosed = errors.New("todo is closed")
