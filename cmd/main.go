package main

import (
	"context"
	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
	"github.com/legoulart/poc-go/internal/port"
	"github.com/legoulart/poc-go/primary/http"
	"github.com/legoulart/poc-go/secondary/mongodb"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	e := echo.New()

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(""))
	database := client.Database("poc")
	session, err := client.StartSession()

	if err != nil {
		panic(err)
	}

	rep := mongodb.NewMongoDbRepository(database.Collection("users"), database.Collection("outbox"))
	dap := mongodb.NewMongoDbAdapter(rep, session)

	u := port.New(dap)

	http.NewController(*e.Group(http.Path), u)

	e.Logger.Fatal(e.Start(":1323"))
}
