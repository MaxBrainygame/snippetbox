package main

import (
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
	"github.com/spf13/viper"
)

type Config struct {
	port        string
	staticfiles string
}

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

	viper.SetConfigName("config")
	viper.AddConfigPath("$Home/Projects/snippetbox")
	err = viper.ReadInConfig()
	if err != nil {
		errorLog.Fatal("fatal error config file: %w", err)
	}

	var config Config
	config.port = viper.GetString("port")
	config.staticfiles = viper.GetString("staticfiles")

	router := httprouter.New()
	router.GET("/", app.home)
	router.GET("/hello/:name", app.home)
	router.GET("/snippet", app.showSnippet)
	router.POST("/snippet/create", app.createSnippet)

	router.ServeFiles("/static/*filepath", http.Dir(config.staticfiles))

	srv := &http.Server{
		Addr:     config.port,
		ErrorLog: errorLog,
		Handler:  router,
	}

	infoLog.Printf("Запуск веб-сервера на http://127.0.0.1%s", config.port)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}
