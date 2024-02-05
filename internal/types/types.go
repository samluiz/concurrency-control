package types

type Cliente struct {
	ID     int
	Limite float64
	Saldo  float64
}

type Transacao struct {
	ID          int
	Valor       float64
	Tipo        string
	Descricao   string
	RealizadaEm string
	IdCliente   int
}