package repository

import (
	"context"
	"time"

	"github.com/hirac1220/go-clean-architecture/domain/model"
)

type TodoRepository interface {
	Close()
	CreateTodo(context.Context, *model.Todo) (int64, error)
	GetTodoById(context.Context, int) (*model.Todo, error)
	PutTodoById(context.Context, int, *model.Todo) (int64, error)
	DeleteTodoById(context.Context, int) (int64, error)
	ListTodos(context.Context) ([]model.Todo, error)
	GetNow() time.Time
}
