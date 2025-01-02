package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5"
)

type App struct {
	router *mux.Router
	db     *pgx.Conn
}

func (app *App) initialize() {
	dbconf := DbConfig{
		Db_user: "postgres",
		Db_pass: os.Getenv("PG_PASS"),
		Db_addr: "10.0.0.246",
		Db_port: "31540",
		Db_name: "inventory",
		Db_args: "sslmode=disable",
	}

	app.db = dbconf.getConnection()
	app.router = mux.NewRouter().StrictSlash(true)
}

func (app *App) run(addr string) {
	fmt.Printf("Serving... on %v\n", addr)
	err := http.ListenAndServe(addr, app.router)
	if err != nil {
		Log.Fatal(err)
	}
}

func (app *App) handleRoutes() {
	app.router.HandleFunc("/", app.getHomepage).Methods("GET")
	app.router.HandleFunc("/books/", app.getAllBooks).Methods("GET")
	app.router.HandleFunc("/book/{id}", app.getBook).Methods("GET")

}

func (app *App) getHomepage(w http.ResponseWriter, r *http.Request) {
	Log.Debugf("Endpoint %v", r.URL.Path)
	sendResponse(w, 200, map[string]string{"msg": "Welcome to HomePage"})
}

func (app *App) getAllBooks(w http.ResponseWriter, r *http.Request) {
	Log.Debugf("Endpoint %v", r.URL.Path)
	app.handleBooks(w, r, "all")
}

func (app *App) getBook(w http.ResponseWriter, r *http.Request) {
	Log.Debugf("Endpoint %v", r.URL.Path)
	app.handleBooks(w, r, mux.Vars(r)["id"])
}

func (app *App) handleBooks(w http.ResponseWriter, r *http.Request, id string) {
	Log.Debugf("Endpoint %v", r.URL.Path)
	books, err := app.fetchData(id)
	if err != nil {
		sendError(w, http.StatusInternalServerError, "Server Error in Processing Request")
		return
	}
	sendResponse(w, http.StatusOK, books)
}
