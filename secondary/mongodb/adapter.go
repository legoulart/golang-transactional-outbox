package mongodb

import "github.com/legoulart/poc-go/internal/domain/model"

type MongoDbAdapter struct {
	repository MongoDbRepository
}

func NewMongoDbAdapter(repository MongoDbRepository) MongoDbAdapter {
	return MongoDbAdapter{
		repository: repository,
	}
}

func (a MongoDbAdapter) Save(user model.User) {
	panic("TODO")
}

func (a MongoDbAdapter) SaveEvent(event model.Event) {
	panic("TODO")
}
