package config

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func OpenDB() (*sqlx.DB, error) {
	connStr := "host=localhost port=2345 user=postgres password=postgres dbname=bank sslmode=disable"

	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("error opening connection to the database: %v", err)
	}

	err = initTables(db.DB)

	if err != nil {
		return nil, fmt.Errorf("error initializing tables: %v", err)
	}

	return db, nil
}

func initTables(db *sql.DB) error {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS clientes (
			id SERIAL PRIMARY KEY,
			limite INT,
			saldo INT DEFAULT 0
		);
	`)

	if err != nil {
		return fmt.Errorf("error creating clientes table: %v", err)
	}

	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM clientes WHERE id BETWEEN 1 AND 5").Scan(&count)
	if err != nil {
		return fmt.Errorf("error checking existing rows: %v", err)
	}

	if count == 0 {
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