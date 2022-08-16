package main

import (
	"log"
	"net/http"
	"os"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {

	f, err := os.OpenFile("info.log", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	infoLog := log.New(f, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	config := app.newConfig()

	srv := &http.Server{
		Addr:     config.port,
		ErrorLog: errorLog,
		Handler:  app.routes(config),
	}

	infoLog.Printf("Запуск веб-сервера на http://127.0.0.1%s", config.port)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}
