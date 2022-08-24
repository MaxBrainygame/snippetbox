package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"golangify.com/snippetbox/pkg/models"
)

func (app *application) home(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	// filesHtml := []string{
	// 	"./ui/html/home.page.tmpl",
	// 	"./ui/html/base.layout.tmpl",
	// 	"./ui/html/footer.partial.tmpl",
	// }

	// 	homeTemplate, err := template.ParseFiles(filesHtml...)
	// 	if err != nil {
	// 		app.serverError(w, err)
	// 		return
	// 	}
	//
	// 	err = homeTemplate.Execute(w, nil)
	// 	if err != nil {
	// 		app.serverError(w, err)
	// 		return
	// 	}
	//
	// 	fmt.Fprintf(w, "hello! time: %s", time.Now())

	tables, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}

	for _, table := range tables {

		if table == nil {
			continue
		}

		fmt.Fprintf(w, "%v\n", table)
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

	fmt.Fprintf(w, "%v", tableSnippet)

}

func (app *application) createSnippet(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	title := "История про улитку"
	content := "Улитка выползла из раковины,\nвытянула рожки,\nи опять подобрала их."
	expires := "7"

	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, err)
	}

	http.Redirect(w, r, fmt.Sprintf("/snippet?id=%d", id), http.StatusSeeOther)

}
