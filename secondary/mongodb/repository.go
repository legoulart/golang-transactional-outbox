package mongodb

import (
	"encoding/json"
	"github.com/legoulart/poc-go/internal/domain/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDbRepository struct {
	context          mongo.SessionContext
	modelCollection  mongo.Collection
	outboxCollection mongo.Collection
}

func NewMongoDbRepository() MongoDbRepository {
	return MongoDbRepository{}
}

func (r MongoDbRepository) save(user model.User) {
	model, _ := json.Marshal(user)
	r.modelCollection.InsertOne(r.context, model)
}

func (r MongoDbRepository) saveEvent(event model.Event) {
	model, _ := json.Marshal(event)
	r.modelCollection.InsertOne(r.context, model)
}
