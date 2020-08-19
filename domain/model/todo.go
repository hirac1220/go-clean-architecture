package model

import (
	"time"
)

type Todo struct {
	Id       int       `json:"id"`
	UserId   int       `json:"user_id"`
	Deadline time.Time `json:"deadline"`
	Todo     string    `json:"todo"`
}

type Affected struct {
	Affected int `json:"affected"`
}
