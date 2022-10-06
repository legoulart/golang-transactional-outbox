package main

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/legoulart/poc-go/config/di"
)

func main() {
	di.Setup()
}
