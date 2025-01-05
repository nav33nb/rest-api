package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
	"time"
)

const addr = ":23456"

type Response struct {
	Msg  string `json:"msg"`
	Data []Book `json:"data"`
}

func TestApp(t *testing.T) {
	setupServerApi()

	client := &http.Client{Timeout: 5 * time.Second}

	// Testing Functions

	var books []Book
	resp, err := client.Get(fmt.Sprintf("http://localhost%v/books", addr))
	if err != nil {
		t.Error(err)
	}
	err = json.Unmarshal(resp.Body)
	if err != nil {
		t.Fatalf("Cannot unmarshal response to json")
	}
	t.Log(resp.StatusCode)
	t.Log(len(books))
}

func setupServerApi() {
	go func() {
		app := App{}
		app.initializeWith(loadDbConf("test"))
		defer app.close()
		app.handleRoutes()
		app.runOn(addr)
	}()
	Log.Infof("Waiting 5s for server to come up")
	time.Sleep(5 * time.Second)
}
