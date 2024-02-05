package repositories

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/jmoiron/sqlx"
	t "github.com/samluiz/concurrency-control/internal/types"
)

var (
	ErrClienteNotFound = errors.New("Cliente não encontrado.")
	ErrInconsistentSaldo = errors.New("Saldo inconsistente.")
)

type Repo struct {
	db *sqlx.DB
}

func NewRepo(db *sqlx.DB) *Repo {
	return &Repo{db}
}

func (r *Repo) CriarTransacao(transacao t.TransacaoRequest, clienteId int) (*t.NovaTransacaoResponse, error) {
	ctx := context.Background()
	tx, err := r.db.BeginTxx(ctx, &sql.TxOptions{Isolation: sql.LevelReadCommitted})

	if err != nil {
		return nil, err
	}
	defer tx.Rollback()
	
	var clienteExists bool
	err = tx.GetContext(ctx, &clienteExists, "SELECT EXISTS (SELECT 1 FROM clientes WHERE id = $1)", clienteId)

	if err != nil {
		return nil, err
	}

	if !clienteExists {
		return nil, ErrClienteNotFound
	}

	now := time.Now()
	
	tx.MustExec("INSERT INTO transacoes (valor, tipo, descricao, realizada_em, id_cliente) VALUES ($1, $2, $3, $4, $5)", transacao.Valor, transacao.Tipo, transacao.Descricao, now, clienteId)
	
	var operation string

	if transacao.Tipo == "d" {
		operation = "-"
	}
	if transacao.Tipo == "c" {
		operation = "+"
	}

	result, err := tx.ExecContext(ctx, "UPDATE clientes SET saldo = saldo " + operation + " $1 WHERE id = $2 FOR UPDATE", transacao.Valor, clienteId)

	if err != nil {
		return nil, err
	}

	if rowsAffected, _ := result.RowsAffected(); rowsAffected == 0 {
		return nil, err
	}

	var novaTransacaoResponse t.NovaTransacaoResponse

	err = tx.GetContext(ctx, &novaTransacaoResponse, "SELECT limite, saldo FROM clientes WHERE id = $1", clienteId)

	if err != nil {
		return nil, err
	}
	
	if (novaTransacaoResponse.Limite < -novaTransacaoResponse.Saldo) {
		return nil, ErrInconsistentSaldo
	}

	err = tx.Commit()

	if err != nil {
		return nil, err
	}

	return &novaTransacaoResponse, nil
}