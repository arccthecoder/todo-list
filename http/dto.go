package http

import (
	"encoding/json"
	"time"
)

// DTO == data transfer object
type TaskDTO struct {
	Title       string
	Description string
}

type ErrorDTO struct {
	Message string
	Time    time.Time
}

func (e ErrorDTO) ToString() string {
	b, err := json.MarshalIndent(e, "", "    ")
	if err != nil {
		panic(err)
	}

	return string(b)
}
