package wutplans

import (
	"context"
	"time"
)

type Todo struct {
	ID          string     `json:"id"`
	Task        string     `json:"task" form:"task"`
	CompletedAt *time.Time `json:"competed_at"`
	CreatedAt   time.Time  `json:"created_at"`
}

type UpdateTodo struct {
	Task        string     `json:"task,omitempty"`
	CompletedAt *time.Time `json:"completed_at,omitempty"`
}

type TodoService interface {
	CreateTodo(context.Context, *Todo) (*Todo, error)
	DeleteTodo(context.Context, string) (*Todo, error)
	FindAllTodos(context.Context) ([]*Todo, error)
	FindTodoById(context.Context, string) (*Todo, error)
	UpdateTodo(context.Context, string, Todo) (*Todo, error)
}
