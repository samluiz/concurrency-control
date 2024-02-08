package repositories

import (
	"context"
	"log"
)

var clienteCache = make(map[int]bool)

func (r *Repo) ClienteExists(clienteId int) bool {
	cachedClienteExists, ok := clienteCache[clienteId]

	if ok {
		return cachedClienteExists
	}
	
	var clienteExists bool
	err := r.db.QueryRow(context.Background(), "SELECT EXISTS (SELECT 1 FROM clientes WHERE id = $1) FOR UPDATE", clienteId).Scan(&clienteExists)

	if err != nil {
		log.Fatal(err)
		return false
	}

	clienteCache[clienteId] = clienteExists

	return clienteExists
}