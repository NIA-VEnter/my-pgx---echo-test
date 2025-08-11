package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
)

func ConnectPGX() *pgx.Conn {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", "postgres", "654321", "localhost", "5432", "test")
	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		log.Fatalf("Error : %v", err)
	}
	defer conn.Close(context.Background())
	return conn
}
