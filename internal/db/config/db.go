package config

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func OpenDB() (*sqlx.DB, error) {
	connStr := "user=username dbname=mydb sslmode=disable"

	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("error opening connection to the database: %v", err)
	}

	return db, nil
}

func initTables(db *sql.DB) error {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS clientes (
			id SERIAL PRIMARY KEY,
			limite INT,
			saldo INT
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
			valor INT,
			tipo CHAR(1),
			descricao VARCHAR(10),
			realizada_em TIMESTAMP,
			id_cliente INT
		);
	`)

	if err != nil {
		return fmt.Errorf("error creating transacoes table: %v", err)
	}

	return nil
}