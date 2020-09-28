package main

import (
	"html/template"
	"log"
	"net/http"

	"./websocketlib"
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
}

func startServer() {
	log.Println("Websocket listening on port 3001")
	http.ListenAndServe(":3001", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conn, err := websocketlib.MakeConn(w, r)
		if err != nil {
			log.Fatal(err)
		}

		go func() {
			defer conn.Close()

			for {
				msg, err := websocketlib.ReadFromClient(conn)
				if err != nil {
					log.Fatal(err)
				}

				err = websocketlib.MessageAllClients(conn, msg)
				if err != nil {
					log.Fatal(err)
				}
			}
		}()
	}))
}

func main() {
	app := &App{
		Router: mux.NewRouter().StrictSlash(true),
	}

	app.setupRouter()

	go startServer()
	log.Println("HTTP listening on port 3000")
	log.Fatal(http.ListenAndServe(":3000", app.Router))
}
