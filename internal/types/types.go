package types

type Cliente struct {
	ID     int `db:"id"`
	Limite int `db:"limite"`
	Saldo  int `db:"saldo"`
}

type Transacao struct {
	ID          int    `db:"id"`
	Valor       int    `db:"valor"`
	Tipo        string `db:"tipo"`
	Descricao   string `db:"descricao"`
	RealizadaEm string `db:"realizada_em"`
	IdCliente   int    `db:"id_cliente"`
}