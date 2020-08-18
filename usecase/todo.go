package usecase

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strconv"

	"github.com/hirac1220/go-clean-architecture/domain/model"
	"github.com/hirac1220/go-clean-architecture/domain/repository"
)

var (
	ErrNotFound   = errors.New("data not found")
	ErrNoAffected = errors.New("data not updated/inserted")
)

type TodoUsecase interface {
	PostTodo(context.Context, *model.Todo) (*model.Todo, error)
	GetTodo(context.Context, string) (*model.Todo, error)
	PutTodo(context.Context, string, *model.Todo) (*model.Affected, error)
	DeleteTodo(context.Context, string) (*model.Affected, error)
	ListTodos(context.Context) ([]model.Todo, error)
}

type todoUseCase struct {
	todoRepository repository.TodoRepository
}

func NewTodoUseCase(tr repository.TodoRepository) TodoUsecase {
	return &todoUseCase{
		todoRepository: tr,
	}
}

func (tu *todoUseCase) PostTodo(ctx context.Context, todo *model.Todo) (*model.Todo, error) {
	var err error
	id, err := tu.todoRepository.CreateTodo(ctx, todo)
	if id == 0 {
		return nil, fmt.Errorf("error: %w", ErrNoAffected)
	} else if err != nil {
		return nil, fmt.Errorf("error: %w", err)
	}
	t := todo
	t.Id = int(id)
	return t, nil
}

func (tu *todoUseCase) GetTodo(ctx context.Context, id string) (*model.Todo, error) {
	i, _ := strconv.Atoi(id)
	var err error
	t, err := tu.todoRepository.GetTodoById(ctx, i)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("id: %v error: %w", id, ErrNotFound)
	} else if err != nil {
		return nil, fmt.Errorf("id: %v error: %w", id, err)
	}
	return t, nil
}

func (tu *todoUseCase) PutTodo(ctx context.Context, id string, todo *model.Todo) (*model.Affected, error) {
	i, _ := strconv.Atoi(id)
	var err error
	affected, err := tu.todoRepository.PutTodoById(ctx, i, todo)
	if affected == 0 {
		return nil, fmt.Errorf("id: %v error: %w", id, ErrNoAffected)
	} else if err != nil {
		return nil, fmt.Errorf("id: %v error: %w", id, err)
	}
	a := &model.Affected{
		Affected: int(affected),
	}
	return a, nil
}

func (tu *todoUseCase) DeleteTodo(ctx context.Context, id string) (*model.Affected, error) {
	i, _ := strconv.Atoi(id)
	var err error
	affected, err := tu.todoRepository.DeleteTodoById(ctx, i)
	if affected == 0 {
		return nil, fmt.Errorf("id: %v error: %w", id, ErrNoAffected)
	} else if err != nil {
		return nil, fmt.Errorf("id: %v error: %w", id, err)
	}
	a := &model.Affected{
		Affected: int(affected),
	}
	return a, nil
}

func (tu *todoUseCase) ListTodos(ctx context.Context) ([]model.Todo, error) {
	var err error
	tlist, err := tu.todoRepository.ListTodos(ctx)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("error: %w", ErrNotFound)
	} else if err != nil {
		return nil, fmt.Errorf("error: %w", err)
	}
	return tlist, nil
}
