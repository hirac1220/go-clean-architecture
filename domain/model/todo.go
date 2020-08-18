package model

import (
	"time"
)

type Todo struct {
	Id       int       `json:"id"`
	Deadline time.Time `json:"deadline"`
	Todo     string    `json:"todo"`
}

type Affected struct {
	Affected int `json:"affected"`
}
