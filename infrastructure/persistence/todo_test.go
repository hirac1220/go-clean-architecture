package persistence

import (
	"context"
	"testing"
	"time"

	"github.com/hirac1220/go-clean-architecture/domain/model"
	persistence "github.com/hirac1220/go-clean-architecture/infrastructure/persistence"
	"github.com/stretchr/testify/assert"
)

var dt time.Time
var id int64

func TestCreateTodo(t *testing.T) {
	ctx := context.Background()
	persistence.SetConfig()
	tp, _ := persistence.NewTodoPersistence()

	dt = tp.GetNow()
	todo := model.MakeTodo(0, dt)
	id, _ = tp.CreateTodo(ctx, todo)
	expected := id
	assert.Equal(t, expected, id)
}
func TestGetTodo(t *testing.T) {
	ctx := context.Background()
	tp, _ := persistence.NewTodoPersistence()

	actual, _ := tp.GetTodoById(ctx, int(id))
	expected := model.MakeTodo(int(id), dt)
	assert.Equal(t, expected, actual)
}
func TestPutTodo(t *testing.T) {
	ctx := context.Background()
	tp, _ := persistence.NewTodoPersistence()

	dt = tp.GetNow()
	todo := model.MakeTodo(int(id), dt)
	_, _ = tp.PutTodoById(ctx, int(id), todo)
	actual, _ := tp.GetTodoById(ctx, int(id))
	assert.Equal(t, todo, actual)
}
func TestDeleteTodo(t *testing.T) {
	ctx := context.Background()
	tp, _ := persistence.NewTodoPersistence()

	dt = tp.GetNow()
	actual, _ := tp.DeleteTodoById(ctx, int(id))
	expected := int64(1)
	assert.Equal(t, expected, actual)
}
func TestListTodos(t *testing.T) {
	ctx := context.Background()
	tp, _ := persistence.NewTodoPersistence()

	actual, _ := tp.ListTodos(ctx)
	expected := model.ListTodos()
	assert.Equal(t, expected, actual)
}
