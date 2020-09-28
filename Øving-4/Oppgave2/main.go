package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var tpl = template.Must(template.ParseFiles("index.html"))

// App : an object for holding the database and router
type App struct {
	Router *mux.Router
}

func (app *App) setupRouter() {
	app.Router.Methods("GET").Path("/").HandlerFunc(app.getPage)
}

func (app *App) getPage(writer http.ResponseWriter, request *http.Request) {
	tpl.Execute(writer, nil)
	writer.Write([]byte("<h1> Hilsen. Du har koblet deg opp til min enkle web-tjener </h1> Header fra klienten er: <ul>"))
	for _, header := range request.Header {
		writer.Write([]byte("<li>" + fmt.Sprint(header) + "</li>"))
	}
	writer.Write([]byte("</ul>"))
}

func main() {
	app := &App{
		Router: mux.NewRouter().StrictSlash(true),
	}

	app.setupRouter()

	log.Fatal(http.ListenAndServe(":80", app.Router))
}
