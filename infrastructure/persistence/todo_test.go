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

func TestCheckUserId(t *testing.T) {
	ctx := context.Background()
	persistence.SetConfig()
	tp, _ := persistence.NewTodoPersistence()

	uid := 1
	actual, _ := tp.CheckUserId(ctx, uid)
	expected := int64(1)
	assert.Equal(t, expected, actual)
}
func TestCreateTodo(t *testing.T) {
	ctx := context.Background()
	tp, _ := persistence.NewTodoPersistence()

	uid := 1
	dt = tp.GetNow()
	todo := model.MakeTodo(1, 0, dt)
	id, _ = tp.CreateTodo(ctx, uid, todo)
	expected := id
	assert.Equal(t, expected, id)
}
func TestGetTodo(t *testing.T) {
	ctx := context.Background()
	tp, _ := persistence.NewTodoPersistence()

	uid := 1
	actual, _ := tp.GetTodoById(ctx, uid, int(id))
	expected := model.MakeTodo(uid, int(id), dt)
	assert.Equal(t, expected, actual)
}
func TestPutTodo(t *testing.T) {
	ctx := context.Background()
	tp, _ := persistence.NewTodoPersistence()

	uid := 1
	dt = tp.GetNow()
	todo := model.MakeTodo(uid, int(id), dt)
	_, _ = tp.PutTodoById(ctx, uid, int(id), todo)
	actual, _ := tp.GetTodoById(ctx, uid, int(id))
	assert.Equal(t, todo, actual)
}
func TestDeleteTodo(t *testing.T) {
	ctx := context.Background()
	tp, _ := persistence.NewTodoPersistence()

	uid := 1
	dt = tp.GetNow()
	actual, _ := tp.DeleteTodoById(ctx, uid, int(id))
	expected := int64(1)
	assert.Equal(t, expected, actual)
}
func TestListTodos(t *testing.T) {
	ctx := context.Background()
	tp, _ := persistence.NewTodoPersistence()

	uid := 1
	actual, _ := tp.ListTodos(ctx, uid)
	expected := model.ListTodos()
	assert.Equal(t, expected, actual)
}
