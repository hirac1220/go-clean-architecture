package model

import (
	"encoding/json"
	"fmt"
	"time"
)

func MakeTodo(uid, id int, dt time.Time) *Todo {
	todo := &Todo{
		Id:       id,
		UserId:   uid,
		Deadline: dt,
		Todo:     "test4",
	}

	return todo
}

func ListTodos() []Todo {
	var list string = `
	[
		{
			"id": 1,
			"user_id": 1,
			"deadline": "2020-09-01T00:00:00Z",
			"todo": "test1"
		},
		{
			"id": 2,
			"user_id": 1,
			"deadline": "2020-09-01T00:00:00Z",
			"todo": "test2"
		}
	]`

	tlist := []Todo{}
	if err := json.Unmarshal([]byte(list), &tlist); err != nil {
		fmt.Println(err)
		return nil
	}
	return tlist
}
