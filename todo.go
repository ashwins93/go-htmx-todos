package wutplans

import (
	"time"
)

type Todo struct {
	ID          string     `json:"id"`
	Task        string     `json:"task"`
	CompletedAt *time.Time `json:"competed_at"`
	CreatedAt   time.Time  `json:"created_at"`
}
