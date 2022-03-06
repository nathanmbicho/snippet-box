package main

import (
	"database/sql"
	"flag"
	"github.com/nathanmbicho/snippetbox/pkg/models/mysql"
	"html/template"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

//hold the application-wide dependencies
type application struct {
	errorLog      *log.Logger
	infoLog       *log.Logger
	snippets      *mysql.SnippetModel
	templateCache map[string]*template.Template
}

//wrap sql.Open() and returns a sql.DB connection pool for given DSN
func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	//if error id found
	if err != nil {
		return nil, err
	}
	//verify/check if connection is still alive and check errors
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

//main function
func main() {
	//command-line flag
	addr := flag.String("addr", ":4040", "HTTP network address")
	//command-line flag for the MySQL DSN string
	dsn := flag.String("dsn", "golang:#G0ph3r1?@/snippetbox?parseTime=true", "MySQL data source name")
	flag.Parse()

	//info and error logger
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	//pass to openDB dsn from command-line flag
	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}

	// defer a call to db.Close(), so that the connection pool is closed before the main() function exits.
	defer db.Close()

	//initialize new template cache
	templateCache, err := newTemplateCache("./ui/html/")
	if err != nil {
		errorLog.Fatal(err)
	}
	//initialize a new instance of application containing the dependencies
	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
		snippets: &mysql.SnippetModel{
			DB: db,
		},
		templateCache: templateCache,
	}

	//initialize http.Server struct
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	//run server
	infoLog.Printf("Server starting on %s", *addr)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}
