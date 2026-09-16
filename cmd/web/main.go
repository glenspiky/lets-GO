package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

type application struct {
	errorLog *log.Logger
	infolog  *log.Logger
}

func main() {

	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	infolog := log.New(os.Stdout, "INFO/t ", log.Ldate|log.Ltime)

	errorLog := log.New(os.Stderr, "ERROR/t ", log.Ldate|log.Ltime|log.Lshortfile)

	// initialize instance of application cantainig the dependanciies
	app := &application{
		errorLog: errorLog,
		infolog:  infolog,
	}

	srv := &http.Server{Addr: *addr, ErrorLog: errorLog, Handler: app.routes()}

	infolog.Printf("Starting server on %s\n", *addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
