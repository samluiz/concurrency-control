package config

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
)

func OpenDB() (*pgxpool.Pool, error) {
	db, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	
	if err != nil {
		return nil, err
	}

	return db, nil
}