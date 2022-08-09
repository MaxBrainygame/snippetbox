package main

import (
	"log"
	"net/http"
	"path/filepath"

	"github.com/julienschmidt/httprouter"
)

type neuteredFileSystem struct {
	fs http.FileSystem
}

func main() {

	router := httprouter.New()
	router.GET("/", home)
	router.GET("/hello/:name", home)
	router.GET("/snippet", showSnippet)
	router.POST("/snippet/create", createSnippet)

	router.ServeFiles("/static/*filepath", http.Dir("./ui/static"))

	log.Println("Запуск веб-сервера на http://127.0.0.1:4000")
	err := http.ListenAndServe(":4000", router)
	log.Fatal(err)
}

func (nfs *neuteredFileSystem) Open(path string) (http.File, error) {

	file, err := nfs.fs.Open(path)
	if err != nil {
		return nil, err
	}

	infoFile, err := file.Stat()
	if err != nil {
		return nil, err
	}

	if infoFile.IsDir() {

		index := filepath.Join(path, "index.html")
		if _, err := nfs.fs.Open(index); err != nil {
			closeErr := file.Close()
			if closeErr != nil {
				return nil, closeErr
			}

			return nil, err
		}

	}

	return file, nil
}
