package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/samluiz/concurrency-control/internal/db/config"
	"github.com/samluiz/concurrency-control/internal/db/repositories"
	"github.com/samluiz/concurrency-control/internal/handlers"
)

func main() {

	db, err := config.OpenDB(); if err != nil {
		panic(err)
	}
	defer db.Close()

	repo := repositories.NewRepo(db)

	handlers	:= handlers.NewHandler(repo)

	app := fiber.New()

	app.Post("/clientes/:id/transacoes", handlers.HandleCreateTransacao)
}