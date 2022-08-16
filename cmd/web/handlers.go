package main

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"
	"time"

	"github.com/julienschmidt/httprouter"
)

func (app *application) home(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	filesHtml := []string{
		"./ui/html/home.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
	}

	homeTemplate, err := template.ParseFiles(filesHtml...)
	if err != nil {
		app.serverError(w, err)
	}

	err = homeTemplate.Execute(w, nil)
	if err != nil {
		app.serverError(w, err)
	}

	fmt.Fprintf(w, "hello! time: %s", time.Now())
}

func (app *application) showSnippet(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	idString := r.URL.Query().Get("id")
	if len(idString) == 0 {
		fmt.Fprintf(w, "Нужно передать параметр \"id\"")
		return
	}

	idNumber, err := strconv.Atoi(idString)
	if err != nil {
		fmt.Fprintf(w, "Передан некорретный ID")
		app.serverError(w, err)
		return
	}

	fmt.Fprintf(w, "Переданный ID=%d", idNumber)

}

func (app *application) createSnippet(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	// fmt.Fprintf(w, "Должна создаваться заметка")
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"name":"Alex"}`))

}
