package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/samluiz/concurrency-control/internal/db/repositories"
	"github.com/samluiz/concurrency-control/internal/types"
)

func (h Handler) HandleCreateTransacao(c *fiber.Ctx) error {
	clienteId, err := c.ParamsInt("id")

	if err != nil {
		return c.SendStatus(fiber.StatusUnprocessableEntity)
	}

	transacaoRequest := types.TransacaoRequest{}

	if err := c.BodyParser(&transacaoRequest); err != nil {
		return c.SendStatus(fiber.StatusUnprocessableEntity)
	}

	if transacaoRequest.Descricao == "" || len(transacaoRequest.Descricao) > 10 {
		return c.SendStatus(fiber.StatusUnprocessableEntity)
	}

	if transacaoRequest.Tipo != "d" && transacaoRequest.Tipo != "c" {
		return c.SendStatus(fiber.StatusUnprocessableEntity)
	}

	novaTransacaoResponse, err := h.repo.CriarTransacao(&transacaoRequest, clienteId)

	if err != nil {
		if err == repositories.ErrClienteNotFound {
			return c.SendStatus(fiber.StatusNotFound)
		}

		if err == repositories.ErrInconsistentSaldo || err == repositories.ErrValidation {
			return c.SendStatus(fiber.StatusUnprocessableEntity)
		}

		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(novaTransacaoResponse)
}