package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"

	"github.com/gorilla/mux"
)

// App : an object for holding the database and router
type App struct {
	Router *mux.Router
}

type msgBody struct {
	Code     string `json:"code"`
	Language string `json:"language"`
}

type resBody struct {
	Output string `json:"output"`
}

func (app *App) setupRouter() {
	app.Router.Methods("POST").Path("/code").HandlerFunc(app.postCode)
}

func (app *App) postCode(writer http.ResponseWriter, request *http.Request) {
	data := &msgBody{}

	err := json.NewDecoder(request.Body).Decode(data)
	if err != nil {
		log.Print("Could not decode json body")
		log.Fatal(err)
	}

	code := []byte(data.Code)

	switch data.Language {
	case "golang":
		ioutil.WriteFile("./runnerfiles/golangrunner.go", code, 0644)
		out, err := exec.Command("go", "run", "runnerfiles/golangrunner.go").Output()
		res := &resBody{}
		if err != nil {
			res.Output = err.Error()
		} else {
			res.Output = string(out)
		}
		json.NewEncoder(writer).Encode(res)
	case "javascript":

	case "python":
		ioutil.WriteFile("./runnerfiles/pythonrunner.py", code, 0644)
		out, err := exec.Command("python3", "runnerfiles/pythonrunner.py").Output()
		res := &resBody{}
		if err != nil {
			res.Output = err.Error()
		} else {
			res.Output = string(out)
		}
		json.NewEncoder(writer).Encode(res)
	case "java":

	case "c++":

	case "c":

	case "rust":

	case "c#":

	}
}

func main() {
	app := &App{
		Router: mux.NewRouter().StrictSlash(true),
	}

	app.setupRouter()

	log.Fatal(http.ListenAndServe(":8080", app.Router))
}
