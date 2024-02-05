package config

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func OpenDB() (*sql.DB, error) {
	connStr := "user=username dbname=mydb sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("error opening connection to the database: %v", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("error pinging database: %v", err)
	}

	return db, nil
}

func initTables(db *sql.DB) error {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS clientes (
			id SERIAL PRIMARY KEY,
			limite FLOAT,
			saldo FLOAT
		);
	`)

	if err != nil {
		return fmt.Errorf("error creating clientes table: %v", err)
	}

	_, err = db.Exec(`
			INSERT INTO clientes (limite) 
			VALUES 
			(1000 * 100),
			(800 * 100),
			(10000 * 100),
			(100000 * 100),
			(5000 * 100);
	`)

	if err != nil {
		return fmt.Errorf("error inserting initial values in clientes table: %v", err)
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS transacoes (
			id SERIAL PRIMARY KEY,
			valor FLOAT,
			tipo CHAR(1),
			descricao TEXT,
			realizada_em TIMESTAMP,
			id_cliente INT
		);
	`)
	if err != nil {
		return fmt.Errorf("error creating transacoes table: %v", err)
	}

	return nil
}