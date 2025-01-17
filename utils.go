package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/sirupsen/logrus"
)

var err_NoMatch = errors.New("NoMatchFound")
var err_NoId = errors.New("IdRequired")
var err_Parse = errors.New("ParseError")
var err_NoEffect = errors.New("NoEffect")

// var err_InvalidId = errors.New("InvalidId")

func sendError(w http.ResponseWriter, statusCode int, err string) {
	response := map[string]string{"err": err}
	sendResponse(w, statusCode, response)
}

func sendResponse(w http.ResponseWriter, statusCode int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		Log.Errorf("Unable to marshal to JSON, review the payload")
		return
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(response)
}

func initLogger() *logrus.Logger {
	var Log = logrus.New()
	// Log.SetLevel(logrus.DebugLevel)
	Log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	Log.Info("Logger Initialized, Welcome to [Books API]")
	return Log
}

func getErrResponse(err error) string {
	switch {
	default:
		errstr := fmt.Sprintf("Server Error: Cannot Process Request [%v]", err.Error())
		Log.Error(errstr)
		return errstr
	case errors.Is(err, err_NoMatch):
		errstr := fmt.Sprintf("Invalid Request: No such record found [%v]", err.Error())
		Log.Error(errstr)
		return errstr
	case errors.Is(err, err_NoId):
		errstr := fmt.Sprintf("Invalid Request: Missing ID field [%v]", err.Error())
		Log.Error(errstr)
		return errstr
	case errors.Is(err, err_Parse):
		errstr := fmt.Sprintf("Invalid Request: Unable to decode JSON request body [%v]", err.Error())
		Log.Error(errstr)
		return errstr
	// case errors.Is(err, err_InvalidId):
	// 	errstr := fmt.Sprintf("Invalid ID: Integer >0 Required [%v]", err.Error())
	// 	Log.Error(errstr)
	// 	return errstr
	case errors.Is(err, err_NoEffect):
		errstr := fmt.Sprintf("Request Processed: No changes required [%v]", err.Error())
		Log.Error(errstr)
		return errstr
	}
}

func makePayload(s string, b []Book) map[string]string {
	res, _ := json.Marshal(b)
	return map[string]string{
		"msg":  s,
		"data": string(res),
	}
}

func loadDbConf(kind string) DbConf {
	switch kind {
	case "dev":
		return DbConf{
			Db_user: "postgres",
			Db_pass: os.Getenv("PG_PASS"),
			Db_addr: "10.0.0.246",
			Db_port: "31540",
			Db_name: "inventory",
			Db_args: "sslmode=disable",
		}
	case "prod":
		{
			return DbConf{
				Db_user: "postgres",
				Db_pass: os.Getenv("PG_PASS"),
				Db_addr: "10.0.0.246",
				Db_port: "31540",
				Db_name: "inventory",
				Db_args: "sslmode=disable",
			}
		}
	case "test":
		{
			return DbConf{
				Db_user: "postgres",
				Db_pass: os.Getenv("PG_PASS"),
				Db_addr: "10.0.0.246",
				Db_port: "31540",
				Db_name: "inventory",
				Db_args: "sslmode=disable",
			}
		}
	default:
		Log.Fatalf("Unknown DB config kind %v", kind)
		return DbConf{}
	}
}
