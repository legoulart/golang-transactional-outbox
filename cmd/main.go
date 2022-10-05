package main

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
	"github.com/legoulart/poc-go/internal/domain/usecase"
	"github.com/legoulart/poc-go/primary/http"
	"github.com/legoulart/poc-go/secondary/mongodb"
)

func main() {
	e := echo.New()

	rep := mongodb.NewMongoDbRepository()
	dap := mongodb.NewMongoDbAdapter(rep)

	u := usecase.NewUserUseCase(dap)

	http.NewController(*e.Group(http.Path), u)

	e.Logger.Fatal(e.Start(":1323"))
}
