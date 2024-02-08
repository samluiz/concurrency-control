package repositories

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
	t "github.com/samluiz/concurrency-control/internal/types"
)

func (r *Repo) ObterExtrato(clienteId int) (*t.ExtratoResponse, error) {

	clienteExists := r.ClienteExists(clienteId)

	if !clienteExists {
		return nil, ErrClienteNotFound
	}

	rows, _ := r.db.Query(context.Background(), "SELECT saldo, limite, now() as data_extrato FROM clientes WHERE id = $1 FOR UPDATE", clienteId)

	saldoResponse, err := pgx.CollectOneRow(rows, pgx.RowToStructByPos[t.SaldoResponse])

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	rows, _ = r.db.Query(context.Background(), "SELECT valor, tipo, descricao, realizada_em FROM transacoes WHERE id_cliente = $1 ORDER BY realizada_em DESC LIMIT 10 FOR UPDATE", clienteId)

	transacoes, err := pgx.CollectRows(rows, pgx.RowToStructByPos[t.TransacaoResponse])

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return &t.ExtratoResponse{Saldo: saldoResponse, UltimasTransacoes: transacoes}, nil
}