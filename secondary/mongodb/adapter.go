package mongodb

import (
	"context"
	"github.com/legoulart/poc-go/internal/domain"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDbAdapter struct {
	repository MongoDbRepository
	session    mongo.Session
}

func NewMongoDbAdapter(repository MongoDbRepository, session mongo.Session) MongoDbAdapter {
	return MongoDbAdapter{repository: repository, session: session}
}

func (a MongoDbAdapter) Save(subscription domain.Subscription) {
	saveExecute := func(sessCtx mongo.SessionContext) (interface{}, error) {
		a.repository.save(subscription)
		a.repository.saveEvent(domain.EventFromSubscription(subscription))
		return nil, nil
	}

	_, err := a.session.WithTransaction(context.TODO(), saveExecute)
	if err != nil {
		panic("TODO")
	}
}
