package main

var Log = initLogger()

func main() {
	dbconf := loadDbConf("dev")
	myapp := App{}
	addr := ":12345"

	func() {
		myapp.initializeWith(dbconf)
		defer myapp.close()
		myapp.handleRoutes()
		myapp.runOn(addr)
	}()
}
