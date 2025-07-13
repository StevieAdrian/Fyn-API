package main

import (
	"log"

	"github.com/StevieAdrian/Fyn-API/auth-service/internal/bootstrap"
)

func main() {
	app := bootstrap.InitializeApp()

	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
