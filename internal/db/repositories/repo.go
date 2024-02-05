package repositories

import (
	"database/sql"

	t "github.com/samluiz/concurrency-control/types"
)

type Repo struct {
	db *sql.DB
}

func NewRepo(db *sql.DB) *Repo {
	return &Repo{db}
}

func (r *Repo) criarTransacao(transacao t.TransacaoRequest)