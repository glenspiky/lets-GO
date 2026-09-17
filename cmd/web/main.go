package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq" // To register the driver.
)

type application struct {
	errorLog *log.Logger
	infolog  *log.Logger
}

func main() {

	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	dsn := flag.String(
		"dsn",
		"postgres://web:1234@localhost:5432/snippetbox?sslmode=disable",
		"PostgreSQL data source name",
	)
	flag.Parse()

	infolog := log.New(os.Stdout, "INFO/t ", log.Ldate|log.Ltime)

	errorLog := log.New(os.Stderr, "ERROR/t ", log.Ldate|log.Ltime|log.Lshortfile)

	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}

	defer db.Close()

	// initialize instance of application cantainig the dependanciies
	app := &application{
		errorLog: errorLog,
		infolog:  infolog,
	}

	srv := &http.Server{Addr: *addr, ErrorLog: errorLog, Handler: app.routes()}

	infolog.Printf("Starting server on %s\n", *addr)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)

	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
