package database

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

func ConnectDB() {
	// urlExample := "postgres://username:password@localhost:5432/database_name"
  dbpoll, err := pgxpool.Connect(context.Background(), os.Getenv("DATABASE_URL"))
  if err != nil {
    fmt.Fprintf(os.Stderr, "Unable to connect to the database: %v\n", err)
    os.Exit(1)
  }

  defer dbpoll.Close()

  var greeting string
  err = dbpoll.QueryRow(context.Background(), "select 'Hello, world!'").Scan(&greeting)
  if err != nil {
    fmt.Fprintf(os.Stderr, "Query Row failed: %v\n", err)
    os.Exit(1)
  }

  fmt.Println(greeting)
}
