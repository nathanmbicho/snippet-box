package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

//hold the application-wide dependencies
type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

//main function
func main() {
	//command-line flag
	addr := flag.String("addr", ":4040", "HTTP network address")
	flag.Parse()

	//info and error logger
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	//initialize a new instance of application containing the dependencies
	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	//initialize http.Server struct
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	//run server
	infoLog.Printf("Server starting on %s", *addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
