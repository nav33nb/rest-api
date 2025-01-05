package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5"
)

type App struct {
	router *mux.Router
	db     *pgx.Conn
}

func (app *App) initializeWith(dbconf DbConf) {
	Log.Infof("Initializing App")

	app.db = dbconf.getConnection()
	initDatabase(app.db)
	app.router = mux.NewRouter().StrictSlash(false)
}

func (app *App) close() error {
	return app.db.Close(context.Background())
}

func (app *App) runOn(addr string) {
	Log.Infof("Serving Application on %v", addr)
	err := http.ListenAndServe(addr, app.router)
	if err != nil {
		Log.Fatal(err)
	}
}

func (app *App) handleRoutes() {
	Log.Infof("Handling Routes")

	app.router.HandleFunc("/", app.getHomepage).Methods("GET")

	app.router.HandleFunc("/books", app.getBook).Methods("GET")
	app.router.HandleFunc("/books", app.postBook).Methods("POST")
	app.router.HandleFunc("/books", app.putBook).Methods("PUT")

	app.router.HandleFunc("/books/{id}", app.deleteBook).Methods("DELETE")
	app.router.HandleFunc("/books/{id}", app.getBook).Methods("GET")
}

func (app *App) getHomepage(w http.ResponseWriter, r *http.Request) {
	Log.Infof("Hit GET Endpoint %v, with Request Params %v", r.URL.Path, mux.Vars(r))
	Log.Debugf("Endpoint %v", r.URL.Path)
	sendResponse(w, 200, makePayload("Welcome to HomePage", nil))
}

func (app *App) getBook(w http.ResponseWriter, r *http.Request) {
	Log.Infof("Hit GET Endpoint %v, with Request Params %v, %v", r.URL.Path, mux.Vars(r), r.URL.Query().Get("id"))
	books, err := fetchData(app.db, mux.Vars(r)["id"])
	if err != nil {
		sendError(w, http.StatusInternalServerError, getErrResponse(err))
		return
	}
	sendResponse(w, http.StatusOK, makePayload("OK: Data Enclosed", books))
}

// 1. decode the request body
// 2. put decoding INTO an object
// 3. call sql with that object
func (app *App) postBook(w http.ResponseWriter, r *http.Request) {
	Log.Infof("Hit POST Endpoint %v, with Request Params %v", r.URL.Path, mux.Vars(r))

	var b Book
	err := json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		sendError(w, http.StatusBadRequest, getErrResponse(err_Parse))
		Log.Errorf("Skipped record creation: Unable to parse request body: %v", err)
		return
	}

	Log.Debugf("POST Request Body %#v", b)
	err = postData(app.db, b)
	if err != nil {
		sendError(w, http.StatusInternalServerError, getErrResponse(err))
		return
	}
	sendResponse(w, http.StatusCreated, makePayload("record created", nil))
}

func (app *App) putBook(w http.ResponseWriter, r *http.Request) {
	Log.Infof("Hit PUT Endpoint %v, with Request Params %v", r.URL.Path, mux.Vars(r))

	var b Book
	err := json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		sendError(w, http.StatusBadRequest, getErrResponse(err_Parse))
		return
	}
	Log.Debugf("PUT Request Body %#v", b)

	err = putData(app.db, b)
	if err != nil {
		sendError(w, http.StatusInternalServerError, getErrResponse(err))
		return
	}
	sendResponse(w, http.StatusOK, makePayload("record updated", nil))
}

func (app *App) deleteBook(w http.ResponseWriter, r *http.Request) {
	Log.Infof("Hit DELETE Endpoint %v, with Request Params %v", r.URL.Path, mux.Vars(r))
	id := mux.Vars(r)["id"]
	err := deleteData(app.db, id)
	if err != nil {
		sendError(w, http.StatusInternalServerError, getErrResponse(err))
		return
	}
	sendResponse(w, http.StatusOK, makePayload("record deleted", nil))
}
