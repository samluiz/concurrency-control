package repositories

import (
	"errors"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewValidationError(message string) error {
	return errors.New(message)
}

var (
	ErrClienteNotFound   = errors.New("Cliente não encontrado.")
	ErrInconsistentSaldo = errors.New("Saldo inconsistente.")
	ErrValidation        = errors.New("Erro de validação.")
)

type Repo struct {
	db *pgxpool.Pool
}

func NewRepo(db *pgxpool.Pool) *Repo {
	return &Repo{db}
}