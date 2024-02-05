package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/samluiz/concurrency-control/internal/db/config"
	"github.com/samluiz/concurrency-control/internal/db/repositories"
	"github.com/samluiz/concurrency-control/internal/handlers"
)

func main() {
	// Opening the database connection
	db, err := config.OpenDB(); if err != nil {
		log.Fatal(err)
		panic(err)
	}
	// Defer the database connection close if error occurs
	defer db.Close()

	// Creating the repository and handler
	repo := repositories.NewRepo(db)
	handlers	:= handlers.NewHandler(repo)

	// Creating the Fiber app
	app := fiber.New()

	// Routes
	app.Post("/clientes/:id/transacoes", handlers.HandleCreateTransacao)
	app.Get("/clientes/:id/extrato", handlers.HandleGetExtrato)

	app.Listen(":3000")
}