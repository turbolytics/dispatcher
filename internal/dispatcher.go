package internal

import "github.com/google/uuid"

type Channel struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type Subscription struct {
}
