package web

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type App struct {
	Router *mux.Router
}

func newApp() *App {
	router := mux.NewRouter()
	router.HandleFunc("/user/search", getUser()).Methods("POST")
	return &App{Router: router}
}

func (a *App) start() {
	log.Fatal(http.ListenAndServe(":8888", a.Router))
}

func main() {
	newApp().start()
}
