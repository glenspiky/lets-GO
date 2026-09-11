package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

func main() {

	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	infolog := log.New(os.Stdout, "INFO/t ", log.Ldate|log.Ltime)

	errorLog := log.New(os.Stderr, "ERROR/t ", log.Ldate|log.Ltime|log.Lshortfile)

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/create", createSnippet)

	fileServer := http.FileServer(http.Dir("./ui/static/"))

	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	infolog.Printf("Starting server on %s\n", *addr)
	if err := http.ListenAndServe(*addr, mux); err != nil {
		errorLog.Fatal(err)
	}
}
