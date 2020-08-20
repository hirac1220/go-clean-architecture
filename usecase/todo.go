package usecase

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/hirac1220/go-clean-architecture/domain/model"
	"github.com/hirac1220/go-clean-architecture/domain/repository"
)

var (
	ErrNotFound   = errors.New("data not found")
	ErrNoAffected = errors.New("data not updated/inserted")
)

type TodoUsecase interface {
	CheckUser(context.Context, string) (int, error)
	PostTodo(context.Context, string, *model.Todo) (*model.Todo, error)
	GetTodo(context.Context, string, string) (*model.Todo, error)
	PutTodo(context.Context, string, string, *model.Todo) (*model.Affected, error)
	DeleteTodo(context.Context, string, string) (*model.Affected, error)
	ListTodos(context.Context, string) ([]model.Todo, error)
}

type todoUseCase struct {
	todoRepository repository.TodoRepository
}

func NewTodoUseCase(tr repository.TodoRepository) TodoUsecase {
	return &todoUseCase{
		todoRepository: tr,
	}
}

func (tu *todoUseCase) CheckUser(ctx context.Context, user_id string) (int, error) {
	uid, _ := strconv.Atoi(user_id)
	var err error
	log.Println(uid)
	id, err := tu.todoRepository.CheckUserId(ctx, uid)
	if err == sql.ErrNoRows {
		return 0, fmt.Errorf("error: %w", ErrNotFound)
	} else if err != nil {
		return 0, fmt.Errorf("error: %w", err)
	}
	return int(id), nil
}

func (tu *todoUseCase) PostTodo(ctx context.Context, user_id string, todo *model.Todo) (*model.Todo, error) {
	uid, _ := strconv.Atoi(user_id)
	var err error
	id, err := tu.todoRepository.CreateTodo(ctx, uid, todo)
	if id == 0 {
		return nil, fmt.Errorf("error: %w", ErrNoAffected)
	} else if err != nil {
		return nil, fmt.Errorf("error: %w", err)
	}
	t := todo
	t.Id = int(id)
	return t, nil
}

func (tu *todoUseCase) GetTodo(ctx context.Context, user_id, id string) (*model.Todo, error) {
	uid, _ := strconv.Atoi(user_id)
	i, _ := strconv.Atoi(id)
	var err error
	t, err := tu.todoRepository.GetTodoById(ctx, uid, i)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("id: %v error: %w", id, ErrNotFound)
	} else if err != nil {
		return nil, fmt.Errorf("id: %v error: %w", id, err)
	}
	return t, nil
}

func (tu *todoUseCase) PutTodo(ctx context.Context, user_id, id string, todo *model.Todo) (*model.Affected, error) {
	uid, _ := strconv.Atoi(user_id)
	i, _ := strconv.Atoi(id)
	var err error
	affected, err := tu.todoRepository.PutTodoById(ctx, uid, i, todo)
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

func (tu *todoUseCase) DeleteTodo(ctx context.Context, user_id, id string) (*model.Affected, error) {
	uid, _ := strconv.Atoi(user_id)
	i, _ := strconv.Atoi(id)
	var err error
	affected, err := tu.todoRepository.DeleteTodoById(ctx, uid, i)
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

func (tu *todoUseCase) ListTodos(ctx context.Context, user_id string) ([]model.Todo, error) {
	uid, _ := strconv.Atoi(user_id)
	var err error
	tlist, err := tu.todoRepository.ListTodos(ctx, uid)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("error: %w", ErrNotFound)
	} else if err != nil {
		return nil, fmt.Errorf("error: %w", err)
	}
	return tlist, nil
}
