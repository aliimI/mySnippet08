package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"se08.com/pkg/models/postgres"

	"github.com/jackc/pgx/v4/pgxpool"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	snippets *postgres.SnippetModel
}

func main() {

	addr := flag.String("addr", ":4000", "HTTP network address")

	os.Setenv("dsn", "user=postgres password=qwe123 host=localhost port=5432 dbname=snippetbox sslmode=disable pool_max_conns=10")

	dsn := flag.String("dsn", os.Getenv("dsn"), "Postgres data source name")

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime)

	pool, err := pgxpool.Connect(context.Background(), *dsn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	defer pool.Close()

	app := &application{
		errorLog: errLog,
		infoLog:  infoLog,
		snippets: &postgres.SnippetModel{Pool: pool},
	}

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errLog,
		Handler:  app.routes(),
	}

	infoLog.Print("Starting server on %s", *addr)
	err = srv.ListenAndServe()
	errLog.Fatal(err)
}
