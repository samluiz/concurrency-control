package main

import (
	"os"

	"github.com/goccy/go-json"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/samluiz/concurrency-control/internal/db/config"
	"github.com/samluiz/concurrency-control/internal/db/repositories"
	"github.com/samluiz/concurrency-control/internal/handlers"
)

func main() {
	// Opening the database connection
	db, err := config.OpenDB(); if err != nil {
		os.Exit(1)
	}
	// Defer the database connection close if error occurs
	defer db.Close()

	// Creating the repository and handler
	repo := repositories.NewRepo(db)
	handlers	:= handlers.NewHandler(repo)

	// Creating the Fiber app
	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	// Middleware that recovers from panics
	app.Use(recover.New())

	// Middleware that checks if the client exists
	clientes := app.Group("/clientes/:id")

	// Routes
	clientes.Post("/transacoes", handlers.HandleCreateTransacao)
	clientes.Get("/extrato", handlers.HandleGetExtrato)

	app.Listen(":3000")
}