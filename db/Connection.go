package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

func GetConnection() *pgxpool.Pool {
	p, err := pgxpool.New(context.Background(), "postgres://zulfikar:zulfikar@localhost:5432/go_rest_pg")
	if err != nil {
		panic(err)
	}
	return p
}
