package domain

import (
	"encoding/json"
	"github.com/google/uuid"
)

type Event struct {
	Id      string
	Topic   string
	Message string //TODO: json type?
}

func EventFromSubscription(subscription Subscription) Event {
	message, err := json.Marshal(subscription)
	if err != nil {
		panic("TODO")
	}

	return Event{
		Id:      uuid.NewString(),
		Topic:   "email-subscribed",
		Message: string(message),
	}
}
