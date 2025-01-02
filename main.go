package main

import (
	"github.com/sirupsen/logrus"
)

var Log = logrus.New()

func main() {
	Log.SetLevel(logrus.DebugLevel)
	Log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	Log.Debug("Inside Main")

	myapp := App{}
	myapp.initialize()
	myapp.handleRoutes()

	myapp.run(":12345")

}
