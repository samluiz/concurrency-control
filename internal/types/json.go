package types

import "time"

type TransacaoRequest struct {
	Valor     float64 `json:"valor"`
	Tipo      string `json:"tipo"`
	Descricao string `json:"descricao"`
}

type TransacaoResponse struct {
	Valor     float64 `json:"valor"`
	Tipo      string `json:"tipo"`
	Descricao string `json:"descricao"`
	RealizadaEm string `json:"realizada_em"`
}

type NovaTransacaoResponse struct {
	Limite float64 `json:"limite"`
	Saldo  float64 `json:"saldo"`
}

type SaldoResponse struct {
	Total       float64 `json:"total"`
	Limite      float64 `json:"limite"`
	DataExtrato time.Time `json:"data_extrato"`
}

type ExtratoResponse struct {
	Saldo SaldoResponse `json:"saldo"`
	UltimasTransacoes []TransacaoResponse `json:"ultimas_transacoes"`
}