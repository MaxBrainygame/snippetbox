package main

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"golangify.com/snippetbox/pkg/models"
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
		return
	}

	tables, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}

	data := &templateData{Snippets: tables}

	err = homeTemplate.Execute(w, data)
	if err != nil {
		app.serverError(w, err)
		return
	}

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

	tableSnippet, err := app.snippets.Get(idNumber)
	if err != nil {

		if errors.Is(err, models.ErrNoRecord) {
			fmt.Fprintf(w, "%v", err)
		} else {
			app.serverError(w, err)
		}
		return
	}

	data := &templateData{Snippet: tableSnippet}

	files := []string{
		"./ui/html/show.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
	}

	template, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
		return
	}

	err = template.Execute(w, data)
	if err != nil {
		app.serverError(w, err)
		return
	}

	fmt.Fprintf(w, "%v", tableSnippet)

}

func (app *application) createSnippet(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	err := r.ParseForm()
	if err != nil {
		app.serverError(w, err)
		return
	}

	id, err := app.snippets.Insert(r.Form.Get("title"), r.Form.Get("content"), r.Form.Get("expires"))
	if err != nil {
		app.serverError(w, err)
	}

	http.Redirect(w, r, fmt.Sprintf("/snippet?id=%d", id), http.StatusSeeOther)

}
