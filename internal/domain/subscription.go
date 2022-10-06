package domain

import (
	"github.com/google/uuid"
	"time"
)

type Subscription struct {
	Id        uuid.UUID
	Email     string
	CreatedAt time.Time
}

func NewSubscription(email string) *Subscription {
	return &Subscription{
		Id:        uuid.New(),
		Email:     email,
		CreatedAt: time.Now(),
	}
}
