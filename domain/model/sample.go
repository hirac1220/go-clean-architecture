package model

import (
	"encoding/json"
	"fmt"
	"time"
)

func MakeTodo(id int, dt time.Time) *Todo {
	todo := &Todo{
		Id:       id,
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
			"deadline": "2020-09-01T00:00:00Z",
			"todo": "test"
		},
		{
			"id": 2,
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
