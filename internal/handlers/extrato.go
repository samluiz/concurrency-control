package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/samluiz/concurrency-control/internal/db/repositories"
)

func (h Handler) HandleGetExtrato(c *fiber.Ctx) error {
	clienteId, err := c.ParamsInt("id")

	if err != nil {
		return c.SendStatus(fiber.StatusUnprocessableEntity)
	}

	extratoResponse, err := h.repo.ObterExtrato(clienteId)

	if err != nil {
		if err == repositories.ErrClienteNotFound {
			return c.SendStatus(fiber.StatusNotFound)
		}

		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(extratoResponse)
}