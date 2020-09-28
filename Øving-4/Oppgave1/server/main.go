package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// App : an object for holding the database and router
type App struct {
	Router *mux.Router
}

type msgBody struct {
	Value1 int `json:"value1"`
	Value2 int `json:"value2"`
}

type resBody struct {
	Output string `json:"output"`
}

func (app *App) setupRouter() {
	app.Router.Methods("POST").Path("/add").HandlerFunc(app.postAdd)
	app.Router.Methods("POST").Path("/subtract").HandlerFunc(app.postSubtract)
}

func (app *App) postAdd(writer http.ResponseWriter, request *http.Request) {
	data := &msgBody{}
	res := &resBody{}

	err := json.NewDecoder(request.Body).Decode(data)
	if err != nil {
		res.Output = "Could not decode request"
		json.NewEncoder(writer).Encode(res)
		log.Fatal(err)
	} else {
		res.Output = strconv.Itoa(data.Value1 + data.Value2)
		json.NewEncoder(writer).Encode(res)
	}
}

func (app *App) postSubtract(writer http.ResponseWriter, request *http.Request) {
	data := &msgBody{}
	res := &resBody{}

	err := json.NewDecoder(request.Body).Decode(data)
	if err != nil {
		res.Output = "Could not decode request"
		json.NewEncoder(writer).Encode(res)
		log.Fatal(err)
	} else {
		res.Output = strconv.Itoa(data.Value1 - data.Value2)
		json.NewEncoder(writer).Encode(res)
	}
}

func main() {
	app := &App{
		Router: mux.NewRouter().StrictSlash(true),
	}

	app.setupRouter()

	log.Fatal(http.ListenAndServe(":8080", app.Router))
}
