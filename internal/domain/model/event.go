package model

type Event struct {
	id            string
	aggregateType string
	payload       string //TODO: json type?
}

func EventFromUser(user User) Event {
	panic("TODO")
}
