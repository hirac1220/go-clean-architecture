package repository

import (
	"context"
	"time"

	"github.com/hirac1220/go-clean-architecture/domain/model"
)

type TodoRepository interface {
	Close()
	CheckUserId(context.Context, int) (int64, error)
	CreateTodo(context.Context, int, *model.Todo) (int64, error)
	GetTodoById(context.Context, int, int) (*model.Todo, error)
	PutTodoById(context.Context, int, int, *model.Todo) (int64, error)
	DeleteTodoById(context.Context, int, int) (int64, error)
	ListTodos(context.Context, int) ([]model.Todo, error)
	GetNow() time.Time
}
