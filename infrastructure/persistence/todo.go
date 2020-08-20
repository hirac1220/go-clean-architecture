package persistence

import (
	"context"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/hirac1220/go-clean-architecture/domain/model"
	_ "github.com/hirac1220/go-clean-architecture/domain/repository"
)

func (tp *todoPersistence) Close() {
	tp.db.Close()
}

func (tp *todoPersistence) Ping() error {
	return tp.db.Ping()
}

func (tp *todoPersistence) CheckUserId(ctx context.Context, uid int) (int64, error) {
	row := tp.db.QueryRowContext(ctx,
		`SELECT id
		 FROM users
 		 WHERE id = ?;`,
		uid,
	)

	var id int64
	if err := row.Scan(&id); err != nil {
		log.Println(err)
		return 0, err
	}
	log.Printf("user_id: %v", id)
	return id, nil
}

// Post Todo
func (tp *todoPersistence) CreateTodo(ctx context.Context, uid int, todo *model.Todo) (int64, error) {
	stmt, err := tp.db.Prepare(
		`INSERT INTO todos
		 (user_id, deadline, todo)
		 VALUES (?, ?, ?);`,
	)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	res, err := stmt.Exec(uid, todo.Deadline, todo.Todo)
	if err != nil {
		return 0, err
	}
	log.Println("new todo created:")
	log.Println(res.LastInsertId())

	return res.LastInsertId()
}

// Get Todo
func (tp *todoPersistence) GetTodoById(ctx context.Context, uid, id int) (*model.Todo, error) {
	row := tp.db.QueryRowContext(ctx,
		`SELECT * 
		 FROM todos
 		 WHERE user_id = ? AND id = ?;`,
		uid,
		id,
	)

	var todo = &model.Todo{}
	if err := row.Scan(&todo.Id, &todo.UserId, &todo.Deadline, &todo.Todo); err != nil {
		log.Println(err)
		return nil, err
	}
	log.Printf("todo: %v", todo)
	return todo, nil
}

// Put Todo
func (tp *todoPersistence) PutTodoById(ctx context.Context, uid, id int, todo *model.Todo) (int64, error) {
	stmt, err := tp.db.Prepare(
		`UPDATE todos 
		 SET user_id = ?, deadline = ?, todo = ? 
		 WHERE user_id = ? AND id = ?;`,
	)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	res, err := stmt.Exec(uid, todo.Deadline, todo.Todo, uid, id)
	if err != nil {
		return 0, err
	}
	log.Println(res.RowsAffected())
	log.Printf("id: %d todo updated", id)

	return res.RowsAffected()
}

// Delete Todo
func (tp *todoPersistence) DeleteTodoById(ctx context.Context, uid, id int) (int64, error) {
	stmt, err := tp.db.Prepare(
		`DELETE FROM todos WHERE user_id = ? AND id = ?;`,
	)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	res, err := stmt.Exec(uid, id)
	if err != nil {
		return 0, err
	}
	log.Println(res.RowsAffected())
	log.Printf("id: %d todo deleted", id)

	return res.RowsAffected()
}

// List Todo
func (tp *todoPersistence) ListTodos(ctx context.Context, uid int) ([]model.Todo, error) {
	rows, err := tp.db.QueryContext(ctx,
		`SELECT t.id, t.user_id, t.deadline, t.todo
		 FROM todos t
		 JOIN users u ON t.user_id = u.id 
		 WHERE t.user_id = ?;`,
		uid,
	)
	if err != nil {
		return nil, err
	}

	list := []model.Todo{}
	for rows.Next() {
		var todo = &model.Todo{}
		if err := rows.Scan(&todo.Id, &todo.UserId, &todo.Deadline, &todo.Todo); err == nil {
			list = append(list, *todo)
		}
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	log.Printf("todo list: %d", list)
	return list, nil
}

func (tp *todoPersistence) GetNow() time.Time {
	location := time.FixedZone("Asia/Tokyo", 9*60*60)
	date := time.Now().In(location)
	const layout = "2006-01-02 15:04:05 -7:00"
	dt, _ := time.Parse(layout, date.Format(layout))
	return dt
}
