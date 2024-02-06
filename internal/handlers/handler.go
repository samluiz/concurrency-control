package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/samluiz/concurrency-control/internal/db/repositories"
	"github.com/samluiz/concurrency-control/internal/types"
)

type Handler struct {
	repo *repositories.Repo
}

func NewHandler(repo *repositories.Repo) *Handler {
	return &Handler{repo}
}

func (h Handler) HandleCreateTransacao(c *fiber.Ctx) error {
	clienteId, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	transacaoRequest := types.TransacaoRequest{}

	if err := c.BodyParser(&transacaoRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if len(transacaoRequest.Descricao) > 10 {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"error": "Erro de validação: O campo descrição deve ter no máximo 10 caracteres."})
	}

	novaTransacaoResponse, err := h.repo.CriarTransacao(transacaoRequest, clienteId)

	if err != nil {
		if err == repositories.ErrClienteNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
		}

		if err == repositories.ErrInconsistentSaldo {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"error": err.Error()})
		}
		
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(novaTransacaoResponse)
}

func (h Handler) HandleGetExtrato(c *fiber.Ctx) error {
	clienteId, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	extratoResponse, err := h.repo.ObterExtrato(clienteId)

	if err != nil {
		if err == repositories.ErrClienteNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(extratoResponse)
}