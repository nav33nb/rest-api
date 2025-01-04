package main

import "github.com/sirupsen/logrus"

var Log *logrus.Logger

func main() {
	Log = initLogger()

	myapp := App{}
	myapp.initialize()
	defer myapp.close()

	myapp.handleRoutes()
	addr := ":12345"
	myapp.runOn(addr)
}
