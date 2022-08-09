package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"
	"time"

	"github.com/julienschmidt/httprouter"
)

func home(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	filesHtml := []string{
		"./ui/html/home.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
	}

	homeTemplate, err := template.ParseFiles(filesHtml...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}

	err = homeTemplate.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}

	fmt.Fprintf(w, "hello! time: %s", time.Now())
}

func showSnippet(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	idString := r.URL.Query().Get("id")
	if len(idString) == 0 {
		fmt.Fprintf(w, "Нужно передать параметр \"id\"")
		return
	}

	idNumber, err := strconv.Atoi(idString)
	if err != nil {
		fmt.Fprintf(w, "Передан некорретный ID")
		log.Println(err)
		return
	}

	fmt.Fprintf(w, "Переданный ID=%d", idNumber)

}

func createSnippet(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	// fmt.Fprintf(w, "Должна создаваться заметка")
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"name":"Alex"}`))

}
