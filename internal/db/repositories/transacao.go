package repositories

import (
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v5"
	t "github.com/samluiz/concurrency-control/internal/types"
)

func (r *Repo) CriarTransacao(transacao *t.TransacaoRequest, clienteId int) (*t.NovaTransacaoResponse, error) {
	clienteExists := r.ClienteExists(clienteId)

	if !clienteExists {
		return nil, ErrClienteNotFound
	}

	tx, err := r.db.Begin(context.Background())
	
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer tx.Rollback(context.Background())
	
	var limite int
	var saldo int
	var novoSaldo int

	err = tx.QueryRow(context.Background(), "SELECT limite, saldo FROM clientes WHERE id = $1", clienteId).Scan(&limite, &saldo)
	
	if transacao.Tipo == "d" {
			novoSaldo = saldo - transacao.Valor
	} else  {
			novoSaldo = saldo + transacao.Valor
	}

	if (limite + novoSaldo) < 0 {
		return nil, ErrInconsistentSaldo
	}

	batch := &pgx.Batch{}

	now := time.Now()

	batch.Queue("INSERT INTO transacoes (valor, tipo, descricao, realizada_em, id_cliente) VALUES ($1, $2, $3, $4, $5)", transacao.Valor, transacao.Tipo, transacao.Descricao, now, clienteId)
	batch.Queue("UPDATE clientes SET saldo = $1 WHERE id = $2", novoSaldo, clienteId)

	batchResults := tx.SendBatch(context.Background(), batch)
	_, err = batchResults.Exec()

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	err = batchResults.Close()

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	err = tx.Commit(context.Background())

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	novaTransacaoResponse := t.NovaTransacaoResponse{
		Limite: limite,
		Saldo: novoSaldo,
	}

	return &novaTransacaoResponse, nil
}