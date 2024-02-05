package types

type Cliente struct {
	ID     int
	Limite int
	Saldo  int
}

type Transacao struct {
	ID          int
	Valor       int
	Tipo        string
	Descricao   string
	RealizadaEm string
	IdCliente   int
}