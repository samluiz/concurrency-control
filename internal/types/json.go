package types

import "time"

type TransacaoRequest struct {
	Valor     int `json:"valor"`
	Tipo      string `json:"tipo"`
	Descricao string `json:"descricao"`
}

type TransacaoResponse struct {
	Valor     int `json:"valor"`
	Tipo      string `json:"tipo"`
	Descricao string `json:"descricao"`
	RealizadaEm string `json:"realizada_em"`
}

type NovaTransacaoResponse struct {
	Limite int `json:"limite"`
	Saldo  int `json:"saldo"`
}

type SaldoResponse struct {
	Total       int `json:"total"`
	Limite      int `json:"limite"`
	DataExtrato time.Time `json:"data_extrato"`
}

type ExtratoResponse struct {
	Saldo SaldoResponse `json:"saldo"`
	UltimasTransacoes []TransacaoResponse `json:"ultimas_transacoes"`
}