package mock_repository

import (
	context "context"
	reflect "reflect"
	"testing"

	gomock "github.com/golang/mock/gomock"
	model "github.com/hirac1220/go-clean-architecture/domain/model"
	mock "github.com/hirac1220/go-clean-architecture/mock_todo"
	"github.com/hirac1220/go-clean-architecture/usecase"
)

func TestPostTodo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var todo *model.Todo
	var id int64
	var expected *model.Todo
	var err error

	ctx := context.Background()

	mockSample := mock.NewMockTodoRepository(ctrl)
	mockSample.EXPECT().CreateTodo(ctx, todo).Return(id, err)

	uc := usecase.NewTodoUseCase(mockSample)
	result, err := uc.PostTodo(ctx, todo)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Actual PostTodo() is not same as expected")
	}
}
func TestGetTodo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var expected *model.Todo
	var err error

	ctx := context.Background()
	id := "1"

	mockSample := mock.NewMockTodoRepository(ctrl)
	mockSample.EXPECT().GetTodoById(ctx, 1).Return(expected, err)

	uc := usecase.NewTodoUseCase(mockSample)
	result, err := uc.GetTodo(ctx, id)

	if err != nil {
		t.Error("Actual GetTodo() is not same as expected")
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Actual GetTodo() is not same as expected")
	}
}
func TestPutTodo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var todo *model.Todo
	var id int64
	var expected *model.Affected
	var err error

	ctx := context.Background()

	mockSample := mock.NewMockTodoRepository(ctrl)
	mockSample.EXPECT().PutTodoById(ctx, 1, todo).Return(id, err)

	uc := usecase.NewTodoUseCase(mockSample)
	result, err := uc.PutTodo(ctx, "1", todo)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Actual PutTodo() is not same as expected")
	}
}
func TestDeleteTodo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var id int64
	var expected *model.Affected
	var err error

	ctx := context.Background()

	mockSample := mock.NewMockTodoRepository(ctrl)
	mockSample.EXPECT().DeleteTodoById(ctx, 1).Return(id, err)

	uc := usecase.NewTodoUseCase(mockSample)
	result, err := uc.DeleteTodo(ctx, "1")

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Actual DeleteTodo() is not same as expected")
	}
}
func TestListTodos(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var expected []model.Todo
	var err error

	ctx := context.Background()

	mockSample := mock.NewMockTodoRepository(ctrl)
	mockSample.EXPECT().ListTodos(ctx).Return(expected, err)

	uc := usecase.NewTodoUseCase(mockSample)
	result, err := uc.ListTodos(ctx)

	if err != nil {
		t.Error("Actual ListTodos() is not same as expected")
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Actual ListTodos() is not same as expected")
	}
}