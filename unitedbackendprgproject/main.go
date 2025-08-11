package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewPostgres(username, password, host, port, dbName string) (*pgx.Conn, error) {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", username, password, host, port, dbName)
	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func main() {
	//echo
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Static("/", "static")

	//pgx
	conn, err := NewPostgres("postgres", "654321", "localhost", "5432", "test")
	if err != nil {
		log.Fatalf("Error : %v", err)
	}
	defer conn.Close(context.Background())

	e.Logger.Fatal(e.Start(":1323"))
}
