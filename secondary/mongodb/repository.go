package mongodb

import (
	"context"
	"encoding/json"
	"github.com/legoulart/poc-go/internal/domain"

	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDbRepository struct {
	modelCollection  *mongo.Collection
	outboxCollection *mongo.Collection
}

func NewMongoDbRepository(modelCollection *mongo.Collection, outboxCollection *mongo.Collection) MongoDbRepository {
	return MongoDbRepository{
		modelCollection:  modelCollection,
		outboxCollection: outboxCollection,
	}
}

func (r MongoDbRepository) save(user domain.Subscription) {
	model, err := json.Marshal(user)

	if err != nil {
		panic("TODO")
	}

	_, err = r.modelCollection.InsertOne(context.TODO(), model)
	if err != nil {
		panic("TODO")
	}
}

func (r MongoDbRepository) saveEvent(event domain.Event) {
	model, err := json.Marshal(event)

	if err != nil {
		panic("TODO")
	}

	_, err = r.modelCollection.InsertOne(context.TODO(), model)
	if err != nil {
		panic("TODO")
	}
}
